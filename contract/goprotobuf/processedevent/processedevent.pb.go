// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: contract/protobuf/processedevent/processedevent.proto

package processedevent

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConnectionType int32

const (
	ConnectionType_EVENT_STREAM ConnectionType = 0
	ConnectionType_STORAGE      ConnectionType = 1
	ConnectionType_REVERSE_ETL  ConnectionType = 2
)

// Enum value maps for ConnectionType.
var (
	ConnectionType_name = map[int32]string{
		0: "EVENT_STREAM",
		1: "STORAGE",
		2: "REVERSE_ETL",
	}
	ConnectionType_value = map[string]int32{
		"EVENT_STREAM": 0,
		"STORAGE":      1,
		"REVERSE_ETL":  2,
	}
)

func (x ConnectionType) Enum() *ConnectionType {
	p := new(ConnectionType)
	*p = x
	return p
}

func (x ConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_protobuf_processedevent_processedevent_proto_enumTypes[0].Descriptor()
}

func (ConnectionType) Type() protoreflect.EnumType {
	return &file_contract_protobuf_processedevent_processedevent_proto_enumTypes[0]
}

func (x ConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionType.Descriptor instead.
func (ConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP(), []int{0}
}

type Type int32

const (
	Type_TRACK    Type = 0
	Type_PAGE     Type = 1
	Type_IDENTITY Type = 2
	Type_GROUP    Type = 3
	Type_ALIAS    Type = 4
	Type_SCREEN   Type = 5
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TRACK",
		1: "PAGE",
		2: "IDENTITY",
		3: "GROUP",
		4: "ALIAS",
		5: "SCREEN",
	}
	Type_value = map[string]int32{
		"TRACK":    0,
		"PAGE":     1,
		"IDENTITY": 2,
		"GROUP":    3,
		"ALIAS":    4,
		"SCREEN":   5,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_protobuf_processedevent_processedevent_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_contract_protobuf_processedevent_processedevent_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP(), []int{1}
}

type DestinationType int32

const (
	DestinationType_webhook DestinationType = 0
)

// Enum value maps for DestinationType.
var (
	DestinationType_name = map[int32]string{
		0: "webhook",
	}
	DestinationType_value = map[string]int32{
		"webhook": 0,
	}
)

func (x DestinationType) Enum() *DestinationType {
	p := new(DestinationType)
	*p = x
	return p
}

func (x DestinationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DestinationType) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_protobuf_processedevent_processedevent_proto_enumTypes[2].Descriptor()
}

func (DestinationType) Type() protoreflect.EnumType {
	return &file_contract_protobuf_processedevent_processedevent_proto_enumTypes[2]
}

