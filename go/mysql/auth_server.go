/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mysql

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net"
	"strings"

	"github.com/mushiyu/vitess/go/vt/log"
)

// AuthServer is the interface that servers must implement to validate
// users and passwords. It has two modes:
//
// 1. using salt the way MySQL native auth does it. In that case, the
// password is not sent in the clear, but the salt is used to hash the
// password both on the client and server side, and the result is sent
// and compared.
//
// 2. sending the user / password in the clear (using MySQL Cleartext
// method). The server then gets access to both user and password, and
// can authenticate using any method. If SSL is not used, it means the
// password is sent in the clear. That may not be suitable for some
// use cases.
type AuthServer interface {
	// AuthMethod returns the authentication method to use for the
	// given user. If this returns MysqlNativePassword
	// (mysql_native_password), then ValidateHash() will be
	// called, and no further roundtrip with the client is
	// expected. If anything else is returned, Negotiate()
	// will be called on the connection, and the AuthServer
	// needs to handle the packets.
	AuthMethod(user string) (string, error)

	// Salt returns the salt to use for a connection.
	// It should be 20 bytes of data.
	// Most implementations should just use mysql.NewSalt().
	// (this is meant to support a plugin that would use an
	// existing MySQL server as the source of auth, and just forward
	// the salt generated by that server).
	// Do not return zero bytes, as a known salt can be the source
	// of a crypto attack.
	Salt() ([]byte, error)

	// ValidateHash validates the data sent by the client matches
	// what the server computes.  It also returns the user data.
	ValidateHash(salt []byte, user string, authResponse []byte, remoteAddr net.Addr) (Getter, error)

	// Negotiate is called if AuthMethod returns anything else
	// than MysqlNativePassword. It is handed the connection after the
	// AuthSwitchRequest packet is sent.
	// - If the negotiation fails, it should just return an error
	// (should be a SQLError if possible).
	// The framework is responsible for writing the Error packet
	// and closing the connection in that case.
	// - If the negotiation works, it should return the Getter,
	// and no error. The framework is responsible for writing the
	// OK packet.
	Negotiate(c *Conn, user string, remoteAddr net.Addr) (Getter, error)
}

// authServers is a registry of AuthServer implementations.
var authServers = make(map[string]AuthServer)

// RegisterAuthServerImpl registers an implementations of AuthServer.
func RegisterAuthServerImpl(name string, authServer AuthServer) {
	if _, ok := authServers[name]; ok {
		log.Fatalf("AuthServer named %v already exists", name)
	}
	authServers[name] = authServer
}

// GetAuthServer returns an AuthServer by name, or log.Exitf.
func GetAuthServer(name string) AuthServer {
	authServer, ok := authServers[name]
	if !ok {
		log.Exitf("no AuthServer name %v registered", name)
	}
	return authServer
}

// NewSalt returns a 20 character salt.
func NewSalt() ([]byte, error) {
	salt := make([]byte, 20)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	// Salt must be a legal UTF8 string.
	for i := 0; i < len(salt); i++ {
		salt[i] &= 0x7f
		if salt[i] == '\x00' || salt[i] == '$' {
			salt[i]++
		}
	}

	return salt, nil
}

// ScramblePassword computes the hash of the password using 4.1+ method.
func ScramblePassword(salt, password []byte) []byte {
	if len(password) == 0 {
		return nil
	}

	// stage1Hash = SHA1(password)
	crypt := sha1.New()
	crypt.Write(password)
	stage1 := crypt.Sum(nil)

	// scrambleHash = SHA1(salt + SHA1(stage1Hash))
	// inner Hash
	crypt.Reset()
	crypt.Write(stage1)
	hash := crypt.Sum(nil)
	// outer Hash
	crypt.Reset()
	crypt.Write(salt)
	crypt.Write(hash)
	scramble := crypt.Sum(nil)

	// token = scrambleHash XOR stage1Hash
	for i := range scramble {
		scramble[i] ^= stage1[i]
	}
	return scramble
}

func isPassScrambleMysqlNativePassword(reply, salt []byte, mysqlNativePassword string) bool {
	/*
		SERVER:  recv(reply)
				 hash_stage1=xor(reply, sha1(salt,hash))
				 candidate_hash2=sha1(hash_stage1)
				 check(candidate_hash2==hash)
	*/
	if len(reply) == 0 {
		return false
	}

	if mysqlNativePassword == "" {
		return false
	}

	if strings.Contains(mysqlNativePassword, "*") {
		mysqlNativePassword = mysqlNativePassword[1:]
	}

	hash, err := hex.DecodeString(mysqlNativePassword)
	if err != nil {
		return false
	}

	// scramble = SHA1(salt+hash)
	crypt := sha1.New()
	crypt.Write(salt)
	crypt.Write(hash)
	scramble := crypt.Sum(nil)

	// token = scramble XOR stage1Hash
	for i := range scramble {
		scramble[i] ^= reply[i]
	}
	hashStage1 := scramble

	crypt.Reset()
	crypt.Write(hashStage1)
	candidateHash2 := crypt.Sum(nil)

	if bytes.Compare(candidateHash2, hash) != 0 {
		return false
	}
	return true
}

// Constants for the dialog plugin.
const (
	mysqlDialogMessage = "Enter password: "

	// Dialog plugin is similar to clear text, but can respond to multiple
	// prompts in a row. This is not yet implemented.
	// Follow questions should be prepended with a `cmd` byte:
	// 0x02 - ordinary question
	// 0x03 - last question
	// 0x04 - password question
	// 0x05 - last password
	mysqlDialogAskPassword = 0x04
)

// authServerDialogSwitchData is a helper method to return the data
// needed in the AuthSwitchRequest packet for the dialog plugin
// to ask for a password.
func authServerDialogSwitchData() []byte {
	result := make([]byte, len(mysqlDialogMessage)+2)
	result[0] = mysqlDialogAskPassword
	writeNullString(result, 1, mysqlDialogMessage)
	return result
}

// AuthServerReadPacketString is a helper method to read a packet
// as a null terminated string. It is used by the mysql_clear_password
// and dialog plugins.
func AuthServerReadPacketString(c *Conn) (string, error) {
	// Read a packet, the password is the payload, as a
	// zero terminated string.
	data, err := c.ReadPacket()
	if err != nil {
		return "", err
	}
	if len(data) == 0 || data[len(data)-1] != 0 {
		return "", fmt.Errorf("received invalid response packet, datalen=%v", len(data))
	}
	return string(data[:len(data)-1]), nil
}

// AuthServerNegotiateClearOrDialog will finish a negotiation based on
// the method type for the connection. Only supports
// MysqlClearPassword and MysqlDialog.
func AuthServerNegotiateClearOrDialog(c *Conn, method string) (string, error) {
	switch method {
	case MysqlClearPassword:
		// The password is the next packet in plain text.
		return AuthServerReadPacketString(c)

	case MysqlDialog:
		return AuthServerReadPacketString(c)

	default:
		return "", fmt.Errorf("unrecognized method: %v", method)
	}
}
