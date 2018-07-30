# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/admin/launch_plan.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import interface_pb2 as flyteidl_dot_core_dot_interface__pb2
from flyteidl.admin import schedule_pb2 as flyteidl_dot_admin_dot_schedule__pb2
from flyteidl.admin import common_pb2 as flyteidl_dot_admin_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/admin/launch_plan.proto',
  package='admin',
  syntax='proto3',
  serialized_pb=_b('\n flyteidl/admin/launch_plan.proto\x12\x05\x61\x64min\x1a\x1d\x66lyteidl/core/interface.proto\x1a\x1d\x66lyteidl/admin/schedule.proto\x1a\x1b\x66lyteidl/admin/common.proto\"n\n\x17LaunchPlanCreateRequest\x12\x1d\n\x02id\x18\x01 \x01(\x0b\x32\x11.admin.Identifier\x12\x0f\n\x07version\x18\x02 \x01(\t\x12#\n\x04spec\x18\x03 \x01(\x0b\x32\x15.admin.LaunchPlanSpec\"n\n\nLaunchPlan\x12\x1d\n\x02id\x18\x01 \x01(\x0b\x32\x11.admin.Identifier\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x0b\n\x03urn\x18\x03 \x01(\t\x12#\n\x04spec\x18\x04 \x01(\x0b\x32\x15.admin.LaunchPlanSpec\"9\n\x0eLaunchPlanList\x12\'\n\x0claunch_plans\x18\x01 \x03(\x0b\x32\x11.admin.LaunchPlan\"\xb6\x01\n\x0eLaunchPlanSpec\x12\x14\n\x0cworkflow_urn\x18\x01 \x01(\t\x12\x32\n\x0f\x65ntity_metadata\x18\x02 \x01(\x0b\x32\x19.admin.LaunchPlanMetadata\x12(\n\x0e\x64\x65\x66\x61ult_inputs\x18\x03 \x03(\x0b\x32\x10.admin.Parameter\x12\x30\n\x0c\x66ixed_inputs\x18\x04 \x01(\x0b\x32\x1a.core.NamedValueCollection\"c\n\x12LaunchPlanMetadata\x12!\n\x08schedule\x18\x01 \x01(\x0b\x32\x0f.admin.Schedule\x12*\n\rnotifications\x18\x02 \x03(\x0b\x32\x13.admin.NotificationB3Z1github.com/lyft/flyteidl/gen/pb-go/flyteidl/adminb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_interface__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_schedule__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_common__pb2.DESCRIPTOR,])




_LAUNCHPLANCREATEREQUEST = _descriptor.Descriptor(
  name='LaunchPlanCreateRequest',
  full_name='admin.LaunchPlanCreateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='admin.LaunchPlanCreateRequest.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version', full_name='admin.LaunchPlanCreateRequest.version', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='spec', full_name='admin.LaunchPlanCreateRequest.spec', index=2,
      number=3, type=11, cpp_type=10, label=1,
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
  serialized_start=134,
  serialized_end=244,
)


_LAUNCHPLAN = _descriptor.Descriptor(
  name='LaunchPlan',
  full_name='admin.LaunchPlan',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='admin.LaunchPlan.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version', full_name='admin.LaunchPlan.version', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='urn', full_name='admin.LaunchPlan.urn', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='spec', full_name='admin.LaunchPlan.spec', index=3,
      number=4, type=11, cpp_type=10, label=1,
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
  serialized_start=246,
  serialized_end=356,
)


_LAUNCHPLANLIST = _descriptor.Descriptor(
  name='LaunchPlanList',
  full_name='admin.LaunchPlanList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='launch_plans', full_name='admin.LaunchPlanList.launch_plans', index=0,
      number=1, type=11, cpp_type=10, label=3,
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
  serialized_start=358,
  serialized_end=415,
)


_LAUNCHPLANSPEC = _descriptor.Descriptor(
  name='LaunchPlanSpec',
  full_name='admin.LaunchPlanSpec',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='workflow_urn', full_name='admin.LaunchPlanSpec.workflow_urn', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='entity_metadata', full_name='admin.LaunchPlanSpec.entity_metadata', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_inputs', full_name='admin.LaunchPlanSpec.default_inputs', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fixed_inputs', full_name='admin.LaunchPlanSpec.fixed_inputs', index=3,
      number=4, type=11, cpp_type=10, label=1,
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
  serialized_start=418,
  serialized_end=600,
)


_LAUNCHPLANMETADATA = _descriptor.Descriptor(
  name='LaunchPlanMetadata',
  full_name='admin.LaunchPlanMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='schedule', full_name='admin.LaunchPlanMetadata.schedule', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='notifications', full_name='admin.LaunchPlanMetadata.notifications', index=1,
      number=2, type=11, cpp_type=10, label=3,
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
  serialized_start=602,
  serialized_end=701,
)

_LAUNCHPLANCREATEREQUEST.fields_by_name['id'].message_type = flyteidl_dot_admin_dot_common__pb2._IDENTIFIER
_LAUNCHPLANCREATEREQUEST.fields_by_name['spec'].message_type = _LAUNCHPLANSPEC
_LAUNCHPLAN.fields_by_name['id'].message_type = flyteidl_dot_admin_dot_common__pb2._IDENTIFIER
_LAUNCHPLAN.fields_by_name['spec'].message_type = _LAUNCHPLANSPEC
_LAUNCHPLANLIST.fields_by_name['launch_plans'].message_type = _LAUNCHPLAN
_LAUNCHPLANSPEC.fields_by_name['entity_metadata'].message_type = _LAUNCHPLANMETADATA
_LAUNCHPLANSPEC.fields_by_name['default_inputs'].message_type = flyteidl_dot_admin_dot_common__pb2._PARAMETER
_LAUNCHPLANSPEC.fields_by_name['fixed_inputs'].message_type = flyteidl_dot_core_dot_interface__pb2._NAMEDVALUECOLLECTION
_LAUNCHPLANMETADATA.fields_by_name['schedule'].message_type = flyteidl_dot_admin_dot_schedule__pb2._SCHEDULE
_LAUNCHPLANMETADATA.fields_by_name['notifications'].message_type = flyteidl_dot_admin_dot_common__pb2._NOTIFICATION
DESCRIPTOR.message_types_by_name['LaunchPlanCreateRequest'] = _LAUNCHPLANCREATEREQUEST
DESCRIPTOR.message_types_by_name['LaunchPlan'] = _LAUNCHPLAN
DESCRIPTOR.message_types_by_name['LaunchPlanList'] = _LAUNCHPLANLIST
DESCRIPTOR.message_types_by_name['LaunchPlanSpec'] = _LAUNCHPLANSPEC
DESCRIPTOR.message_types_by_name['LaunchPlanMetadata'] = _LAUNCHPLANMETADATA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

LaunchPlanCreateRequest = _reflection.GeneratedProtocolMessageType('LaunchPlanCreateRequest', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANCREATEREQUEST,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:admin.LaunchPlanCreateRequest)
  ))
_sym_db.RegisterMessage(LaunchPlanCreateRequest)

LaunchPlan = _reflection.GeneratedProtocolMessageType('LaunchPlan', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLAN,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:admin.LaunchPlan)
  ))
_sym_db.RegisterMessage(LaunchPlan)

LaunchPlanList = _reflection.GeneratedProtocolMessageType('LaunchPlanList', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANLIST,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:admin.LaunchPlanList)
  ))
_sym_db.RegisterMessage(LaunchPlanList)

LaunchPlanSpec = _reflection.GeneratedProtocolMessageType('LaunchPlanSpec', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANSPEC,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:admin.LaunchPlanSpec)
  ))
_sym_db.RegisterMessage(LaunchPlanSpec)

LaunchPlanMetadata = _reflection.GeneratedProtocolMessageType('LaunchPlanMetadata', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANMETADATA,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:admin.LaunchPlanMetadata)
  ))
_sym_db.RegisterMessage(LaunchPlanMetadata)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z1github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin'))
# @@protoc_insertion_point(module_scope)
