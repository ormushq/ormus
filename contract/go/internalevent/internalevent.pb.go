// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: internalevent/internalevent.proto

package internalevent

import (
	reflect "reflect"
	sync "sync"

	project "github.com/ormushq/ormus/contract/go/project"
	task "github.com/ormushq/ormus/contract/go/task"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventName int32

const (
	EventName_EVENT_NAME_USER_CREATED        EventName = 0
	EventName_EVENT_NAME_PROJECT_CREATED     EventName = 1
	EventName_EVENT_NAME_WRITE_KEY_GENERATED EventName = 2
	EventName_EVENT_NAME_TASK_CREATED        EventName = 3
)

// Enum value maps for EventName.
var (
	EventName_name = map[int32]string{
		0: "EVENT_NAME_USER_CREATED",
		1: "EVENT_NAME_PROJECT_CREATED",
		2: "EVENT_NAME_WRITE_KEY_GENERATED",
		3: "EVENT_NAME_TASK_CREATED",
	}
	EventName_value = map[string]int32{
		"EVENT_NAME_USER_CREATED":        0,
		"EVENT_NAME_PROJECT_CREATED":     1,
		"EVENT_NAME_WRITE_KEY_GENERATED": 2,
		"EVENT_NAME_TASK_CREATED":        3,
	}
)

func (x EventName) Enum() *EventName {
	p := new(EventName)
	*p = x
	return p
}

func (x EventName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventName) Descriptor() protoreflect.EnumDescriptor {
	return file_internalevent_internalevent_proto_enumTypes[0].Descriptor()
}

func (EventName) Type() protoreflect.EnumType {
	return &file_internalevent_internalevent_proto_enumTypes[0]
}

func (x EventName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventName.Descriptor instead.
func (EventName) EnumDescriptor() ([]byte, []int) {
	return file_internalevent_internalevent_proto_rawDescGZIP(), []int{0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventName     EventName              `protobuf:"varint,2,opt,name=event_name,json=eventName,proto3,enum=internalevent.EventName" json:"event_name,omitempty"`
	Time          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	TracerCarrier map[string]string      `protobuf:"bytes,4,rep,name=tracer_carrier,json=tracerCarrier,proto3" json:"tracer_carrier,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to Payload:
	//
	//	*Event_UserCreatedEvent
	//	*Event_ProjectCreatedEvent
	//	*Event_WriteKeyGeneratedEvent
	//	*Event_TaskCreatedEvent
	Payload isEvent_Payload `protobuf_oneof:"payload"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalevent_internalevent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_internalevent_internalevent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_internalevent_internalevent_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetEventName() EventName {
	if x != nil {
		return x.EventName
	}
	return EventName_EVENT_NAME_USER_CREATED
}

func (x *Event) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Event) GetTracerCarrier() map[string]string {
	if x != nil {
		return x.TracerCarrier
	}
	return nil
}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Event) GetUserCreatedEvent() *UserCreatedEvent {
	if x, ok := x.GetPayload().(*Event_UserCreatedEvent); ok {
		return x.UserCreatedEvent
	}
	return nil
}

func (x *Event) GetProjectCreatedEvent() *ProjectCreatedEvent {
	if x, ok := x.GetPayload().(*Event_ProjectCreatedEvent); ok {
		return x.ProjectCreatedEvent
	}
	return nil
}

func (x *Event) GetWriteKeyGeneratedEvent() *WriteKeyGeneratedEvent {
	if x, ok := x.GetPayload().(*Event_WriteKeyGeneratedEvent); ok {
		return x.WriteKeyGeneratedEvent
	}
	return nil
}

func (x *Event) GetTaskCreatedEvent() *TaskCreatedEvent {
	if x, ok := x.GetPayload().(*Event_TaskCreatedEvent); ok {
		return x.TaskCreatedEvent
	}
	return nil
}

type isEvent_Payload interface {
	isEvent_Payload()
}

type Event_UserCreatedEvent struct {
	UserCreatedEvent *UserCreatedEvent `protobuf:"bytes,100,opt,name=user_created_event,json=userCreatedEvent,proto3,oneof"`
}

type Event_ProjectCreatedEvent struct {
	ProjectCreatedEvent *ProjectCreatedEvent `protobuf:"bytes,101,opt,name=project_created_event,json=projectCreatedEvent,proto3,oneof"`
}

type Event_WriteKeyGeneratedEvent struct {
	WriteKeyGeneratedEvent *WriteKeyGeneratedEvent `protobuf:"bytes,102,opt,name=write_key_generated_event,json=writeKeyGeneratedEvent,proto3,oneof"`
}

type Event_TaskCreatedEvent struct {
	TaskCreatedEvent *TaskCreatedEvent `protobuf:"bytes,103,opt,name=task_created_event,json=taskCreatedEvent,proto3,oneof"`
}

func (*Event_UserCreatedEvent) isEvent_Payload() {}

func (*Event_ProjectCreatedEvent) isEvent_Payload() {}

func (*Event_WriteKeyGeneratedEvent) isEvent_Payload() {}

func (*Event_TaskCreatedEvent) isEvent_Payload() {}

type UserCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *UserCreatedEvent) Reset() {
	*x = UserCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalevent_internalevent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreatedEvent) ProtoMessage() {}

func (x *UserCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_internalevent_internalevent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreatedEvent.ProtoReflect.Descriptor instead.
func (*UserCreatedEvent) Descriptor() ([]byte, []int) {
	return file_internalevent_internalevent_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreatedEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserCreatedEvent) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ProjectCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *project.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ProjectCreatedEvent) Reset() {
	*x = ProjectCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalevent_internalevent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectCreatedEvent) ProtoMessage() {}

func (x *ProjectCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_internalevent_internalevent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectCreatedEvent.ProtoReflect.Descriptor instead.
func (*ProjectCreatedEvent) Descriptor() ([]byte, []int) {
	return file_internalevent_internalevent_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectCreatedEvent) GetProject() *project.Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type WriteKeyGeneratedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *project.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *WriteKeyGeneratedEvent) Reset() {
	*x = WriteKeyGeneratedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalevent_internalevent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteKeyGeneratedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteKeyGeneratedEvent) ProtoMessage() {}

func (x *WriteKeyGeneratedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_internalevent_internalevent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteKeyGeneratedEvent.ProtoReflect.Descriptor instead.
func (*WriteKeyGeneratedEvent) Descriptor() ([]byte, []int) {
	return file_internalevent_internalevent_proto_rawDescGZIP(), []int{3}
}

func (x *WriteKeyGeneratedEvent) GetProject() *project.Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type TaskCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *task.Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *TaskCreatedEvent) Reset() {
	*x = TaskCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalevent_internalevent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreatedEvent) ProtoMessage() {}

func (x *TaskCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_internalevent_internalevent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreatedEvent.ProtoReflect.Descriptor instead.
func (*TaskCreatedEvent) Descriptor() ([]byte, []int) {
	return file_internalevent_internalevent_proto_rawDescGZIP(), []int{4}
}

func (x *TaskCreatedEvent) GetTask() *task.Task {
	if x != nil {
		return x.Task
	}
	return nil
}

var File_internalevent_internalevent_proto protoreflect.FileDescriptor

var file_internalevent_internalevent_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x61, 0x73, 0x6b,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x04, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x4e,
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x72, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x4f,
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10, 0x75,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x58, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x62, 0x0a, 0x19, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x16, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4f, 0x0a,
	0x12, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10, 0x74, 0x61,
	0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x40,
	0x0a, 0x12, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x66, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x41, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x44, 0x0a, 0x16, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x2a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x32, 0x0a, 0x10,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1e, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b,
	0x2a, 0x89, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43,
	0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f,
	0x4b, 0x45, 0x59, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x1b, 0x0a, 0x17, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x34, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x6d, 0x75, 0x73,
	0x68, 0x71, 0x2f, 0x6f, 0x72, 0x6d, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalevent_internalevent_proto_rawDescOnce sync.Once
	file_internalevent_internalevent_proto_rawDescData = file_internalevent_internalevent_proto_rawDesc
)

func file_internalevent_internalevent_proto_rawDescGZIP() []byte {
	file_internalevent_internalevent_proto_rawDescOnce.Do(func() {
		file_internalevent_internalevent_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalevent_internalevent_proto_rawDescData)
	})
	return file_internalevent_internalevent_proto_rawDescData
}

