# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: queryservice.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import query_pb2 as query__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='queryservice.proto',
  package='queryservice',
  syntax='proto3',
  serialized_pb=_b('\n\x12queryservice.proto\x12\x0cqueryservice\x1a\x0bquery.proto2\xa7\x0c\n\x05Query\x12:\n\x07\x45xecute\x12\x15.query.ExecuteRequest\x1a\x16.query.ExecuteResponse\"\x00\x12I\n\x0c\x45xecuteBatch\x12\x1a.query.ExecuteBatchRequest\x1a\x1b.query.ExecuteBatchResponse\"\x00\x12N\n\rStreamExecute\x12\x1b.query.StreamExecuteRequest\x1a\x1c.query.StreamExecuteResponse\"\x00\x30\x01\x12\x34\n\x05\x42\x65gin\x12\x13.query.BeginRequest\x1a\x14.query.BeginResponse\"\x00\x12\x37\n\x06\x43ommit\x12\x14.query.CommitRequest\x1a\x15.query.CommitResponse\"\x00\x12=\n\x08Rollback\x12\x16.query.RollbackRequest\x1a\x17.query.RollbackResponse\"\x00\x12:\n\x07Prepare\x12\x15.query.PrepareRequest\x1a\x16.query.PrepareResponse\"\x00\x12O\n\x0e\x43ommitPrepared\x12\x1c.query.CommitPreparedRequest\x1a\x1d.query.CommitPreparedResponse\"\x00\x12U\n\x10RollbackPrepared\x12\x1e.query.RollbackPreparedRequest\x1a\x1f.query.RollbackPreparedResponse\"\x00\x12X\n\x11\x43reateTransaction\x12\x1f.query.CreateTransactionRequest\x1a .query.CreateTransactionResponse\"\x00\x12\x46\n\x0bStartCommit\x12\x19.query.StartCommitRequest\x1a\x1a.query.StartCommitResponse\"\x00\x12\x46\n\x0bSetRollback\x12\x19.query.SetRollbackRequest\x1a\x1a.query.SetRollbackResponse\"\x00\x12^\n\x13\x43oncludeTransaction\x12!.query.ConcludeTransactionRequest\x1a\".query.ConcludeTransactionResponse\"\x00\x12R\n\x0fReadTransaction\x12\x1d.query.ReadTransactionRequest\x1a\x1e.query.ReadTransactionResponse\"\x00\x12I\n\x0c\x42\x65ginExecute\x12\x1a.query.BeginExecuteRequest\x1a\x1b.query.BeginExecuteResponse\"\x00\x12X\n\x11\x42\x65ginExecuteBatch\x12\x1f.query.BeginExecuteBatchRequest\x1a .query.BeginExecuteBatchResponse\"\x00\x12N\n\rMessageStream\x12\x1b.query.MessageStreamRequest\x1a\x1c.query.MessageStreamResponse\"\x00\x30\x01\x12\x43\n\nMessageAck\x12\x18.query.MessageAckRequest\x1a\x19.query.MessageAckResponse\"\x00\x12\x43\n\nSplitQuery\x12\x18.query.SplitQueryRequest\x1a\x19.query.SplitQueryResponse\"\x00\x12K\n\x0cStreamHealth\x12\x1a.query.StreamHealthRequest\x1a\x1b.query.StreamHealthResponse\"\x00\x30\x01\x12K\n\x0cUpdateStream\x12\x1a.query.UpdateStreamRequest\x1a\x1b.query.UpdateStreamResponse\"\x00\x30\x01\x42+Z)github.com/mushiyu/vitess/go/vt/proto/queryserviceb\x06proto3')
  ,
  dependencies=[query__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z)github.com/mushiyu/vitess/go/vt/proto/queryservice'))