func (x DestinationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DestinationType.Descriptor instead.
func (DestinationType) EnumDescriptor() ([]byte, []int) {
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP(), []int{2}
}

type DestinationCategory int32

const (
	DestinationCategory_ANALYTICS       DestinationCategory = 0
	DestinationCategory_ADVERTISING     DestinationCategory = 1
	DestinationCategory_CRM             DestinationCategory = 2
	DestinationCategory_EMAIL_MARKETING DestinationCategory = 3
	DestinationCategory_LIVECHAT        DestinationCategory = 4
	DestinationCategory_PAYMENTS        DestinationCategory = 5
	DestinationCategory_SURVEYS         DestinationCategory = 6
)

// Enum value maps for DestinationCategory.
var (
	DestinationCategory_name = map[int32]string{
		0: "ANALYTICS",
		1: "ADVERTISING",
		2: "CRM",
		3: "EMAIL_MARKETING",
		4: "LIVECHAT",
		5: "PAYMENTS",
		6: "SURVEYS",
	}
	DestinationCategory_value = map[string]int32{
		"ANALYTICS":       0,
		"ADVERTISING":     1,
		"CRM":             2,
		"EMAIL_MARKETING": 3,
		"LIVECHAT":        4,
		"PAYMENTS":        5,
		"SURVEYS":         6,
	}
)

func (x DestinationCategory) Enum() *DestinationCategory {
	p := new(DestinationCategory)
	*p = x
	return p
}

func (x DestinationCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DestinationCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_protobuf_processedevent_processedevent_proto_enumTypes[3].Descriptor()
}

func (DestinationCategory) Type() protoreflect.EnumType {
	return &file_contract_protobuf_processedevent_processedevent_proto_enumTypes[3]
}

func (x DestinationCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DestinationCategory.Descriptor instead.
func (DestinationCategory) EnumDescriptor() ([]byte, []int) {
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP(), []int{3}
}

type WebhookMethod int32

const (
	WebhookMethod_POST  WebhookMethod = 0
	WebhookMethod_PUT   WebhookMethod = 1
	WebhookMethod_PATCH WebhookMethod = 2
)

// Enum value maps for WebhookMethod.
var (
	WebhookMethod_name = map[int32]string{
		0: "POST",
		1: "PUT",
		2: "PATCH",
	}
	WebhookMethod_value = map[string]int32{
		"POST":  0,
		"PUT":   1,
		"PATCH": 2,
	}
)

func (x WebhookMethod) Enum() *WebhookMethod {
	p := new(WebhookMethod)
	*p = x
	return p
}

func (x WebhookMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WebhookMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_protobuf_processedevent_processedevent_proto_enumTypes[4].Descriptor()
}

func (WebhookMethod) Type() protoreflect.EnumType {
	return &file_contract_protobuf_processedevent_processedevent_proto_enumTypes[4]
}

func (x WebhookMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WebhookMethod.Descriptor instead.
func (WebhookMethod) EnumDescriptor() ([]byte, []int) {
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP(), []int{4}
}

type ProcessedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId          string                 `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	TracerCarrier     map[string]string      `protobuf:"bytes,2,rep,name=tracer_carrier,json=tracerCarrier,proto3" json:"tracer_carrier,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Integration       *Integration           `protobuf:"bytes,3,opt,name=integration,proto3" json:"integration,omitempty"`
	MessageId         string                 `protobuf:"bytes,4,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	EventType         Type                   `protobuf:"varint,5,opt,name=event_type,json=eventType,proto3,enum=event.Type" json:"event_type,omitempty"`
	Version           uint32                 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	SentAt            *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	ReceivedAt        *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
	OriginalTimestamp *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=original_timestamp,json=originalTimestamp,proto3" json:"original_timestamp,omitempty"`
	Timestamp         *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ProcessedEvent) Reset() {
	*x = ProcessedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_processedevent_processedevent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessedEvent) ProtoMessage() {}

func (x *ProcessedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_processedevent_processedevent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessedEvent.ProtoReflect.Descriptor instead.
func (*ProcessedEvent) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessedEvent) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *ProcessedEvent) GetTracerCarrier() map[string]string {
	if x != nil {
		return x.TracerCarrier
	}
	return nil
}

func (x *ProcessedEvent) GetIntegration() *Integration {
	if x != nil {
		return x.Integration
	}
	return nil
}

func (x *ProcessedEvent) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *ProcessedEvent) GetEventType() Type {
	if x != nil {
		return x.EventType
	}
	return Type_TRACK
}

func (x *ProcessedEvent) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ProcessedEvent) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

func (x *ProcessedEvent) GetReceivedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedAt
	}
	return nil
}

func (x *ProcessedEvent) GetOriginalTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.OriginalTimestamp
	}
	return nil
}

