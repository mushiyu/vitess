# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: tabletmanagerservice.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import tabletmanagerdata_pb2 as tabletmanagerdata__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='tabletmanagerservice.proto',
  package='tabletmanagerservice',
  syntax='proto3',
  serialized_pb=_b('\n\x1atabletmanagerservice.proto\x12\x14tabletmanagerservice\x1a\x17tabletmanagerdata.proto2\x95!\n\rTabletManager\x12I\n\x04Ping\x12\x1e.tabletmanagerdata.PingRequest\x1a\x1f.tabletmanagerdata.PingResponse\"\x00\x12L\n\x05Sleep\x12\x1f.tabletmanagerdata.SleepRequest\x1a .tabletmanagerdata.SleepResponse\"\x00\x12^\n\x0b\x45xecuteHook\x12%.tabletmanagerdata.ExecuteHookRequest\x1a&.tabletmanagerdata.ExecuteHookResponse\"\x00\x12X\n\tGetSchema\x12#.tabletmanagerdata.GetSchemaRequest\x1a$.tabletmanagerdata.GetSchemaResponse\"\x00\x12g\n\x0eGetPermissions\x12(.tabletmanagerdata.GetPermissionsRequest\x1a).tabletmanagerdata.GetPermissionsResponse\"\x00\x12^\n\x0bSetReadOnly\x12%.tabletmanagerdata.SetReadOnlyRequest\x1a&.tabletmanagerdata.SetReadOnlyResponse\"\x00\x12\x61\n\x0cSetReadWrite\x12&.tabletmanagerdata.SetReadWriteRequest\x1a\'.tabletmanagerdata.SetReadWriteResponse\"\x00\x12[\n\nChangeType\x12$.tabletmanagerdata.ChangeTypeRequest\x1a%.tabletmanagerdata.ChangeTypeResponse\"\x00\x12\x61\n\x0cRefreshState\x12&.tabletmanagerdata.RefreshStateRequest\x1a\'.tabletmanagerdata.RefreshStateResponse\"\x00\x12g\n\x0eRunHealthCheck\x12(.tabletmanagerdata.RunHealthCheckRequest\x1a).tabletmanagerdata.RunHealthCheckResponse\"\x00\x12p\n\x11IgnoreHealthError\x12+.tabletmanagerdata.IgnoreHealthErrorRequest\x1a,.tabletmanagerdata.IgnoreHealthErrorResponse\"\x00\x12\x61\n\x0cReloadSchema\x12&.tabletmanagerdata.ReloadSchemaRequest\x1a\'.tabletmanagerdata.ReloadSchemaResponse\"\x00\x12j\n\x0fPreflightSchema\x12).tabletmanagerdata.PreflightSchemaRequest\x1a*.tabletmanagerdata.PreflightSchemaResponse\"\x00\x12^\n\x0b\x41pplySchema\x12%.tabletmanagerdata.ApplySchemaRequest\x1a&.tabletmanagerdata.ApplySchemaResponse\"\x00\x12p\n\x11\x45xecuteFetchAsDba\x12+.tabletmanagerdata.ExecuteFetchAsDbaRequest\x1a,.tabletmanagerdata.ExecuteFetchAsDbaResponse\"\x00\x12\x7f\n\x16\x45xecuteFetchAsAllPrivs\x12\x30.tabletmanagerdata.ExecuteFetchAsAllPrivsRequest\x1a\x31.tabletmanagerdata.ExecuteFetchAsAllPrivsResponse\"\x00\x12p\n\x11\x45xecuteFetchAsApp\x12+.tabletmanagerdata.ExecuteFetchAsAppRequest\x1a,.tabletmanagerdata.ExecuteFetchAsAppResponse\"\x00\x12^\n\x0bSlaveStatus\x12%.tabletmanagerdata.SlaveStatusRequest\x1a&.tabletmanagerdata.SlaveStatusResponse\"\x00\x12g\n\x0eMasterPosition\x12(.tabletmanagerdata.MasterPositionRequest\x1a).tabletmanagerdata.MasterPositionResponse\"\x00\x12X\n\tStopSlave\x12#.tabletmanagerdata.StopSlaveRequest\x1a$.tabletmanagerdata.StopSlaveResponse\"\x00\x12m\n\x10StopSlaveMinimum\x12*.tabletmanagerdata.StopSlaveMinimumRequest\x1a+.tabletmanagerdata.StopSlaveMinimumResponse\"\x00\x12[\n\nStartSlave\x12$.tabletmanagerdata.StartSlaveRequest\x1a%.tabletmanagerdata.StartSlaveResponse\"\x00\x12\x8b\x01\n\x1aTabletExternallyReparented\x12\x34.tabletmanagerdata.TabletExternallyReparentedRequest\x1a\x35.tabletmanagerdata.TabletExternallyReparentedResponse\"\x00\x12\x82\x01\n\x17TabletExternallyElected\x12\x31.tabletmanagerdata.TabletExternallyElectedRequest\x1a\x32.tabletmanagerdata.TabletExternallyElectedResponse\"\x00\x12X\n\tGetSlaves\x12#.tabletmanagerdata.GetSlavesRequest\x1a$.tabletmanagerdata.GetSlavesResponse\"\x00\x12m\n\x10VReplicationExec\x12*.tabletmanagerdata.VReplicationExecRequest\x1a+.tabletmanagerdata.VReplicationExecResponse\"\x00\x12\x7f\n\x16VReplicationWaitForPos\x12\x30.tabletmanagerdata.VReplicationWaitForPosRequest\x1a\x31.tabletmanagerdata.VReplicationWaitForPosResponse\"\x00\x12m\n\x10ResetReplication\x12*.tabletmanagerdata.ResetReplicationRequest\x1a+.tabletmanagerdata.ResetReplicationResponse\"\x00\x12[\n\nInitMaster\x12$.tabletmanagerdata.InitMasterRequest\x1a%.tabletmanagerdata.InitMasterResponse\"\x00\x12\x82\x01\n\x17PopulateReparentJournal\x12\x31.tabletmanagerdata.PopulateReparentJournalRequest\x1a\x32.tabletmanagerdata.PopulateReparentJournalResponse\"\x00\x12X\n\tInitSlave\x12#.tabletmanagerdata.InitSlaveRequest\x1a$.tabletmanagerdata.InitSlaveResponse\"\x00\x12\x61\n\x0c\x44\x65moteMaster\x12&.tabletmanagerdata.DemoteMasterRequest\x1a\'.tabletmanagerdata.DemoteMasterResponse\"\x00\x12\x85\x01\n\x18PromoteSlaveWhenCaughtUp\x12\x32.tabletmanagerdata.PromoteSlaveWhenCaughtUpRequest\x1a\x33.tabletmanagerdata.PromoteSlaveWhenCaughtUpResponse\"\x00\x12m\n\x10SlaveWasPromoted\x12*.tabletmanagerdata.SlaveWasPromotedRequest\x1a+.tabletmanagerdata.SlaveWasPromotedResponse\"\x00\x12X\n\tSetMaster\x12#.tabletmanagerdata.SetMasterRequest\x1a$.tabletmanagerdata.SetMasterResponse\"\x00\x12p\n\x11SlaveWasRestarted\x12+.tabletmanagerdata.SlaveWasRestartedRequest\x1a,.tabletmanagerdata.SlaveWasRestartedResponse\"\x00\x12\x8e\x01\n\x1bStopReplicationAndGetStatus\x12\x35.tabletmanagerdata.StopReplicationAndGetStatusRequest\x1a\x36.tabletmanagerdata.StopReplicationAndGetStatusResponse\"\x00\x12\x61\n\x0cPromoteSlave\x12&.tabletmanagerdata.PromoteSlaveRequest\x1a\'.tabletmanagerdata.PromoteSlaveResponse\"\x00\x12Q\n\x06\x42\x61\x63kup\x12 .tabletmanagerdata.BackupRequest\x1a!.tabletmanagerdata.BackupResponse\"\x00\x30\x01\x12r\n\x11RestoreFromBackup\x12+.tabletmanagerdata.RestoreFromBackupRequest\x1a,.tabletmanagerdata.RestoreFromBackupResponse\"\x00\x30\x01\x42\x33Z1github.com/mushiyu/vitess/go/vt/proto/tabletmanagerserviceb\x06proto3')
  ,
  dependencies=[tabletmanagerdata__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z1github.com/mushiyu/vitess/go/vt/proto/tabletmanagerservice'))

_TABLETMANAGER = _descriptor.ServiceDescriptor(
  name='TabletManager',
  full_name='tabletmanagerservice.TabletManager',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=78,
  serialized_end=4323,
  methods=[
  _descriptor.MethodDescriptor(
    name='Ping',
    full_name='tabletmanagerservice.TabletManager.Ping',
    index=0,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._PINGREQUEST,
    output_type=tabletmanagerdata__pb2._PINGRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Sleep',
    full_name='tabletmanagerservice.TabletManager.Sleep',
    index=1,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._SLEEPREQUEST,
    output_type=tabletmanagerdata__pb2._SLEEPRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ExecuteHook',
    full_name='tabletmanagerservice.TabletManager.ExecuteHook',
    index=2,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._EXECUTEHOOKREQUEST,
    output_type=tabletmanagerdata__pb2._EXECUTEHOOKRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetSchema',
    full_name='tabletmanagerservice.TabletManager.GetSchema',
    index=3,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._GETSCHEMAREQUEST,
    output_type=tabletmanagerdata__pb2._GETSCHEMARESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetPermissions',
    full_name='tabletmanagerservice.TabletManager.GetPermissions',
    index=4,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._GETPERMISSIONSREQUEST,
    output_type=tabletmanagerdata__pb2._GETPERMISSIONSRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetReadOnly',
    full_name='tabletmanagerservice.TabletManager.SetReadOnly',
    index=5,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._SETREADONLYREQUEST,
    output_type=tabletmanagerdata__pb2._SETREADONLYRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetReadWrite',
    full_name='tabletmanagerservice.TabletManager.SetReadWrite',
    index=6,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._SETREADWRITEREQUEST,
    output_type=tabletmanagerdata__pb2._SETREADWRITERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ChangeType',
    full_name='tabletmanagerservice.TabletManager.ChangeType',
    index=7,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._CHANGETYPEREQUEST,
    output_type=tabletmanagerdata__pb2._CHANGETYPERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RefreshState',
    full_name='tabletmanagerservice.TabletManager.RefreshState',
    index=8,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._REFRESHSTATEREQUEST,
    output_type=tabletmanagerdata__pb2._REFRESHSTATERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RunHealthCheck',
    full_name='tabletmanagerservice.TabletManager.RunHealthCheck',
    index=9,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._RUNHEALTHCHECKREQUEST,
    output_type=tabletmanagerdata__pb2._RUNHEALTHCHECKRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='IgnoreHealthError',
    full_name='tabletmanagerservice.TabletManager.IgnoreHealthError',
    index=10,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._IGNOREHEALTHERRORREQUEST,
    output_type=tabletmanagerdata__pb2._IGNOREHEALTHERRORRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ReloadSchema',
    full_name='tabletmanagerservice.TabletManager.ReloadSchema',
    index=11,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._RELOADSCHEMAREQUEST,
    output_type=tabletmanagerdata__pb2._RELOADSCHEMARESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PreflightSchema',
    full_name='tabletmanagerservice.TabletManager.PreflightSchema',
    index=12,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._PREFLIGHTSCHEMAREQUEST,
    output_type=tabletmanagerdata__pb2._PREFLIGHTSCHEMARESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ApplySchema',
    full_name='tabletmanagerservice.TabletManager.ApplySchema',
    index=13,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._APPLYSCHEMAREQUEST,
    output_type=tabletmanagerdata__pb2._APPLYSCHEMARESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ExecuteFetchAsDba',
    full_name='tabletmanagerservice.TabletManager.ExecuteFetchAsDba',
    index=14,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._EXECUTEFETCHASDBAREQUEST,
    output_type=tabletmanagerdata__pb2._EXECUTEFETCHASDBARESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ExecuteFetchAsAllPrivs',
    full_name='tabletmanagerservice.TabletManager.ExecuteFetchAsAllPrivs',
    index=15,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._EXECUTEFETCHASALLPRIVSREQUEST,
    output_type=tabletmanagerdata__pb2._EXECUTEFETCHASALLPRIVSRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ExecuteFetchAsApp',
    full_name='tabletmanagerservice.TabletManager.ExecuteFetchAsApp',
    index=16,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._EXECUTEFETCHASAPPREQUEST,
    output_type=tabletmanagerdata__pb2._EXECUTEFETCHASAPPRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SlaveStatus',
    full_name='tabletmanagerservice.TabletManager.SlaveStatus',
    index=17,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._SLAVESTATUSREQUEST,
    output_type=tabletmanagerdata__pb2._SLAVESTATUSRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='MasterPosition',
    full_name='tabletmanagerservice.TabletManager.MasterPosition',
    index=18,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._MASTERPOSITIONREQUEST,
    output_type=tabletmanagerdata__pb2._MASTERPOSITIONRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='StopSlave',
    full_name='tabletmanagerservice.TabletManager.StopSlave',
    index=19,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._STOPSLAVEREQUEST,
    output_type=tabletmanagerdata__pb2._STOPSLAVERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='StopSlaveMinimum',
    full_name='tabletmanagerservice.TabletManager.StopSlaveMinimum',
    index=20,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._STOPSLAVEMINIMUMREQUEST,
    output_type=tabletmanagerdata__pb2._STOPSLAVEMINIMUMRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='StartSlave',
    full_name='tabletmanagerservice.TabletManager.StartSlave',
    index=21,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._STARTSLAVEREQUEST,
    output_type=tabletmanagerdata__pb2._STARTSLAVERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='TabletExternallyReparented',
    full_name='tabletmanagerservice.TabletManager.TabletExternallyReparented',
    index=22,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._TABLETEXTERNALLYREPARENTEDREQUEST,
    output_type=tabletmanagerdata__pb2._TABLETEXTERNALLYREPARENTEDRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='TabletExternallyElected',
    full_name='tabletmanagerservice.TabletManager.TabletExternallyElected',
    index=23,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._TABLETEXTERNALLYELECTEDREQUEST,
    output_type=tabletmanagerdata__pb2._TABLETEXTERNALLYELECTEDRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetSlaves',
    full_name='tabletmanagerservice.TabletManager.GetSlaves',
    index=24,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._GETSLAVESREQUEST,
    output_type=tabletmanagerdata__pb2._GETSLAVESRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VReplicationExec',
    full_name='tabletmanagerservice.TabletManager.VReplicationExec',
    index=25,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._VREPLICATIONEXECREQUEST,
    output_type=tabletmanagerdata__pb2._VREPLICATIONEXECRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VReplicationWaitForPos',
    full_name='tabletmanagerservice.TabletManager.VReplicationWaitForPos',
    index=26,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._VREPLICATIONWAITFORPOSREQUEST,
    output_type=tabletmanagerdata__pb2._VREPLICATIONWAITFORPOSRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ResetReplication',
    full_name='tabletmanagerservice.TabletManager.ResetReplication',
    index=27,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._RESETREPLICATIONREQUEST,
    output_type=tabletmanagerdata__pb2._RESETREPLICATIONRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='InitMaster',
    full_name='tabletmanagerservice.TabletManager.InitMaster',
    index=28,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._INITMASTERREQUEST,
    output_type=tabletmanagerdata__pb2._INITMASTERRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PopulateReparentJournal',
    full_name='tabletmanagerservice.TabletManager.PopulateReparentJournal',
    index=29,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._POPULATEREPARENTJOURNALREQUEST,
    output_type=tabletmanagerdata__pb2._POPULATEREPARENTJOURNALRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='InitSlave',
    full_name='tabletmanagerservice.TabletManager.InitSlave',
    index=30,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._INITSLAVEREQUEST,
    output_type=tabletmanagerdata__pb2._INITSLAVERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='DemoteMaster',
    full_name='tabletmanagerservice.TabletManager.DemoteMaster',
    index=31,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._DEMOTEMASTERREQUEST,
    output_type=tabletmanagerdata__pb2._DEMOTEMASTERRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PromoteSlaveWhenCaughtUp',
    full_name='tabletmanagerservice.TabletManager.PromoteSlaveWhenCaughtUp',
    index=32,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._PROMOTESLAVEWHENCAUGHTUPREQUEST,
    output_type=tabletmanagerdata__pb2._PROMOTESLAVEWHENCAUGHTUPRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SlaveWasPromoted',
    full_name='tabletmanagerservice.TabletManager.SlaveWasPromoted',
    index=33,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._SLAVEWASPROMOTEDREQUEST,
    output_type=tabletmanagerdata__pb2._SLAVEWASPROMOTEDRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetMaster',
    full_name='tabletmanagerservice.TabletManager.SetMaster',
    index=34,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._SETMASTERREQUEST,
    output_type=tabletmanagerdata__pb2._SETMASTERRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SlaveWasRestarted',
    full_name='tabletmanagerservice.TabletManager.SlaveWasRestarted',
    index=35,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._SLAVEWASRESTARTEDREQUEST,
    output_type=tabletmanagerdata__pb2._SLAVEWASRESTARTEDRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='StopReplicationAndGetStatus',
    full_name='tabletmanagerservice.TabletManager.StopReplicationAndGetStatus',
    index=36,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._STOPREPLICATIONANDGETSTATUSREQUEST,
    output_type=tabletmanagerdata__pb2._STOPREPLICATIONANDGETSTATUSRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PromoteSlave',
    full_name='tabletmanagerservice.TabletManager.PromoteSlave',
    index=37,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._PROMOTESLAVEREQUEST,
    output_type=tabletmanagerdata__pb2._PROMOTESLAVERESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Backup',
    full_name='tabletmanagerservice.TabletManager.Backup',
    index=38,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._BACKUPREQUEST,
    output_type=tabletmanagerdata__pb2._BACKUPRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RestoreFromBackup',
    full_name='tabletmanagerservice.TabletManager.RestoreFromBackup',
    index=39,
    containing_service=None,
    input_type=tabletmanagerdata__pb2._RESTOREFROMBACKUPREQUEST,
    output_type=tabletmanagerdata__pb2._RESTOREFROMBACKUPRESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_TABLETMANAGER)

DESCRIPTOR.services_by_name['TabletManager'] = _TABLETMANAGER

# @@protoc_insertion_point(module_scope)
