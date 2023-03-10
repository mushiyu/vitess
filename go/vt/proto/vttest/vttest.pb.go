// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vttest.proto

package vttest // import "github.com/mushiyu/vitess/go/vt/proto/vttest"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Shard describes a single shard in a keyspace.
type Shard struct {
	// name has to be unique in a keyspace. For unsharded keyspaces, it
	// should be '0'. For sharded keyspace, it should be derived from
	// the keyrange, like '-80' or '40-80'.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// db_name_override is the mysql db name for this shard. Has to be
	// globally unique. If not specified, we will by default use
	// 'vt_<keyspace>_<shard>'.
	DbNameOverride       string   `protobuf:"bytes,2,opt,name=db_name_override,json=dbNameOverride,proto3" json:"db_name_override,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Shard) Reset()         { *m = Shard{} }
func (m *Shard) String() string { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()    {}
func (*Shard) Descriptor() ([]byte, []int) {
	return fileDescriptor_vttest_beaece7261b82562, []int{0}
}
func (m *Shard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shard.Unmarshal(m, b)
}
func (m *Shard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shard.Marshal(b, m, deterministic)
}
func (dst *Shard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shard.Merge(dst, src)
}
func (m *Shard) XXX_Size() int {
	return xxx_messageInfo_Shard.Size(m)
}
func (m *Shard) XXX_DiscardUnknown() {
	xxx_messageInfo_Shard.DiscardUnknown(m)
}

var xxx_messageInfo_Shard proto.InternalMessageInfo

func (m *Shard) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Shard) GetDbNameOverride() string {
	if m != nil {
		return m.DbNameOverride
	}
	return ""
}

// Keyspace describes a single keyspace.
type Keyspace struct {
	// name has to be unique in a VTTestTopology.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// shards inside this keyspace. Ignored if redirect is set.
	Shards []*Shard `protobuf:"bytes,2,rep,name=shards,proto3" json:"shards,omitempty"`
	// sharding_column_name for this keyspace. Used for v2 calls, but not for v3.
	ShardingColumnName string `protobuf:"bytes,3,opt,name=sharding_column_name,json=shardingColumnName,proto3" json:"sharding_column_name,omitempty"`
	// sharding_column_type for this keyspace. Used for v2 calls, but not for v3.
	ShardingColumnType string `protobuf:"bytes,4,opt,name=sharding_column_type,json=shardingColumnType,proto3" json:"sharding_column_type,omitempty"`
	// redirects all traffic to another keyspace. If set, shards is ignored.
	ServedFrom string `protobuf:"bytes,5,opt,name=served_from,json=servedFrom,proto3" json:"served_from,omitempty"`
	// number of replica tablets to instantiate. This includes the master tablet.
	ReplicaCount int32 `protobuf:"varint,6,opt,name=replica_count,json=replicaCount,proto3" json:"replica_count,omitempty"`
	// number of rdonly tablets to instantiate.
	RdonlyCount          int32    `protobuf:"varint,7,opt,name=rdonly_count,json=rdonlyCount,proto3" json:"rdonly_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Keyspace) Reset()         { *m = Keyspace{} }
func (m *Keyspace) String() string { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()    {}
func (*Keyspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_vttest_beaece7261b82562, []int{1}
}
func (m *Keyspace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyspace.Unmarshal(m, b)
}
func (m *Keyspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyspace.Marshal(b, m, deterministic)
}
func (dst *Keyspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyspace.Merge(dst, src)
}
func (m *Keyspace) XXX_Size() int {
	return xxx_messageInfo_Keyspace.Size(m)
}
func (m *Keyspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyspace.DiscardUnknown(m)
}

var xxx_messageInfo_Keyspace proto.InternalMessageInfo

func (m *Keyspace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Keyspace) GetShards() []*Shard {
	if m != nil {
		return m.Shards
	}
	return nil
}

func (m *Keyspace) GetShardingColumnName() string {
	if m != nil {
		return m.ShardingColumnName
	}
	return ""
}

func (m *Keyspace) GetShardingColumnType() string {
	if m != nil {
		return m.ShardingColumnType
	}
	return ""
}

func (m *Keyspace) GetServedFrom() string {
	if m != nil {
		return m.ServedFrom
	}
	return ""
}

func (m *Keyspace) GetReplicaCount() int32 {
	if m != nil {
		return m.ReplicaCount
	}
	return 0
}

func (m *Keyspace) GetRdonlyCount() int32 {
	if m != nil {
		return m.RdonlyCount
	}
	return 0
}

// VTTestTopology describes the keyspaces in the topology.
type VTTestTopology struct {
	// all keyspaces in the topology.
	Keyspaces []*Keyspace `protobuf:"bytes,1,rep,name=keyspaces,proto3" json:"keyspaces,omitempty"`
	// list of cells the keyspaces reside in. Vtgate is started in only the first cell.
	Cells                []string `protobuf:"bytes,2,rep,name=cells,proto3" json:"cells,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VTTestTopology) Reset()         { *m = VTTestTopology{} }
func (m *VTTestTopology) String() string { return proto.CompactTextString(m) }
func (*VTTestTopology) ProtoMessage()    {}
func (*VTTestTopology) Descriptor() ([]byte, []int) {
	return fileDescriptor_vttest_beaece7261b82562, []int{2}
}
func (m *VTTestTopology) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VTTestTopology.Unmarshal(m, b)
}
func (m *VTTestTopology) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VTTestTopology.Marshal(b, m, deterministic)
}
func (dst *VTTestTopology) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VTTestTopology.Merge(dst, src)
}
func (m *VTTestTopology) XXX_Size() int {
	return xxx_messageInfo_VTTestTopology.Size(m)
}
func (m *VTTestTopology) XXX_DiscardUnknown() {
	xxx_messageInfo_VTTestTopology.DiscardUnknown(m)
}

var xxx_messageInfo_VTTestTopology proto.InternalMessageInfo

func (m *VTTestTopology) GetKeyspaces() []*Keyspace {
	if m != nil {
		return m.Keyspaces
	}
	return nil
}

func (m *VTTestTopology) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

func init() {
	proto.RegisterType((*Shard)(nil), "vttest.Shard")
	proto.RegisterType((*Keyspace)(nil), "vttest.Keyspace")
	proto.RegisterType((*VTTestTopology)(nil), "vttest.VTTestTopology")
}

func init() { proto.RegisterFile("vttest.proto", fileDescriptor_vttest_beaece7261b82562) }

var fileDescriptor_vttest_beaece7261b82562 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xcb, 0x6a, 0xe3, 0x40,
	0x10, 0x44, 0xb6, 0xa5, 0x5d, 0xb7, 0x1f, 0x98, 0xc1, 0x87, 0xb9, 0xad, 0xd7, 0xc6, 0xa0, 0x93,
	0xb4, 0x6c, 0xfe, 0x20, 0x26, 0xb9, 0x04, 0x12, 0x50, 0x84, 0x0f, 0xb9, 0x08, 0x59, 0xea, 0x38,
	0x22, 0x92, 0x5a, 0xcc, 0x8c, 0x05, 0xfa, 0x8d, 0x7c, 0x71, 0x50, 0x8f, 0x4c, 0x2e, 0xbe, 0x55,
	0x57, 0xd5, 0x74, 0x35, 0x35, 0x30, 0x6f, 0x8d, 0x41, 0x6d, 0x82, 0x46, 0x91, 0x21, 0xe1, 0xd9,
	0x69, 0xfb, 0x00, 0xee, 0xeb, 0x47, 0xaa, 0x72, 0x21, 0x60, 0x52, 0xa7, 0x15, 0x4a, 0x67, 0xe3,
	0xf8, 0xd3, 0x88, 0xb1, 0xf0, 0x61, 0x95, 0x9f, 0x92, 0x1e, 0x26, 0xd4, 0xa2, 0x52, 0x45, 0x8e,
	0x72, 0xc4, 0xfa, 0x32, 0x3f, 0x3d, 0xa7, 0x15, 0xbe, 0x0c, 0xec, 0xf6, 0x6b, 0x04, 0xbf, 0x9f,
	0xb0, 0xd3, 0x4d, 0x9a, 0xe1, 0xcd, 0x55, 0x7b, 0xf0, 0x74, 0x9f, 0xa3, 0xe5, 0x68, 0x33, 0xf6,
	0x67, 0xff, 0x17, 0xc1, 0x70, 0x0e, 0xa7, 0x47, 0x83, 0x28, 0xfe, 0xc1, 0x9a, 0x51, 0x51, 0x9f,
	0x93, 0x8c, 0xca, 0x4b, 0x55, 0x73, 0xbc, 0x1c, 0xf3, 0x2a, 0x71, 0xd5, 0x0e, 0x2c, 0xf5, 0x17,
	0xdc, 0x7a, 0x61, 0xba, 0x06, 0xe5, 0xe4, 0xd6, 0x8b, 0xb8, 0x6b, 0x50, 0xfc, 0x81, 0x99, 0x46,
	0xd5, 0x62, 0x9e, 0xbc, 0x2b, 0xaa, 0xa4, 0xcb, 0x46, 0xb0, 0xd4, 0xa3, 0xa2, 0x4a, 0xec, 0x60,
	0xa1, 0xb0, 0x29, 0x8b, 0x2c, 0x4d, 0x32, 0xba, 0xd4, 0x46, 0x7a, 0x1b, 0xc7, 0x77, 0xa3, 0xf9,
	0x40, 0x1e, 0x7a, 0x4e, 0xfc, 0x85, 0xb9, 0xca, 0xa9, 0x2e, 0xbb, 0xc1, 0xf3, 0x8b, 0x3d, 0x33,
	0xcb, 0xb1, 0x65, 0x7b, 0x84, 0xe5, 0x31, 0x8e, 0x51, 0x9b, 0x98, 0x1a, 0x2a, 0xe9, 0xdc, 0x89,
	0x00, 0xa6, 0x9f, 0x43, 0x4b, 0x5a, 0x3a, 0x5c, 0xc4, 0xea, 0x5a, 0xc4, 0xb5, 0xbe, 0xe8, 0xc7,
	0x22, 0xd6, 0xe0, 0x66, 0x58, 0x96, 0xb6, 0xb4, 0x69, 0x64, 0x87, 0xfb, 0xfd, 0xdb, 0xae, 0x2d,
	0x0c, 0x6a, 0x1d, 0x14, 0x14, 0x5a, 0x14, 0x9e, 0x29, 0x6c, 0x4d, 0xc8, 0x7f, 0x1b, 0xda, 0x85,
	0x27, 0x8f, 0xa7, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x0f, 0xe6, 0xb4, 0xf9, 0x01,
	0x00, 0x00,
}