func (x *ProcessedEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Integration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SourceId       string               `protobuf:"bytes,2,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	Name           string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Metadata       *DestinationMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ConnectionType ConnectionType       `protobuf:"varint,5,opt,name=connection_type,json=connectionType,proto3,enum=event.ConnectionType" json:"connection_type,omitempty"`
	Enabled        bool                 `protobuf:"varint,6,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Types that are assignable to Config:
	//
	//	*Integration_Fake
	//	*Integration_Webhook
	Config    isIntegration_Config   `protobuf_oneof:"config"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Integration) Reset() {
	*x = Integration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_processedevent_processedevent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integration) ProtoMessage() {}

func (x *Integration) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_processedevent_processedevent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integration.ProtoReflect.Descriptor instead.
func (*Integration) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP(), []int{1}
}

func (x *Integration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Integration) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *Integration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Integration) GetMetadata() *DestinationMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Integration) GetConnectionType() ConnectionType {
	if x != nil {
		return x.ConnectionType
	}
	return ConnectionType_EVENT_STREAM
}

func (x *Integration) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (m *Integration) GetConfig() isIntegration_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *Integration) GetFake() *FakeConfig {
	if x, ok := x.GetConfig().(*Integration_Fake); ok {
		return x.Fake
	}
	return nil
}

func (x *Integration) GetWebhook() *WebhookConfig {
	if x, ok := x.GetConfig().(*Integration_Webhook); ok {
		return x.Webhook
	}
	return nil
}

func (x *Integration) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type isIntegration_Config interface {
	isIntegration_Config()
}

type Integration_Fake struct {
	Fake *FakeConfig `protobuf:"bytes,100,opt,name=fake,proto3,oneof"`
}

type Integration_Webhook struct {
	Webhook *WebhookConfig `protobuf:"bytes,101,opt,name=webhook,proto3,oneof"`
}

func (*Integration_Fake) isIntegration_Config() {}

func (*Integration_Webhook) isIntegration_Config() {}

type DestinationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slug       DestinationType       `protobuf:"varint,3,opt,name=slug,proto3,enum=event.DestinationType" json:"slug,omitempty"`
	Categories []DestinationCategory `protobuf:"varint,4,rep,packed,name=categories,proto3,enum=event.DestinationCategory" json:"categories,omitempty"`
}

func (x *DestinationMetadata) Reset() {
	*x = DestinationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_processedevent_processedevent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationMetadata) ProtoMessage() {}

func (x *DestinationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_processedevent_processedevent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationMetadata.ProtoReflect.Descriptor instead.
func (*DestinationMetadata) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP(), []int{2}
}

func (x *DestinationMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DestinationMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DestinationMetadata) GetSlug() DestinationType {
	if x != nil {
		return x.Slug
	}
	return DestinationType_webhook
}

func (x *DestinationMetadata) GetCategories() []DestinationCategory {
	if x != nil {
		return x.Categories
	}
	return nil
}

// integration's protobuf
type FakeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FakeConfig) Reset() {
	*x = FakeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_processedevent_processedevent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FakeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FakeConfig) ProtoMessage() {}

func (x *FakeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_processedevent_processedevent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FakeConfig.ProtoReflect.Descriptor instead.
func (*FakeConfig) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP(), []int{3}
}

func (x *FakeConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WebhookConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers map[string]string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload map[string]string `protobuf:"bytes,2,rep,name=payload,proto3" json:"payload,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Method  WebhookMethod     `protobuf:"varint,3,opt,name=method,proto3,enum=event.WebhookMethod" json:"method,omitempty"`
	Url     string            `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *WebhookConfig) Reset() {
	*x = WebhookConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_processedevent_processedevent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookConfig) ProtoMessage() {}

func (x *WebhookConfig) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_processedevent_processedevent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookConfig.ProtoReflect.Descriptor instead.
func (*WebhookConfig) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP(), []int{4}
}

func (x *WebhookConfig) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *WebhookConfig) GetPayload() map[string]string {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *WebhookConfig) GetMethod() WebhookMethod {
	if x != nil {
		return x.Method
	}
	return WebhookMethod_POST
}

func (x *WebhookConfig) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_contract_protobuf_processedevent_processedevent_proto protoreflect.FileDescriptor

var file_contract_protobuf_processedevent_processedevent_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd2, 0x04, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x4f, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72,
	0x12, 0x34, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x07, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74,
	0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x49, 0x0a,
	0x12, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x1a, 0x40, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x43, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x03, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a,
	0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x66, 0x61, 0x6b, 0x65, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x61,
	0x6b, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x04, 0x66, 0x61, 0x6b, 0x65,
	0x12, 0x30, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x08, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12,
	0x3a, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x20, 0x0a, 0x0a, 0x46,
	0x61, 0x6b, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc1, 0x02,
	0x0a, 0x0d, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x3b, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x2a, 0x40, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x52,
	0x45, 0x41, 0x4d, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x56, 0x45, 0x52, 0x53, 0x45, 0x5f, 0x45, 0x54,
	0x4c, 0x10, 0x02, 0x2a, 0x4b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x54,
	0x52, 0x41, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x47, 0x45, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x02, 0x12, 0x09,
	0x0a, 0x05, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x49,
	0x41, 0x53, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x05,
	0x2a, 0x1e, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x10, 0x00,
	0x2a, 0x7c, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x4e, 0x41, 0x4c, 0x59,
	0x54, 0x49, 0x43, 0x53, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x44, 0x56, 0x45, 0x52, 0x54,
	0x49, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x52, 0x4d, 0x10, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x49, 0x56, 0x45, 0x43, 0x48, 0x41,
	0x54, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10,
	0x05, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x52, 0x56, 0x45, 0x59, 0x53, 0x10, 0x06, 0x2a, 0x2d,
	0x0a, 0x0d, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x02, 0x42, 0x24, 0x5a,
	0x22, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_protobuf_processedevent_processedevent_proto_rawDescOnce sync.Once
	file_contract_protobuf_processedevent_processedevent_proto_rawDescData = file_contract_protobuf_processedevent_processedevent_proto_rawDesc
)

func file_contract_protobuf_processedevent_processedevent_proto_rawDescGZIP() []byte {
	file_contract_protobuf_processedevent_processedevent_proto_rawDescOnce.Do(func() {
		file_contract_protobuf_processedevent_processedevent_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_protobuf_processedevent_processedevent_proto_rawDescData)
	})
	return file_contract_protobuf_processedevent_processedevent_proto_rawDescData
}

var file_contract_protobuf_processedevent_processedevent_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_contract_protobuf_processedevent_processedevent_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_contract_protobuf_processedevent_processedevent_proto_goTypes = []interface{}{
	(ConnectionType)(0),           // 0: event.ConnectionType
	(Type)(0),                     // 1: event.Type
	(DestinationType)(0),          // 2: event.DestinationType
	(DestinationCategory)(0),      // 3: event.DestinationCategory
	(WebhookMethod)(0),            // 4: event.WebhookMethod
	(*ProcessedEvent)(nil),        // 5: event.ProcessedEvent
	(*Integration)(nil),           // 6: event.Integration
	(*DestinationMetadata)(nil),   // 7: event.DestinationMetadata
	(*FakeConfig)(nil),            // 8: event.FakeConfig
	(*WebhookConfig)(nil),         // 9: event.WebhookConfig
	nil,                           // 10: event.ProcessedEvent.TracerCarrierEntry
	nil,                           // 11: event.WebhookConfig.HeadersEntry
	nil,                           // 12: event.WebhookConfig.PayloadEntry
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_contract_protobuf_processedevent_processedevent_proto_depIdxs = []int32{
	10, // 0: event.ProcessedEvent.tracer_carrier:type_name -> event.ProcessedEvent.TracerCarrierEntry
	6,  // 1: event.ProcessedEvent.integration:type_name -> event.Integration
	1,  // 2: event.ProcessedEvent.event_type:type_name -> event.Type
	13, // 3: event.ProcessedEvent.sent_at:type_name -> google.protobuf.Timestamp
	13, // 4: event.ProcessedEvent.received_at:type_name -> google.protobuf.Timestamp
	13, // 5: event.ProcessedEvent.original_timestamp:type_name -> google.protobuf.Timestamp
	13, // 6: event.ProcessedEvent.timestamp:type_name -> google.protobuf.Timestamp
	7,  // 7: event.Integration.metadata:type_name -> event.DestinationMetadata
	0,  // 8: event.Integration.connection_type:type_name -> event.ConnectionType
	8,  // 9: event.Integration.fake:type_name -> event.FakeConfig
	9,  // 10: event.Integration.webhook:type_name -> event.WebhookConfig
	13, // 11: event.Integration.created_at:type_name -> google.protobuf.Timestamp
	2,  // 12: event.DestinationMetadata.slug:type_name -> event.DestinationType
	3,  // 13: event.DestinationMetadata.categories:type_name -> event.DestinationCategory
	11, // 14: event.WebhookConfig.headers:type_name -> event.WebhookConfig.HeadersEntry
	12, // 15: event.WebhookConfig.payload:type_name -> event.WebhookConfig.PayloadEntry
	4,  // 16: event.WebhookConfig.method:type_name -> event.WebhookMethod
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_contract_protobuf_processedevent_processedevent_proto_init() }
func file_contract_protobuf_processedevent_processedevent_proto_init() {
	if File_contract_protobuf_processedevent_processedevent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_protobuf_processedevent_processedevent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessedEvent); i {
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
		file_contract_protobuf_processedevent_processedevent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integration); i {
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
		file_contract_protobuf_processedevent_processedevent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationMetadata); i {
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
		file_contract_protobuf_processedevent_processedevent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FakeConfig); i {
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
		file_contract_protobuf_processedevent_processedevent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookConfig); i {
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
	file_contract_protobuf_processedevent_processedevent_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Integration_Fake)(nil),
		(*Integration_Webhook)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contract_protobuf_processedevent_processedevent_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_protobuf_processedevent_processedevent_proto_goTypes,
		DependencyIndexes: file_contract_protobuf_processedevent_processedevent_proto_depIdxs,
		EnumInfos:         file_contract_protobuf_processedevent_processedevent_proto_enumTypes,
		MessageInfos:      file_contract_protobuf_processedevent_processedevent_proto_msgTypes,
	}.Build()
	File_contract_protobuf_processedevent_processedevent_proto = out.File
	file_contract_protobuf_processedevent_processedevent_proto_rawDesc = nil
	file_contract_protobuf_processedevent_processedevent_proto_goTypes = nil
	file_contract_protobuf_processedevent_processedevent_proto_depIdxs = nil
}
