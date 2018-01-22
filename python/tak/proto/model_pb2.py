# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: model.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='model.proto',
  package='tak.proto',
  syntax='proto3',
  serialized_pb=_b('\n\x0bmodel.proto\x12\ttak.proto\"Y\n\x08ModelDef\x12\x0c\n\x04size\x18\x01 \x01(\x05\x12\x0e\n\x06layers\x18\x02 \x01(\x05\x12\x0e\n\x06kernel\x18\x03 \x01(\x05\x12\x0f\n\x07\x66ilters\x18\x04 \x01(\x05\x12\x0e\n\x06hidden\x18\x05 \x01(\x05\x42\x04Z\x02pbb\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_MODELDEF = _descriptor.Descriptor(
  name='ModelDef',
  full_name='tak.proto.ModelDef',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='size', full_name='tak.proto.ModelDef.size', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='layers', full_name='tak.proto.ModelDef.layers', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='kernel', full_name='tak.proto.ModelDef.kernel', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='filters', full_name='tak.proto.ModelDef.filters', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='hidden', full_name='tak.proto.ModelDef.hidden', index=4,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
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
  serialized_start=26,
  serialized_end=115,
)

DESCRIPTOR.message_types_by_name['ModelDef'] = _MODELDEF

ModelDef = _reflection.GeneratedProtocolMessageType('ModelDef', (_message.Message,), dict(
  DESCRIPTOR = _MODELDEF,
  __module__ = 'model_pb2'
  # @@protoc_insertion_point(class_scope:tak.proto.ModelDef)
  ))
_sym_db.RegisterMessage(ModelDef)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\002pb'))
# @@protoc_insertion_point(module_scope)