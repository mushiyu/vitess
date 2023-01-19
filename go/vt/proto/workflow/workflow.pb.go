// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workflow.proto

package workflow // import "github.com/mushiyu/vitess/go/vt/proto/workflow"

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

// WorkflowState describes the state of a workflow.
// This constant should match the Node object described in
// web/vtctld2/src/app/workflows/node.ts as it is exposed as JSON to
// the Angular 2 web app.
type WorkflowState int32

const (
	WorkflowState_NotStarted WorkflowState = 0
	WorkflowState_Running    WorkflowState = 1
	WorkflowState_Done       WorkflowState = 2
)

var WorkflowState_name = map[int32]string{
	0: "NotStarted",
	1: "Running",
	2: "Done",
}
var WorkflowState_value = map[string]int32{
	"NotStarted": 0,
	"Running":    1,
	"Done":       2,
}

func (x WorkflowState) String() string {
	return proto.EnumName(WorkflowState_name, int32(x))
}
func (WorkflowState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_workflow_cc5eebeb403313d8, []int{0}
}

type TaskState int32

const (
	TaskState_TaskNotStarted TaskState = 0
	TaskState_TaskRunning    TaskState = 1
	TaskState_TaskDone       TaskState = 2
)

var TaskState_name = map[int32]string{
	0: "TaskNotStarted",
	1: "TaskRunning",
	2: "TaskDone",
}
var TaskState_value = map[string]int32{
	"TaskNotStarted": 0,
	"TaskRunning":    1,
	"TaskDone":       2,
}

func (x TaskState) String() string {
	return proto.EnumName(TaskState_name, int32(x))
}
func (TaskState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_workflow_cc5eebeb403313d8, []int{1}
}

// Workflow is the persisted state of a long-running workflow.
type Workflow struct {
	// uuid is set when the workflow is created, and immutable after
	// that.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// factory_name is set with the name of the factory that created the
	// job (and can also restart it). It is set at creation time, and
	// immutable after that.
	FactoryName string `protobuf:"bytes,2,opt,name=factory_name,json=factoryName,proto3" json:"factory_name,omitempty"`
	// name is the display name of the workflow.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// state describes the state of the job. A job is created as
	// NotStarted, then the Workflow Manager picks it up and starts it,
	// switching it to Running (and populating 'start_time').  The
	// workflow can then fail over to a new Workflow Manager is
	// necessary, and still be in Running state.  When done, it goes to
	// Done, 'end_time' is populated, and 'error' is set if there was an
	// error.
	State WorkflowState `protobuf:"varint,4,opt,name=state,proto3,enum=workflow.WorkflowState" json:"state,omitempty"`
	// data is workflow-specific stored data. It is usually a binary
	// proto-encoded data structure. It can vary throughout the
	// execution of the workflow.  It will not change after the workflow
	// is Done.
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	// error is set if the job finished with an error. This field only
	// makes sense if 'state' is Done.
	Error string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	// start_time is set when the workflow manager starts a workflow for
	// the first time. This field only makes sense if 'state' is Running
	// or Done.
	StartTime int64 `protobuf:"varint,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// end_time is set when the workflow is finished.
	// This field only makes sense if 'state' is Done.
	EndTime int64 `protobuf:"varint,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// create_time is set when the workflow is created.
	CreateTime           int64    `protobuf:"varint,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Workflow) Reset()         { *m = Workflow{} }
func (m *Workflow) String() string { return proto.CompactTextString(m) }
func (*Workflow) ProtoMessage()    {}
func (*Workflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_cc5eebeb403313d8, []int{0}
}
func (m *Workflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workflow.Unmarshal(m, b)
}
func (m *Workflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workflow.Marshal(b, m, deterministic)
}
func (dst *Workflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workflow.Merge(dst, src)
}
func (m *Workflow) XXX_Size() int {
	return xxx_messageInfo_Workflow.Size(m)
}
func (m *Workflow) XXX_DiscardUnknown() {
	xxx_messageInfo_Workflow.DiscardUnknown(m)
}

var xxx_messageInfo_Workflow proto.InternalMessageInfo

func (m *Workflow) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Workflow) GetFactoryName() string {
	if m != nil {
		return m.FactoryName
	}
	return ""
}

func (m *Workflow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Workflow) GetState() WorkflowState {
	if m != nil {
		return m.State
	}
	return WorkflowState_NotStarted
}

func (m *Workflow) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Workflow) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Workflow) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Workflow) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Workflow) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type WorkflowCheckpoint struct {
	// code_version is used to detect incompabilities between the version of the
	// running workflow and the one which wrote the checkpoint. If they don't
	// match, the workflow must not continue. The author of workflow must update
	// this variable in their implementation when incompabilities are introduced.
	CodeVersion int32 `protobuf:"varint,1,opt,name=code_version,json=codeVersion,proto3" json:"code_version,omitempty"`
	// Task is the data structure that stores the execution status and the
	// attributes of a task.
	Tasks map[string]*Task `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// settings includes workflow specific data, e.g. the resharding workflow
	// would store the source shards and destination shards.
	Settings             map[string]string `protobuf:"bytes,3,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WorkflowCheckpoint) Reset()         { *m = WorkflowCheckpoint{} }
func (m *WorkflowCheckpoint) String() string { return proto.CompactTextString(m) }
func (*WorkflowCheckpoint) ProtoMessage()    {}
func (*WorkflowCheckpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_cc5eebeb403313d8, []int{1}
}
func (m *WorkflowCheckpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowCheckpoint.Unmarshal(m, b)
}
func (m *WorkflowCheckpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowCheckpoint.Marshal(b, m, deterministic)
}
func (dst *WorkflowCheckpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowCheckpoint.Merge(dst, src)
}
func (m *WorkflowCheckpoint) XXX_Size() int {
	return xxx_messageInfo_WorkflowCheckpoint.Size(m)
}
func (m *WorkflowCheckpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowCheckpoint.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowCheckpoint proto.InternalMessageInfo

func (m *WorkflowCheckpoint) GetCodeVersion() int32 {
	if m != nil {
		return m.CodeVersion
	}
	return 0
}

func (m *WorkflowCheckpoint) GetTasks() map[string]*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowCheckpoint) GetSettings() map[string]string {
	if m != nil {
		return m.Settings
	}
	return nil
}

type Task struct {
	Id    string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State TaskState `protobuf:"varint,2,opt,name=state,proto3,enum=workflow.TaskState" json:"state,omitempty"`
	// attributes includes the parameters the task needs.
	Attributes           map[string]string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Error                string            `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_cc5eebeb403313d8, []int{2}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetState() TaskState {
	if m != nil {
		return m.State
	}
	return TaskState_TaskNotStarted
}

func (m *Task) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Task) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Workflow)(nil), "workflow.Workflow")
	proto.RegisterType((*WorkflowCheckpoint)(nil), "workflow.WorkflowCheckpoint")
	proto.RegisterMapType((map[string]string)(nil), "workflow.WorkflowCheckpoint.SettingsEntry")
	proto.RegisterMapType((map[string]*Task)(nil), "workflow.WorkflowCheckpoint.TasksEntry")
	proto.RegisterType((*Task)(nil), "workflow.Task")
	proto.RegisterMapType((map[string]string)(nil), "workflow.Task.AttributesEntry")
	proto.RegisterEnum("workflow.WorkflowState", WorkflowState_name, WorkflowState_value)
	proto.RegisterEnum("workflow.TaskState", TaskState_name, TaskState_value)
}

