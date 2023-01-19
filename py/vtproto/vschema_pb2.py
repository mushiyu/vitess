# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vschema.proto

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
  name='vschema.proto',
  package='vschema',
  syntax='proto3',
  serialized_pb=_b('\n\rvschema.proto\x12\x07vschema\x1a\x0bquery.proto\"\xfe\x01\n\x08Keyspace\x12\x0f\n\x07sharded\x18\x01 \x01(\x08\x12\x31\n\x08vindexes\x18\x02 \x03(\x0b\x32\x1f.vschema.Keyspace.VindexesEntry\x12-\n\x06tables\x18\x03 \x03(\x0b\x32\x1d.vschema.Keyspace.TablesEntry\x1a@\n\rVindexesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x1e\n\x05value\x18\x02 \x01(\x0b\x32\x0f.vschema.Vindex:\x02\x38\x01\x1a=\n\x0bTablesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x1d\n\x05value\x18\x02 \x01(\x0b\x32\x0e.vschema.Table:\x02\x38\x01\"\x81\x01\n\x06Vindex\x12\x0c\n\x04type\x18\x01 \x01(\t\x12+\n\x06params\x18\x02 \x03(\x0b\x32\x1b.vschema.Vindex.ParamsEntry\x12\r\n\x05owner\x18\x03 \x01(\t\x1a-\n\x0bParamsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xca\x01\n\x05Table\x12\x0c\n\x04type\x18\x01 \x01(\t\x12.\n\x0f\x63olumn_vindexes\x18\x02 \x03(\x0b\x32\x15.vschema.ColumnVindex\x12.\n\x0e\x61uto_increment\x18\x03 \x01(\x0b\x32\x16.vschema.AutoIncrement\x12 \n\x07\x63olumns\x18\x04 \x03(\x0b\x32\x0f.vschema.Column\x12\x0e\n\x06pinned\x18\x05 \x01(\t\x12!\n\x19\x63olumn_list_authoritative\x18\x06 \x01(\x08\"=\n\x0c\x43olumnVindex\x12\x0e\n\x06\x63olumn\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0f\n\x07\x63olumns\x18\x03 \x03(\t\"1\n\rAutoIncrement\x12\x0e\n\x06\x63olumn\x18\x01 \x01(\t\x12\x10\n\x08sequence\x18\x02 \x01(\t\"1\n\x06\x43olumn\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x19\n\x04type\x18\x02 \x01(\x0e\x32\x0b.query.Type\"\x88\x01\n\nSrvVSchema\x12\x35\n\tkeyspaces\x18\x01 \x03(\x0b\x32\".vschema.SrvVSchema.KeyspacesEntry\x1a\x43\n\x0eKeyspacesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12 \n\x05value\x18\x02 \x01(\x0b\x32\x11.vschema.Keyspace:\x02\x38\x01\x42&Z$github.com/mushiyu/vitess/go/vt/proto/vschemab\x06proto3')
  ,
  dependencies=[query__pb2.DESCRIPTOR,])




_KEYSPACE_VINDEXESENTRY = _descriptor.Descriptor(
  name='VindexesEntry',
  full_name='vschema.Keyspace.VindexesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='vschema.Keyspace.VindexesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='vschema.Keyspace.VindexesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=167,
  serialized_end=231,
)

_KEYSPACE_TABLESENTRY = _descriptor.Descriptor(
  name='TablesEntry',
  full_name='vschema.Keyspace.TablesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='vschema.Keyspace.TablesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='vschema.Keyspace.TablesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=233,
  serialized_end=294,
)

_KEYSPACE = _descriptor.Descriptor(
  name='Keyspace',
  full_name='vschema.Keyspace',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='sharded', full_name='vschema.Keyspace.sharded', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vindexes', full_name='vschema.Keyspace.vindexes', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tables', full_name='vschema.Keyspace.tables', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_KEYSPACE_VINDEXESENTRY, _KEYSPACE_TABLESENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=40,
  serialized_end=294,
)


_VINDEX_PARAMSENTRY = _descriptor.Descriptor(
  name='ParamsEntry',
  full_name='vschema.Vindex.ParamsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='vschema.Vindex.ParamsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='vschema.Vindex.ParamsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=381,
  serialized_end=426,
)

_VINDEX = _descriptor.Descriptor(
  name='Vindex',
  full_name='vschema.Vindex',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='vschema.Vindex.type', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='params', full_name='vschema.Vindex.params', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='owner', full_name='vschema.Vindex.owner', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_VINDEX_PARAMSENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=297,
  serialized_end=426,
)


