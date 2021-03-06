// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: hexes/v1/event_payload.proto

package hexesv1

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

type PayloadChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text       string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	AlliesOnly bool   `protobuf:"varint,2,opt,name=allies_only,json=alliesOnly,proto3" json:"allies_only,omitempty"`
}

func (x *PayloadChat) Reset() {
	*x = PayloadChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hexes_v1_event_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadChat) ProtoMessage() {}

func (x *PayloadChat) ProtoReflect() protoreflect.Message {
	mi := &file_hexes_v1_event_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadChat.ProtoReflect.Descriptor instead.
func (*PayloadChat) Descriptor() ([]byte, []int) {
	return file_hexes_v1_event_payload_proto_rawDescGZIP(), []int{0}
}

func (x *PayloadChat) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PayloadChat) GetAlliesOnly() bool {
	if x != nil {
		return x.AlliesOnly
	}
	return false
}

var File_hexes_v1_event_payload_proto protoreflect.FileDescriptor

var file_hexes_v1_event_payload_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x68, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x42, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x43, 0x68, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61,
	0x6c, 0x6c, 0x69, 0x65, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x69, 0x65, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x42, 0x77, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x13, 0x67, 0x6f, 0x2f, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x68,
	0x65, 0x78, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x48,
	0x65, 0x78, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x48, 0x65, 0x78, 0x65, 0x73, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x14, 0x48, 0x65, 0x78, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x48, 0x65, 0x78, 0x65,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hexes_v1_event_payload_proto_rawDescOnce sync.Once
	file_hexes_v1_event_payload_proto_rawDescData = file_hexes_v1_event_payload_proto_rawDesc
)

func file_hexes_v1_event_payload_proto_rawDescGZIP() []byte {
	file_hexes_v1_event_payload_proto_rawDescOnce.Do(func() {
		file_hexes_v1_event_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_hexes_v1_event_payload_proto_rawDescData)
	})
	return file_hexes_v1_event_payload_proto_rawDescData
}

var file_hexes_v1_event_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hexes_v1_event_payload_proto_goTypes = []interface{}{
	(*PayloadChat)(nil), // 0: hexes.v1.PayloadChat
}
var file_hexes_v1_event_payload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hexes_v1_event_payload_proto_init() }
func file_hexes_v1_event_payload_proto_init() {
	if File_hexes_v1_event_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hexes_v1_event_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadChat); i {
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
			RawDescriptor: file_hexes_v1_event_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hexes_v1_event_payload_proto_goTypes,
		DependencyIndexes: file_hexes_v1_event_payload_proto_depIdxs,
		MessageInfos:      file_hexes_v1_event_payload_proto_msgTypes,
	}.Build()
	File_hexes_v1_event_payload_proto = out.File
	file_hexes_v1_event_payload_proto_rawDesc = nil
	file_hexes_v1_event_payload_proto_goTypes = nil
	file_hexes_v1_event_payload_proto_depIdxs = nil
}