var file_internalevent_internalevent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internalevent_internalevent_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internalevent_internalevent_proto_goTypes = []any{
	(EventName)(0),                 // 0: internalevent.EventName
	(*Event)(nil),                  // 1: internalevent.Event
	(*UserCreatedEvent)(nil),       // 2: internalevent.UserCreatedEvent
	(*ProjectCreatedEvent)(nil),    // 3: internalevent.ProjectCreatedEvent
	(*WriteKeyGeneratedEvent)(nil), // 4: internalevent.WriteKeyGeneratedEvent
	(*TaskCreatedEvent)(nil),       // 5: internalevent.TaskCreatedEvent
	nil,                            // 6: internalevent.Event.TracerCarrierEntry
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*project.Project)(nil),        // 8: project.Project
	(*task.Task)(nil),              // 9: task.Task
}
var file_internalevent_internalevent_proto_depIdxs = []int32{
	0,  // 0: internalevent.Event.event_name:type_name -> internalevent.EventName
	7,  // 1: internalevent.Event.time:type_name -> google.protobuf.Timestamp
	6,  // 2: internalevent.Event.tracer_carrier:type_name -> internalevent.Event.TracerCarrierEntry
	2,  // 3: internalevent.Event.user_created_event:type_name -> internalevent.UserCreatedEvent
	3,  // 4: internalevent.Event.project_created_event:type_name -> internalevent.ProjectCreatedEvent
	4,  // 5: internalevent.Event.write_key_generated_event:type_name -> internalevent.WriteKeyGeneratedEvent
	5,  // 6: internalevent.Event.task_created_event:type_name -> internalevent.TaskCreatedEvent
	7,  // 7: internalevent.UserCreatedEvent.created_at:type_name -> google.protobuf.Timestamp
	8,  // 8: internalevent.ProjectCreatedEvent.project:type_name -> project.Project
	8,  // 9: internalevent.WriteKeyGeneratedEvent.project:type_name -> project.Project
	9,  // 10: internalevent.TaskCreatedEvent.task:type_name -> task.Task
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_internalevent_internalevent_proto_init() }
func file_internalevent_internalevent_proto_init() {
	if File_internalevent_internalevent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internalevent_internalevent_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internalevent_internalevent_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserCreatedEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internalevent_internalevent_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ProjectCreatedEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internalevent_internalevent_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WriteKeyGeneratedEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internalevent_internalevent_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TaskCreatedEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_internalevent_internalevent_proto_msgTypes[0].OneofWrappers = []any{
		(*Event_UserCreatedEvent)(nil),
		(*Event_ProjectCreatedEvent)(nil),
		(*Event_WriteKeyGeneratedEvent)(nil),
		(*Event_TaskCreatedEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalevent_internalevent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalevent_internalevent_proto_goTypes,
		DependencyIndexes: file_internalevent_internalevent_proto_depIdxs,
		EnumInfos:         file_internalevent_internalevent_proto_enumTypes,
		MessageInfos:      file_internalevent_internalevent_proto_msgTypes,
	}.Build()
	File_internalevent_internalevent_proto = out.File
	file_internalevent_internalevent_proto_rawDesc = nil
	file_internalevent_internalevent_proto_goTypes = nil
	file_internalevent_internalevent_proto_depIdxs = nil
}
