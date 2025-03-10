// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: source/source.proto

package source

import (
	reflect "reflect"
	sync "sync"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_STATUS_ACTIVE     Status = 0
	Status_STATUS_NOT_ACTIVE Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_ACTIVE",
		1: "STATUS_NOT_ACTIVE",
	}
	Status_value = map[string]int32{
		"STATUS_ACTIVE":     0,
		"STATUS_NOT_ACTIVE": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_source_source_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_source_source_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_source_source_proto_rawDescGZIP(), []int{0}
}

type SourceMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slug     string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Category string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *SourceMetadata) Reset() {
	*x = SourceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceMetadata) ProtoMessage() {}

func (x *SourceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_source_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceMetadata.ProtoReflect.Descriptor instead.
func (*SourceMetadata) Descriptor() ([]byte, []int) {
	return file_source_source_proto_rawDescGZIP(), []int{0}
}

func (x *SourceMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SourceMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SourceMetadata) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *SourceMetadata) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WriteKey    string               `protobuf:"bytes,2,opt,name=write_key,json=writeKey,proto3" json:"write_key,omitempty"`
	Name        string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string               `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ProjectId   string               `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	OwnerId     string               `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Status      Status               `protobuf:"varint,7,opt,name=status,proto3,enum=source.Status" json:"status,omitempty"`
	Metadata    *SourceMetadata      `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   *timestamp.Timestamp `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_source_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_source_source_proto_rawDescGZIP(), []int{1}
}

func (x *Source) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Source) GetWriteKey() string {
	if x != nil {
		return x.WriteKey
	}
	return ""
}

func (x *Source) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Source) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Source) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Source) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Source) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_ACTIVE
}

func (x *Source) GetMetadata() *SourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Source) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Source) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Source) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type NewSourceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	OwnerId   string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	WriteKey  string `protobuf:"bytes,3,opt,name=write_key,json=writeKey,proto3" json:"write_key,omitempty"`
}

func (x *NewSourceEvent) Reset() {
	*x = NewSourceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSourceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSourceEvent) ProtoMessage() {}

func (x *NewSourceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_source_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSourceEvent.ProtoReflect.Descriptor instead.
func (*NewSourceEvent) Descriptor() ([]byte, []int) {
	return file_source_source_proto_rawDescGZIP(), []int{2}
}

func (x *NewSourceEvent) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *NewSourceEvent) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *NewSourceEvent) GetWriteKey() string {
	if x != nil {
		return x.WriteKey
	}
	return ""
}

type ValidateWriteKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WriteKey string `protobuf:"bytes,1,opt,name=write_key,json=writeKey,proto3" json:"write_key,omitempty"`
}

func (x *ValidateWriteKeyReq) Reset() {
	*x = ValidateWriteKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_source_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateWriteKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateWriteKeyReq) ProtoMessage() {}

func (x *ValidateWriteKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_source_source_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateWriteKeyReq.ProtoReflect.Descriptor instead.
func (*ValidateWriteKeyReq) Descriptor() ([]byte, []int) {
	return file_source_source_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateWriteKeyReq) GetWriteKey() string {
	if x != nil {
		return x.WriteKey
	}
	return ""
}

type ValidateWriteKeyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid   bool   `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	OwnerId   string `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	WriteKey  string `protobuf:"bytes,4,opt,name=write_key,json=writeKey,proto3" json:"write_key,omitempty"`
}

func (x *ValidateWriteKeyResp) Reset() {
	*x = ValidateWriteKeyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_source_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateWriteKeyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateWriteKeyResp) ProtoMessage() {}

func (x *ValidateWriteKeyResp) ProtoReflect() protoreflect.Message {
	mi := &file_source_source_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateWriteKeyResp.ProtoReflect.Descriptor instead.
func (*ValidateWriteKeyResp) Descriptor() ([]byte, []int) {
	return file_source_source_proto_rawDescGZIP(), []int{4}
}

func (x *ValidateWriteKeyResp) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *ValidateWriteKeyResp) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ValidateWriteKeyResp) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *ValidateWriteKeyResp) GetWriteKey() string {
	if x != nil {
		return x.WriteKey
	}
	return ""
}

type NewEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type       string               `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name       string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Properties map[string]string    `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SendAt     *timestamp.Timestamp `protobuf:"bytes,5,opt,name=send_at,json=sendAt,proto3" json:"send_at,omitempty"`
	ReceivedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
	Timestamp  *timestamp.Timestamp `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Event      string               `protobuf:"bytes,8,opt,name=event,proto3" json:"event,omitempty"`
	WriteKey   string               `protobuf:"bytes,9,opt,name=write_key,json=writeKey,proto3" json:"write_key,omitempty"`
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *NewEvent) Reset() {
	*x = NewEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_source_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEvent) ProtoMessage() {}

