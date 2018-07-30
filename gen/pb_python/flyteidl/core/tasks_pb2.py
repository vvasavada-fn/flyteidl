# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/core/tasks.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import interface_pb2 as flyteidl_dot_core_dot_interface__pb2
from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/core/tasks.proto',
  package='core',
  syntax='proto3',
  serialized_pb=_b('\n\x19\x66lyteidl/core/tasks.proto\x12\x04\x63ore\x1a\x1d\x66lyteidl/core/interface.proto\x1a\x1c\x66lyteidl/core/literals.proto\x1a\x1egoogle/protobuf/duration.proto\"\xff\x01\n\tResources\x12/\n\x08requests\x18\x01 \x03(\x0b\x32\x1d.core.Resources.ResourceEntry\x12-\n\x06limits\x18\x02 \x03(\x0b\x32\x1d.core.Resources.ResourceEntry\x1aJ\n\rResourceEntry\x12*\n\x04name\x18\x01 \x01(\x0e\x32\x1c.core.Resources.ResourceName\x12\r\n\x05value\x18\x02 \x01(\t\"F\n\x0cResourceName\x12\x0b\n\x07Unknown\x10\x00\x12\x07\n\x03\x43pu\x10\x01\x12\x07\n\x03Gpu\x10\x02\x12\n\n\x06Memory\x10\x03\x12\x0b\n\x07Storage\x10\x04\"\x8b\x01\n\x0fRuntimeMetadata\x12/\n\x04type\x18\x01 \x01(\x0e\x32!.core.RuntimeMetadata.RuntimeType\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x0e\n\x06\x66lavor\x18\x03 \x01(\t\"&\n\x0bRuntimeType\x12\x0c\n\x08\x46lyteSDK\x10\x00\x12\t\n\x05Other\x10\x01\"\x9e\x01\n\x0cTaskMetadata\x12\x14\n\x0c\x64iscoverable\x18\x01 \x01(\x08\x12&\n\x07runtime\x18\x02 \x01(\x0b\x32\x15.core.RuntimeMetadata\x12*\n\x07timeout\x18\x04 \x01(\x0b\x32\x19.google.protobuf.Duration\x12$\n\x07retries\x18\x05 \x01(\x0b\x32\x13.core.RetryStrategy\"\xdb\x01\n\x0cTaskTemplate\x12\n\n\x02id\x18\x01 \x01(\t\x12$\n\x08\x63\x61tegory\x18\x02 \x01(\x0e\x32\x12.core.TaskCategory\x12\x0c\n\x04type\x18\x03 \x01(\t\x12$\n\x08metadata\x18\x04 \x01(\x0b\x32\x12.core.TaskMetadata\x12\'\n\tinterface\x18\x05 \x01(\x0b\x32\x14.core.TypedInterface\x12\x0e\n\x06\x63ustom\x18\x06 \x01(\x0c\x12$\n\tcontainer\x18\x07 \x01(\x0b\x32\x0f.core.ContainerH\x00\x42\x06\n\x04task\"\xa2\x01\n\tContainer\x12\r\n\x05image\x18\x01 \x01(\t\x12\x0f\n\x07\x63ommand\x18\x02 \x03(\t\x12\x0c\n\x04\x61rgs\x18\x03 \x03(\t\x12\"\n\tresources\x18\x04 \x01(\x0b\x32\x0f.core.Resources\x12\x1f\n\x03\x65nv\x18\x05 \x03(\x0b\x32\x12.core.KeyValuePair\x12\"\n\x06\x63onfig\x18\x06 \x03(\x0b\x32\x12.core.KeyValuePair*5\n\x0cTaskCategory\x12\x12\n\x0eSingleStepTask\x10\x00\x12\x11\n\rMultiStepTask\x10\x01\x42\x32Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/coreb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_interface__pb2.DESCRIPTOR,flyteidl_dot_core_dot_literals__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,])

_TASKCATEGORY = _descriptor.EnumDescriptor(
  name='TaskCategory',
  full_name='core.TaskCategory',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SingleStepTask', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MultiStepTask', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1076,
  serialized_end=1129,
)
_sym_db.RegisterEnumDescriptor(_TASKCATEGORY)

TaskCategory = enum_type_wrapper.EnumTypeWrapper(_TASKCATEGORY)
SingleStepTask = 0
MultiStepTask = 1


_RESOURCES_RESOURCENAME = _descriptor.EnumDescriptor(
  name='ResourceName',
  full_name='core.Resources.ResourceName',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Unknown', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Cpu', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Gpu', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Memory', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Storage', index=4, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=314,
  serialized_end=384,
)
_sym_db.RegisterEnumDescriptor(_RESOURCES_RESOURCENAME)

_RUNTIMEMETADATA_RUNTIMETYPE = _descriptor.EnumDescriptor(
  name='RuntimeType',
  full_name='core.RuntimeMetadata.RuntimeType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='FlyteSDK', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Other', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=488,
  serialized_end=526,
)
_sym_db.RegisterEnumDescriptor(_RUNTIMEMETADATA_RUNTIMETYPE)


_RESOURCES_RESOURCEENTRY = _descriptor.Descriptor(
  name='ResourceEntry',
  full_name='core.Resources.ResourceEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='core.Resources.ResourceEntry.name', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='core.Resources.ResourceEntry.value', index=1,
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
  serialized_start=238,
  serialized_end=312,
)

_RESOURCES = _descriptor.Descriptor(
  name='Resources',
  full_name='core.Resources',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='requests', full_name='core.Resources.requests', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limits', full_name='core.Resources.limits', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_RESOURCES_RESOURCEENTRY, ],
  enum_types=[
    _RESOURCES_RESOURCENAME,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=129,
  serialized_end=384,
)


_RUNTIMEMETADATA = _descriptor.Descriptor(
  name='RuntimeMetadata',
  full_name='core.RuntimeMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='core.RuntimeMetadata.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version', full_name='core.RuntimeMetadata.version', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='flavor', full_name='core.RuntimeMetadata.flavor', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _RUNTIMEMETADATA_RUNTIMETYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=387,
  serialized_end=526,
)


_TASKMETADATA = _descriptor.Descriptor(
  name='TaskMetadata',
  full_name='core.TaskMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='discoverable', full_name='core.TaskMetadata.discoverable', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='runtime', full_name='core.TaskMetadata.runtime', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timeout', full_name='core.TaskMetadata.timeout', index=2,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='retries', full_name='core.TaskMetadata.retries', index=3,
      number=5, type=11, cpp_type=10, label=1,
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
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=529,
  serialized_end=687,
)


_TASKTEMPLATE = _descriptor.Descriptor(
  name='TaskTemplate',
  full_name='core.TaskTemplate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='core.TaskTemplate.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='category', full_name='core.TaskTemplate.category', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='core.TaskTemplate.type', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='core.TaskTemplate.metadata', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='interface', full_name='core.TaskTemplate.interface', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='custom', full_name='core.TaskTemplate.custom', index=5,
      number=6, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='container', full_name='core.TaskTemplate.container', index=6,
      number=7, type=11, cpp_type=10, label=1,
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
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='task', full_name='core.TaskTemplate.task',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=690,
  serialized_end=909,
)


_CONTAINER = _descriptor.Descriptor(
  name='Container',
  full_name='core.Container',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image', full_name='core.Container.image', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='command', full_name='core.Container.command', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='args', full_name='core.Container.args', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resources', full_name='core.Container.resources', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='env', full_name='core.Container.env', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='config', full_name='core.Container.config', index=5,
      number=6, type=11, cpp_type=10, label=3,
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
  serialized_start=912,
  serialized_end=1074,
)

_RESOURCES_RESOURCEENTRY.fields_by_name['name'].enum_type = _RESOURCES_RESOURCENAME
_RESOURCES_RESOURCEENTRY.containing_type = _RESOURCES
_RESOURCES.fields_by_name['requests'].message_type = _RESOURCES_RESOURCEENTRY
_RESOURCES.fields_by_name['limits'].message_type = _RESOURCES_RESOURCEENTRY
_RESOURCES_RESOURCENAME.containing_type = _RESOURCES
_RUNTIMEMETADATA.fields_by_name['type'].enum_type = _RUNTIMEMETADATA_RUNTIMETYPE
_RUNTIMEMETADATA_RUNTIMETYPE.containing_type = _RUNTIMEMETADATA
_TASKMETADATA.fields_by_name['runtime'].message_type = _RUNTIMEMETADATA
_TASKMETADATA.fields_by_name['timeout'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_TASKMETADATA.fields_by_name['retries'].message_type = flyteidl_dot_core_dot_literals__pb2._RETRYSTRATEGY
_TASKTEMPLATE.fields_by_name['category'].enum_type = _TASKCATEGORY
_TASKTEMPLATE.fields_by_name['metadata'].message_type = _TASKMETADATA
_TASKTEMPLATE.fields_by_name['interface'].message_type = flyteidl_dot_core_dot_interface__pb2._TYPEDINTERFACE
_TASKTEMPLATE.fields_by_name['container'].message_type = _CONTAINER
_TASKTEMPLATE.oneofs_by_name['task'].fields.append(
  _TASKTEMPLATE.fields_by_name['container'])
_TASKTEMPLATE.fields_by_name['container'].containing_oneof = _TASKTEMPLATE.oneofs_by_name['task']
_CONTAINER.fields_by_name['resources'].message_type = _RESOURCES
_CONTAINER.fields_by_name['env'].message_type = flyteidl_dot_core_dot_literals__pb2._KEYVALUEPAIR
_CONTAINER.fields_by_name['config'].message_type = flyteidl_dot_core_dot_literals__pb2._KEYVALUEPAIR
DESCRIPTOR.message_types_by_name['Resources'] = _RESOURCES
DESCRIPTOR.message_types_by_name['RuntimeMetadata'] = _RUNTIMEMETADATA
DESCRIPTOR.message_types_by_name['TaskMetadata'] = _TASKMETADATA
DESCRIPTOR.message_types_by_name['TaskTemplate'] = _TASKTEMPLATE
DESCRIPTOR.message_types_by_name['Container'] = _CONTAINER
DESCRIPTOR.enum_types_by_name['TaskCategory'] = _TASKCATEGORY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Resources = _reflection.GeneratedProtocolMessageType('Resources', (_message.Message,), dict(

  ResourceEntry = _reflection.GeneratedProtocolMessageType('ResourceEntry', (_message.Message,), dict(
    DESCRIPTOR = _RESOURCES_RESOURCEENTRY,
    __module__ = 'flyteidl.core.tasks_pb2'
    # @@protoc_insertion_point(class_scope:core.Resources.ResourceEntry)
    ))
  ,
  DESCRIPTOR = _RESOURCES,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:core.Resources)
  ))
_sym_db.RegisterMessage(Resources)
_sym_db.RegisterMessage(Resources.ResourceEntry)

RuntimeMetadata = _reflection.GeneratedProtocolMessageType('RuntimeMetadata', (_message.Message,), dict(
  DESCRIPTOR = _RUNTIMEMETADATA,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:core.RuntimeMetadata)
  ))
_sym_db.RegisterMessage(RuntimeMetadata)

TaskMetadata = _reflection.GeneratedProtocolMessageType('TaskMetadata', (_message.Message,), dict(
  DESCRIPTOR = _TASKMETADATA,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:core.TaskMetadata)
  ))
_sym_db.RegisterMessage(TaskMetadata)

TaskTemplate = _reflection.GeneratedProtocolMessageType('TaskTemplate', (_message.Message,), dict(
  DESCRIPTOR = _TASKTEMPLATE,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:core.TaskTemplate)
  ))
_sym_db.RegisterMessage(TaskTemplate)

Container = _reflection.GeneratedProtocolMessageType('Container', (_message.Message,), dict(
  DESCRIPTOR = _CONTAINER,
  __module__ = 'flyteidl.core.tasks_pb2'
  # @@protoc_insertion_point(class_scope:core.Container)
  ))
_sym_db.RegisterMessage(Container)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z0github.com/lyft/flyteidl/gen/pb-go/flyteidl/core'))
# @@protoc_insertion_point(module_scope)
