// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: contract/protobuf/messages/message.proto

package message

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BrokerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic   string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *BrokerMessage) Reset() {
	*x = BrokerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_messages_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerMessage) ProtoMessage() {}

func (x *BrokerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_messages_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerMessage.ProtoReflect.Descriptor instead.
func (*BrokerMessage) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_messages_message_proto_rawDescGZIP(), []int{0}
}

func (x *BrokerMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BrokerMessage) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *BrokerMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ChannelMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ChannelMessage) Reset() {
	*x = ChannelMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_protobuf_messages_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelMessage) ProtoMessage() {}

func (x *ChannelMessage) ProtoReflect() protoreflect.Message {
	mi := &file_contract_protobuf_messages_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelMessage.ProtoReflect.Descriptor instead.
func (*ChannelMessage) Descriptor() ([]byte, []int) {
	return file_contract_protobuf_messages_message_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelMessage) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_contract_protobuf_messages_message_proto protoreflect.FileDescriptor

var file_contract_protobuf_messages_message_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x1d, 0x5a, 0x1b, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_contract_protobuf_messages_message_proto_rawDescOnce sync.Once
	file_contract_protobuf_messages_message_proto_rawDescData = file_contract_protobuf_messages_message_proto_rawDesc
)

func file_contract_protobuf_messages_message_proto_rawDescGZIP() []byte {
	file_contract_protobuf_messages_message_proto_rawDescOnce.Do(func() {
		file_contract_protobuf_messages_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_protobuf_messages_message_proto_rawDescData)
	})
	return file_contract_protobuf_messages_message_proto_rawDescData
}

var file_contract_protobuf_messages_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contract_protobuf_messages_message_proto_goTypes = []any{
	(*BrokerMessage)(nil),  // 0: message.BrokerMessage
	(*ChannelMessage)(nil), // 1: message.ChannelMessage
}
var file_contract_protobuf_messages_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contract_protobuf_messages_message_proto_init() }
func file_contract_protobuf_messages_message_proto_init() {
	if File_contract_protobuf_messages_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_protobuf_messages_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BrokerMessage); i {
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
		file_contract_protobuf_messages_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ChannelMessage); i {
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
			RawDescriptor: file_contract_protobuf_messages_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_protobuf_messages_message_proto_goTypes,
		DependencyIndexes: file_contract_protobuf_messages_message_proto_depIdxs,
		MessageInfos:      file_contract_protobuf_messages_message_proto_msgTypes,
	}.Build()
	File_contract_protobuf_messages_message_proto = out.File
	file_contract_protobuf_messages_message_proto_rawDesc = nil
	file_contract_protobuf_messages_message_proto_goTypes = nil
	file_contract_protobuf_messages_message_proto_depIdxs = nil
}