func (x *NewEvent) ProtoReflect() protoreflect.Message {
	mi := &file_source_source_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEvent.ProtoReflect.Descriptor instead.
func (*NewEvent) Descriptor() ([]byte, []int) {
	return file_source_source_proto_rawDescGZIP(), []int{5}
}

func (x *NewEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NewEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NewEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewEvent) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *NewEvent) GetSendAt() *timestamp.Timestamp {
	if x != nil {
		return x.SendAt
	}
	return nil
}

func (x *NewEvent) GetReceivedAt() *timestamp.Timestamp {
	if x != nil {
		return x.ReceivedAt
	}
	return nil
}

func (x *NewEvent) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *NewEvent) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *NewEvent) GetWriteKey() string {
	if x != nil {
		return x.WriteKey
	}
	return ""
}

func (x *NewEvent) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *NewEvent) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_source_source_proto protoreflect.FileDescriptor

var file_source_source_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64,
	0x0a, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x22, 0xb2, 0x03, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x67, 0x0a, 0x0e, 0x4e, 0x65, 0x77,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x22, 0x32, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x88, 0x01, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x22, 0x98, 0x04, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a,
	0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x3d, 0x0a,
	0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x32, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01,
	0x32, 0x5f, 0x0a, 0x0f, 0x49, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x4c, 0x0a, 0x0f, 0x49, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x1b, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x72, 0x6d, 0x75, 0x73, 0x68, 0x71, 0x2f, 0x6f, 0x72, 0x6d, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_source_source_proto_rawDescOnce sync.Once
	file_source_source_proto_rawDescData = file_source_source_proto_rawDesc
)

func file_source_source_proto_rawDescGZIP() []byte {
	file_source_source_proto_rawDescOnce.Do(func() {
		file_source_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_source_source_proto_rawDescData)
	})
	return file_source_source_proto_rawDescData
}

var file_source_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_source_source_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_source_source_proto_goTypes = []interface{}{
	(Status)(0),                  // 0: source.Status
	(*SourceMetadata)(nil),       // 1: source.SourceMetadata
	(*Source)(nil),               // 2: source.Source
	(*NewSourceEvent)(nil),       // 3: source.NewSourceEvent
	(*ValidateWriteKeyReq)(nil),  // 4: source.ValidateWriteKeyReq
	(*ValidateWriteKeyResp)(nil), // 5: source.ValidateWriteKeyResp
	(*NewEvent)(nil),             // 6: source.NewEvent
	nil,                          // 7: source.NewEvent.PropertiesEntry
	(*timestamp.Timestamp)(nil),  // 8: google.protobuf.Timestamp
}
var file_source_source_proto_depIdxs = []int32{
	0,  // 0: source.Source.status:type_name -> source.Status
	1,  // 1: source.Source.metadata:type_name -> source.SourceMetadata
	8,  // 2: source.Source.created_at:type_name -> google.protobuf.Timestamp
	8,  // 3: source.Source.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 4: source.Source.deleted_at:type_name -> google.protobuf.Timestamp
	7,  // 5: source.NewEvent.properties:type_name -> source.NewEvent.PropertiesEntry
	8,  // 6: source.NewEvent.send_at:type_name -> google.protobuf.Timestamp
	8,  // 7: source.NewEvent.received_at:type_name -> google.protobuf.Timestamp
	8,  // 8: source.NewEvent.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 9: source.NewEvent.created_at:type_name -> google.protobuf.Timestamp
	8,  // 10: source.NewEvent.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 11: source.IsWriteKeyValid.IsWriteKeyValid:input_type -> source.ValidateWriteKeyReq
	5,  // 12: source.IsWriteKeyValid.IsWriteKeyValid:output_type -> source.ValidateWriteKeyResp
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_source_source_proto_init() }
func file_source_source_proto_init() {
	if File_source_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_source_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceMetadata); i {
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
		file_source_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
		file_source_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSourceEvent); i {
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
		file_source_source_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateWriteKeyReq); i {
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
		file_source_source_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateWriteKeyResp); i {
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
		file_source_source_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEvent); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_source_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_source_source_proto_goTypes,
		DependencyIndexes: file_source_source_proto_depIdxs,
		EnumInfos:         file_source_source_proto_enumTypes,
		MessageInfos:      file_source_source_proto_msgTypes,
	}.Build()
	File_source_source_proto = out.File
	file_source_source_proto_rawDesc = nil
	file_source_source_proto_goTypes = nil
	file_source_source_proto_depIdxs = nil
}