func init() { proto.RegisterFile("workflow.proto", fileDescriptor_workflow_cc5eebeb403313d8) }

var fileDescriptor_workflow_cc5eebeb403313d8 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x6f, 0x8b, 0xd3, 0x4e,
	0x10, 0xfe, 0x25, 0x6d, 0xae, 0xe9, 0xa4, 0x97, 0x2b, 0xf3, 0x3b, 0x30, 0x16, 0xd4, 0x5a, 0x94,
	0xab, 0x05, 0x5b, 0xa8, 0x20, 0xa2, 0xdc, 0x81, 0x7f, 0xf1, 0xd5, 0xbd, 0x48, 0x0f, 0x05, 0xdf,
	0x94, 0xbd, 0x66, 0xaf, 0x2e, 0xbd, 0xee, 0x1e, 0x9b, 0x69, 0x8f, 0x7e, 0x04, 0x3f, 0x98, 0x5f,
	0xc1, 0xcf, 0x23, 0xbb, 0xdb, 0xa4, 0x8d, 0x8a, 0xe0, 0xbb, 0x99, 0x79, 0xe6, 0x79, 0x26, 0x3b,
	0xf3, 0x04, 0xe2, 0x5b, 0xa5, 0x17, 0x57, 0xd7, 0xea, 0x76, 0x78, 0xa3, 0x15, 0x29, 0x0c, 0x8b,
	0xbc, 0xf7, 0xcd, 0x87, 0xf0, 0xf3, 0x36, 0x41, 0x84, 0xfa, 0x6a, 0x25, 0xb2, 0xc4, 0xeb, 0x7a,
	0xfd, 0x66, 0x6a, 0x63, 0x7c, 0x08, 0xad, 0x2b, 0x36, 0x23, 0xa5, 0x37, 0x53, 0xc9, 0x96, 0x3c,
	0xf1, 0x2d, 0x16, 0x6d, 0x6b, 0xe7, 0x6c, 0xc9, 0x0d, 0xcd, 0x42, 0x35, 0x47, 0x33, 0x31, 0x3e,
	0x85, 0x20, 0x27, 0x46, 0x3c, 0xa9, 0x77, 0xbd, 0x7e, 0x3c, 0xbe, 0x33, 0x2c, 0xbf, 0xa0, 0x98,
	0x36, 0x31, 0x70, 0xea, 0xba, 0x8c, 0x44, 0xc6, 0x88, 0x25, 0x41, 0xd7, 0xeb, 0xb7, 0x52, 0x1b,
	0xe3, 0x31, 0x04, 0x5c, 0x6b, 0xa5, 0x93, 0x03, 0xab, 0xeb, 0x12, 0xbc, 0x07, 0x90, 0x13, 0xd3,
	0x34, 0x25, 0xb1, 0xe4, 0x49, 0xa3, 0xeb, 0xf5, 0x6b, 0x69, 0xd3, 0x56, 0x2e, 0xc4, 0x92, 0xe3,
	0x5d, 0x08, 0xb9, 0xcc, 0x1c, 0x18, 0x5a, 0xb0, 0xc1, 0x65, 0x66, 0xa1, 0x07, 0x10, 0xcd, 0x34,
	0x67, 0xc4, 0x1d, 0xda, 0xb4, 0x28, 0xb8, 0x92, 0x69, 0xe8, 0x7d, 0xf7, 0x01, 0x8b, 0xaf, 0x7b,
	0xfb, 0x95, 0xcf, 0x16, 0x37, 0x4a, 0x48, 0x32, 0x1b, 0x98, 0xa9, 0x8c, 0x4f, 0xd7, 0x5c, 0xe7,
	0x42, 0x49, 0xbb, 0x9d, 0x20, 0x8d, 0x4c, 0xed, 0x93, 0x2b, 0xe1, 0x29, 0x04, 0xc4, 0xf2, 0x45,
	0x9e, 0xf8, 0xdd, 0x5a, 0x3f, 0x1a, 0x9f, 0xfc, 0xfe, 0xda, 0x9d, 0xde, 0xf0, 0xc2, 0x74, 0xbe,
	0x97, 0xa4, 0x37, 0xa9, 0x63, 0xe1, 0x07, 0x08, 0x73, 0x4e, 0x24, 0xe4, 0x3c, 0x4f, 0x6a, 0x56,
	0x61, 0xf0, 0x57, 0x85, 0xc9, 0xb6, 0xd9, 0x89, 0x94, 0xdc, 0xce, 0x47, 0x80, 0x9d, 0x38, 0xb6,
	0xa1, 0xb6, 0xe0, 0x9b, 0xed, 0x31, 0x4d, 0x88, 0x8f, 0x20, 0x58, 0xb3, 0xeb, 0x95, 0x3b, 0x62,
	0x34, 0x8e, 0x77, 0x43, 0x0c, 0x2d, 0x75, 0xe0, 0x4b, 0xff, 0x85, 0xd7, 0x79, 0x05, 0x87, 0x95,
	0x21, 0x7f, 0x10, 0x3b, 0xde, 0x17, 0x6b, 0xee, 0x91, 0x7b, 0x3f, 0x3c, 0xa8, 0x1b, 0x41, 0x8c,
	0xc1, 0x2f, 0xdd, 0xe4, 0x8b, 0x0c, 0x9f, 0x14, 0xa6, 0xf0, 0xad, 0x29, 0xfe, 0xaf, 0xce, 0xaf,
	0x18, 0xe2, 0x0c, 0x80, 0x11, 0x69, 0x71, 0xb9, 0x22, 0x5e, 0x2c, 0xe5, 0x7e, 0xb5, 0x7f, 0xf8,
	0xba, 0x6c, 0x70, 0x8b, 0xd8, 0x63, 0xec, 0xcc, 0x53, 0xdf, 0x33, 0x4f, 0xe7, 0x14, 0x8e, 0x7e,
	0x21, 0xfd, 0xcb, 0xc3, 0x06, 0xcf, 0xe1, 0xb0, 0xe2, 0x5e, 0x8c, 0x01, 0xce, 0x15, 0x4d, 0x8c,
	0xfb, 0x78, 0xd6, 0xfe, 0x0f, 0x23, 0x68, 0xa4, 0x2b, 0x29, 0x85, 0x9c, 0xb7, 0x3d, 0x0c, 0xa1,
	0xfe, 0x4e, 0x49, 0xde, 0xf6, 0x07, 0x67, 0xd0, 0x2c, 0x1f, 0x88, 0x08, 0xb1, 0x49, 0x2a, 0xbc,
	0x23, 0x88, 0xec, 0x05, 0x4a, 0x6e, 0x0b, 0x42, 0x53, 0x70, 0xfc, 0x37, 0x27, 0x5f, 0x1e, 0xaf,
	0x05, 0xf1, 0x3c, 0x1f, 0x0a, 0x35, 0x72, 0xd1, 0x68, 0xae, 0x46, 0x6b, 0x1a, 0xd9, 0xdf, 0x79,
	0x54, 0xac, 0xe5, 0xf2, 0xc0, 0xe6, 0xcf, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x75, 0x1d, 0xcd,
	0x85, 0xf0, 0x03, 0x00, 0x00,
}