_QUERY = _descriptor.ServiceDescriptor(
  name='Query',
  full_name='queryservice.Query',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=50,
  serialized_end=1625,
  methods=[
  _descriptor.MethodDescriptor(
    name='Execute',
    full_name='queryservice.Query.Execute',
    index=0,
    containing_service=None,
    input_type=query__pb2._EXECUTEREQUEST,
    output_type=query__pb2._EXECUTERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ExecuteBatch',
    full_name='queryservice.Query.ExecuteBatch',
    index=1,
    containing_service=None,
    input_type=query__pb2._EXECUTEBATCHREQUEST,
    output_type=query__pb2._EXECUTEBATCHRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='StreamExecute',
    full_name='queryservice.Query.StreamExecute',
    index=2,
    containing_service=None,
    input_type=query__pb2._STREAMEXECUTEREQUEST,
    output_type=query__pb2._STREAMEXECUTERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Begin',
    full_name='queryservice.Query.Begin',
    index=3,
    containing_service=None,
    input_type=query__pb2._BEGINREQUEST,
    output_type=query__pb2._BEGINRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Commit',
    full_name='queryservice.Query.Commit',
    index=4,
    containing_service=None,
    input_type=query__pb2._COMMITREQUEST,
    output_type=query__pb2._COMMITRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Rollback',
    full_name='queryservice.Query.Rollback',
    index=5,
    containing_service=None,
    input_type=query__pb2._ROLLBACKREQUEST,
    output_type=query__pb2._ROLLBACKRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Prepare',
    full_name='queryservice.Query.Prepare',
    index=6,
    containing_service=None,
    input_type=query__pb2._PREPAREREQUEST,
    output_type=query__pb2._PREPARERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='CommitPrepared',
    full_name='queryservice.Query.CommitPrepared',
    index=7,
    containing_service=None,
    input_type=query__pb2._COMMITPREPAREDREQUEST,
    output_type=query__pb2._COMMITPREPAREDRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RollbackPrepared',
    full_name='queryservice.Query.RollbackPrepared',
    index=8,
    containing_service=None,
    input_type=query__pb2._ROLLBACKPREPAREDREQUEST,
    output_type=query__pb2._ROLLBACKPREPAREDRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='CreateTransaction',
    full_name='queryservice.Query.CreateTransaction',
    index=9,
    containing_service=None,
    input_type=query__pb2._CREATETRANSACTIONREQUEST,
    output_type=query__pb2._CREATETRANSACTIONRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='StartCommit',
    full_name='queryservice.Query.StartCommit',
    index=10,
    containing_service=None,
    input_type=query__pb2._STARTCOMMITREQUEST,
    output_type=query__pb2._STARTCOMMITRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetRollback',
    full_name='queryservice.Query.SetRollback',
    index=11,
    containing_service=None,
    input_type=query__pb2._SETROLLBACKREQUEST,
    output_type=query__pb2._SETROLLBACKRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ConcludeTransaction',
    full_name='queryservice.Query.ConcludeTransaction',
    index=12,
    containing_service=None,
    input_type=query__pb2._CONCLUDETRANSACTIONREQUEST,
    output_type=query__pb2._CONCLUDETRANSACTIONRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ReadTransaction',
    full_name='queryservice.Query.ReadTransaction',
    index=13,
    containing_service=None,
    input_type=query__pb2._READTRANSACTIONREQUEST,
    output_type=query__pb2._READTRANSACTIONRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='BeginExecute',
    full_name='queryservice.Query.BeginExecute',
    index=14,
    containing_service=None,
    input_type=query__pb2._BEGINEXECUTEREQUEST,
    output_type=query__pb2._BEGINEXECUTERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='BeginExecuteBatch',
    full_name='queryservice.Query.BeginExecuteBatch',
    index=15,
    containing_service=None,
    input_type=query__pb2._BEGINEXECUTEBATCHREQUEST,
    output_type=query__pb2._BEGINEXECUTEBATCHRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='MessageStream',
    full_name='queryservice.Query.MessageStream',
    index=16,
    containing_service=None,
    input_type=query__pb2._MESSAGESTREAMREQUEST,
    output_type=query__pb2._MESSAGESTREAMRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='MessageAck',
    full_name='queryservice.Query.MessageAck',
    index=17,
    containing_service=None,
    input_type=query__pb2._MESSAGEACKREQUEST,
    output_type=query__pb2._MESSAGEACKRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SplitQuery',
    full_name='queryservice.Query.SplitQuery',
    index=18,
    containing_service=None,
    input_type=query__pb2._SPLITQUERYREQUEST,
    output_type=query__pb2._SPLITQUERYRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='StreamHealth',
    full_name='queryservice.Query.StreamHealth',
    index=19,
    containing_service=None,
    input_type=query__pb2._STREAMHEALTHREQUEST,
    output_type=query__pb2._STREAMHEALTHRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateStream',
    full_name='queryservice.Query.UpdateStream',
    index=20,
    containing_service=None,
    input_type=query__pb2._UPDATESTREAMREQUEST,
    output_type=query__pb2._UPDATESTREAMRESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_QUERY)

DESCRIPTOR.services_by_name['Query'] = _QUERY

# @@protoc_insertion_point(module_scope)