_TABLE = _descriptor.Descriptor(
  name='Table',
  full_name='vschema.Table',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='vschema.Table.type', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='column_vindexes', full_name='vschema.Table.column_vindexes', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='auto_increment', full_name='vschema.Table.auto_increment', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='columns', full_name='vschema.Table.columns', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pinned', full_name='vschema.Table.pinned', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='column_list_authoritative', full_name='vschema.Table.column_list_authoritative', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=429,
  serialized_end=631,
)


_COLUMNVINDEX = _descriptor.Descriptor(
  name='ColumnVindex',
  full_name='vschema.ColumnVindex',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='column', full_name='vschema.ColumnVindex.column', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='vschema.ColumnVindex.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='columns', full_name='vschema.ColumnVindex.columns', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=633,
  serialized_end=694,
)


_AUTOINCREMENT = _descriptor.Descriptor(
  name='AutoIncrement',
  full_name='vschema.AutoIncrement',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='column', full_name='vschema.AutoIncrement.column', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sequence', full_name='vschema.AutoIncrement.sequence', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=696,
  serialized_end=745,
)


_COLUMN = _descriptor.Descriptor(
  name='Column',
  full_name='vschema.Column',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='vschema.Column.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='vschema.Column.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=747,
  serialized_end=796,
)


_SRVVSCHEMA_KEYSPACESENTRY = _descriptor.Descriptor(
  name='KeyspacesEntry',
  full_name='vschema.SrvVSchema.KeyspacesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='vschema.SrvVSchema.KeyspacesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='vschema.SrvVSchema.KeyspacesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=868,
  serialized_end=935,
)

_SRVVSCHEMA = _descriptor.Descriptor(
  name='SrvVSchema',
  full_name='vschema.SrvVSchema',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='keyspaces', full_name='vschema.SrvVSchema.keyspaces', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SRVVSCHEMA_KEYSPACESENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=799,
  serialized_end=935,
)

_KEYSPACE_VINDEXESENTRY.fields_by_name['value'].message_type = _VINDEX
_KEYSPACE_VINDEXESENTRY.containing_type = _KEYSPACE
_KEYSPACE_TABLESENTRY.fields_by_name['value'].message_type = _TABLE
_KEYSPACE_TABLESENTRY.containing_type = _KEYSPACE
_KEYSPACE.fields_by_name['vindexes'].message_type = _KEYSPACE_VINDEXESENTRY
_KEYSPACE.fields_by_name['tables'].message_type = _KEYSPACE_TABLESENTRY
_VINDEX_PARAMSENTRY.containing_type = _VINDEX
_VINDEX.fields_by_name['params'].message_type = _VINDEX_PARAMSENTRY
_TABLE.fields_by_name['column_vindexes'].message_type = _COLUMNVINDEX
_TABLE.fields_by_name['auto_increment'].message_type = _AUTOINCREMENT
_TABLE.fields_by_name['columns'].message_type = _COLUMN
_COLUMN.fields_by_name['type'].enum_type = query__pb2._TYPE
_SRVVSCHEMA_KEYSPACESENTRY.fields_by_name['value'].message_type = _KEYSPACE
_SRVVSCHEMA_KEYSPACESENTRY.containing_type = _SRVVSCHEMA
_SRVVSCHEMA.fields_by_name['keyspaces'].message_type = _SRVVSCHEMA_KEYSPACESENTRY
DESCRIPTOR.message_types_by_name['Keyspace'] = _KEYSPACE
DESCRIPTOR.message_types_by_name['Vindex'] = _VINDEX
DESCRIPTOR.message_types_by_name['Table'] = _TABLE
DESCRIPTOR.message_types_by_name['ColumnVindex'] = _COLUMNVINDEX
DESCRIPTOR.message_types_by_name['AutoIncrement'] = _AUTOINCREMENT
DESCRIPTOR.message_types_by_name['Column'] = _COLUMN
DESCRIPTOR.message_types_by_name['SrvVSchema'] = _SRVVSCHEMA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Keyspace = _reflection.GeneratedProtocolMessageType('Keyspace', (_message.Message,), dict(

  VindexesEntry = _reflection.GeneratedProtocolMessageType('VindexesEntry', (_message.Message,), dict(
    DESCRIPTOR = _KEYSPACE_VINDEXESENTRY,
    __module__ = 'vschema_pb2'
    # @@protoc_insertion_point(class_scope:vschema.Keyspace.VindexesEntry)
    ))
  ,

  TablesEntry = _reflection.GeneratedProtocolMessageType('TablesEntry', (_message.Message,), dict(
    DESCRIPTOR = _KEYSPACE_TABLESENTRY,
    __module__ = 'vschema_pb2'
    # @@protoc_insertion_point(class_scope:vschema.Keyspace.TablesEntry)
    ))
  ,
  DESCRIPTOR = _KEYSPACE,
  __module__ = 'vschema_pb2'
  # @@protoc_insertion_point(class_scope:vschema.Keyspace)
  ))
