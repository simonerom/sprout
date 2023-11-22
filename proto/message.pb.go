// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: message.proto

package proto

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

type MessageState int32

const (
	MessageState_MESSAGE_STATE_INVALID    MessageState = 0
	MessageState_MESSAGE_STATE_RECEIVED   MessageState = 1
	MessageState_MESSAGE_STATE_FETCHED    MessageState = 2
	MessageState_MESSAGE_STATE_PROVING    MessageState = 3
	MessageState_MESSAGE_STATE_PROVED     MessageState = 4
	MessageState_MESSAGE_STATE_OUTPUTTING MessageState = 5
	MessageState_MESSAGE_STATE_OUTPUTTED  MessageState = 6
	MessageState_MESSAGE_STATE_FAILED     MessageState = 7
)

// Enum value maps for MessageState.
var (
	MessageState_name = map[int32]string{
		0: "MESSAGE_STATE_INVALID",
		1: "MESSAGE_STATE_RECEIVED",
		2: "MESSAGE_STATE_FETCHED",
		3: "MESSAGE_STATE_PROVING",
		4: "MESSAGE_STATE_PROVED",
		5: "MESSAGE_STATE_OUTPUTTING",
		6: "MESSAGE_STATE_OUTPUTTED",
		7: "MESSAGE_STATE_FAILED",
	}
	MessageState_value = map[string]int32{
		"MESSAGE_STATE_INVALID":    0,
		"MESSAGE_STATE_RECEIVED":   1,
		"MESSAGE_STATE_FETCHED":    2,
		"MESSAGE_STATE_PROVING":    3,
		"MESSAGE_STATE_PROVED":     4,
		"MESSAGE_STATE_OUTPUTTING": 5,
		"MESSAGE_STATE_OUTPUTTED":  6,
		"MESSAGE_STATE_FAILED":     7,
	}
)

func (x MessageState) Enum() *MessageState {
	p := new(MessageState)
	*p = x
	return p
}

func (x MessageState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageState) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (MessageState) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x MessageState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageState.Descriptor instead.
func (MessageState) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageID string `protobuf:"bytes,1,opt,name=messageID,proto3" json:"messageID,omitempty"`
	ProjectID uint64 `protobuf:"varint,2,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Data      string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessageID() string {
	if x != nil {
		return x.MessageID
	}
	return ""
}

func (x *Message) GetProjectID() uint64 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *Message) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x2a, 0xea, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a,
	0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x45, 0x54, 0x43, 0x48,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12,
	0x18, 0x0a, 0x14, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55,
	0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x53, 0x53, 0x41,
	0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x54,
	0x45, 0x44, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x07, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_proto_goTypes = []interface{}{
	(MessageState)(0), // 0: proto.MessageState
	(*Message)(nil),   // 1: proto.Message
}
var file_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
