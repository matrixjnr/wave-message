// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: pkg/service/service.proto

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId    string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	SenderId     string `protobuf:"bytes,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Payload      []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp    int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MessageId    string `protobuf:"bytes,5,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	IsPersistent bool   `protobuf:"varint,6,opt,name=is_persistent,json=isPersistent,proto3" json:"is_persistent,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_message_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_message_message_proto_msgTypes[0]
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
	return file_pkg_message_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Message) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Message) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetIsPersistent() bool {
	if x != nil {
		return x.IsPersistent
	}
	return false
}

var File_pkg_message_message_proto protoreflect.FileDescriptor

var file_pkg_message_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0xc1, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x50, 0x65,
	0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2e, 0x2f, 0x2e,
	0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_message_message_proto_rawDescOnce sync.Once
	file_pkg_message_message_proto_rawDescData = file_pkg_message_message_proto_rawDesc
)

func file_pkg_message_message_proto_rawDescGZIP() []byte {
	file_pkg_message_message_proto_rawDescOnce.Do(func() {
		file_pkg_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_message_message_proto_rawDescData)
	})
	return file_pkg_message_message_proto_rawDescData
}

var file_pkg_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_message_message_proto_goTypes = []any{
	(*Message)(nil), // 0: service.Message
}
var file_pkg_message_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_message_message_proto_init() }
func file_pkg_message_message_proto_init() {
	if File_pkg_message_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_message_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
			RawDescriptor: file_pkg_message_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_message_message_proto_goTypes,
		DependencyIndexes: file_pkg_message_message_proto_depIdxs,
		MessageInfos:      file_pkg_message_message_proto_msgTypes,
	}.Build()
	File_pkg_message_message_proto = out.File
	file_pkg_message_message_proto_rawDesc = nil
	file_pkg_message_message_proto_goTypes = nil
	file_pkg_message_message_proto_depIdxs = nil
}