_sym_db.RegisterMessage(Keyspace)
_sym_db.RegisterMessage(Keyspace.VindexesEntry)
_sym_db.RegisterMessage(Keyspace.TablesEntry)

Vindex = _reflection.GeneratedProtocolMessageType('Vindex', (_message.Message,), dict(

  ParamsEntry = _reflection.GeneratedProtocolMessageType('ParamsEntry', (_message.Message,), dict(
    DESCRIPTOR = _VINDEX_PARAMSENTRY,
    __module__ = 'vschema_pb2'
    # @@protoc_insertion_point(class_scope:vschema.Vindex.ParamsEntry)
    ))
  ,
  DESCRIPTOR = _VINDEX,
  __module__ = 'vschema_pb2'
  # @@protoc_insertion_point(class_scope:vschema.Vindex)
  ))
_sym_db.RegisterMessage(Vindex)
_sym_db.RegisterMessage(Vindex.ParamsEntry)

Table = _reflection.GeneratedProtocolMessageType('Table', (_message.Message,), dict(
  DESCRIPTOR = _TABLE,
  __module__ = 'vschema_pb2'
  # @@protoc_insertion_point(class_scope:vschema.Table)
  ))
_sym_db.RegisterMessage(Table)

ColumnVindex = _reflection.GeneratedProtocolMessageType('ColumnVindex', (_message.Message,), dict(
  DESCRIPTOR = _COLUMNVINDEX,
  __module__ = 'vschema_pb2'
  # @@protoc_insertion_point(class_scope:vschema.ColumnVindex)
  ))
_sym_db.RegisterMessage(ColumnVindex)

AutoIncrement = _reflection.GeneratedProtocolMessageType('AutoIncrement', (_message.Message,), dict(
  DESCRIPTOR = _AUTOINCREMENT,
  __module__ = 'vschema_pb2'
  # @@protoc_insertion_point(class_scope:vschema.AutoIncrement)
  ))
_sym_db.RegisterMessage(AutoIncrement)

Column = _reflection.GeneratedProtocolMessageType('Column', (_message.Message,), dict(
  DESCRIPTOR = _COLUMN,
  __module__ = 'vschema_pb2'
  # @@protoc_insertion_point(class_scope:vschema.Column)
  ))
_sym_db.RegisterMessage(Column)

SrvVSchema = _reflection.GeneratedProtocolMessageType('SrvVSchema', (_message.Message,), dict(

  KeyspacesEntry = _reflection.GeneratedProtocolMessageType('KeyspacesEntry', (_message.Message,), dict(
    DESCRIPTOR = _SRVVSCHEMA_KEYSPACESENTRY,
    __module__ = 'vschema_pb2'
    # @@protoc_insertion_point(class_scope:vschema.SrvVSchema.KeyspacesEntry)
    ))
  ,
  DESCRIPTOR = _SRVVSCHEMA,
  __module__ = 'vschema_pb2'
  # @@protoc_insertion_point(class_scope:vschema.SrvVSchema)
  ))
_sym_db.RegisterMessage(SrvVSchema)
_sym_db.RegisterMessage(SrvVSchema.KeyspacesEntry)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z$github.com/mushiyu/vitess/go/vt/proto/vschema'))
_KEYSPACE_VINDEXESENTRY.has_options = True
_KEYSPACE_VINDEXESENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
_KEYSPACE_TABLESENTRY.has_options = True
_KEYSPACE_TABLESENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
_VINDEX_PARAMSENTRY.has_options = True
_VINDEX_PARAMSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
_SRVVSCHEMA_KEYSPACESENTRY.has_options = True
_SRVVSCHEMA_KEYSPACESENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
# @@protoc_insertion_point(module_scope)
