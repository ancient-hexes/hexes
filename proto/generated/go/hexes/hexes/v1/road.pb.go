// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: hexes/v1/road.proto

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

type Road_Type int32

const (
	Road_TYPE_UNSPECIFIED     Road_Type = 0
	Road_TYPE_DIRT            Road_Type = 1
	Road_TYPE_GRAVEL          Road_Type = 2
	Road_TYPE_COBBLESTONE     Road_Type = 3
	Road_TYPE_FAVORABLE_WINDS Road_Type = 4
)

// Enum value maps for Road_Type.
var (
	Road_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_DIRT",
		2: "TYPE_GRAVEL",
		3: "TYPE_COBBLESTONE",
		4: "TYPE_FAVORABLE_WINDS",
	}
	Road_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":     0,
		"TYPE_DIRT":            1,
		"TYPE_GRAVEL":          2,
		"TYPE_COBBLESTONE":     3,
		"TYPE_FAVORABLE_WINDS": 4,
	}
)

func (x Road_Type) Enum() *Road_Type {
	p := new(Road_Type)
	*p = x
	return p
}

func (x Road_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Road_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_hexes_v1_road_proto_enumTypes[0].Descriptor()
}

func (Road_Type) Type() protoreflect.EnumType {
	return &file_hexes_v1_road_proto_enumTypes[0]
}

func (x Road_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Road_Type.Descriptor instead.
func (Road_Type) EnumDescriptor() ([]byte, []int) {
	return file_hexes_v1_road_proto_rawDescGZIP(), []int{0, 0}
}

type Road struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Road_Type `protobuf:"varint,1,opt,name=type,proto3,enum=hexes.v1.Road_Type" json:"type,omitempty"`
	// Amount of move points required to move from adjacent tile.
	BaseCost uint32 `protobuf:"varint,2,opt,name=base_cost,json=baseCost,proto3" json:"base_cost,omitempty"`
	// `base_cost` * 1.41 (rounded)
	DiagonalCost uint32 `protobuf:"varint,3,opt,name=diagonal_cost,json=diagonalCost,proto3" json:"diagonal_cost,omitempty"`
}

func (x *Road) Reset() {
	*x = Road{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hexes_v1_road_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Road) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Road) ProtoMessage() {}

func (x *Road) ProtoReflect() protoreflect.Message {
	mi := &file_hexes_v1_road_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Road.ProtoReflect.Descriptor instead.
func (*Road) Descriptor() ([]byte, []int) {
	return file_hexes_v1_road_proto_rawDescGZIP(), []int{0}
}

func (x *Road) GetType() Road_Type {
	if x != nil {
		return x.Type
	}
	return Road_TYPE_UNSPECIFIED
}

func (x *Road) GetBaseCost() uint32 {
	if x != nil {
		return x.BaseCost
	}
	return 0
}

func (x *Road) GetDiagonalCost() uint32 {
	if x != nil {
		return x.DiagonalCost
	}
	return 0
}

var File_hexes_v1_road_proto protoreflect.FileDescriptor

var file_hexes_v1_road_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x61, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22,
	0xdf, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x61, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x69, 0x61, 0x67, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x69, 0x61, 0x67, 0x6f, 0x6e, 0x61, 0x6c, 0x43,
	0x6f, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x52, 0x54, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x41, 0x56, 0x45, 0x4c, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x42, 0x42, 0x4c, 0x45,
	0x53, 0x54, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x46, 0x41, 0x56, 0x4f, 0x52, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x53, 0x10,
	0x04, 0x42, 0x72, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x42, 0x09, 0x52, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x16,
	0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x68,
	0x65, 0x78, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x48,
	0x65, 0x78, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x48, 0x65, 0x78, 0x65, 0x73, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x14, 0x48, 0x65, 0x78, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x48, 0x65, 0x78, 0x65,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hexes_v1_road_proto_rawDescOnce sync.Once
	file_hexes_v1_road_proto_rawDescData = file_hexes_v1_road_proto_rawDesc
)

func file_hexes_v1_road_proto_rawDescGZIP() []byte {
	file_hexes_v1_road_proto_rawDescOnce.Do(func() {
		file_hexes_v1_road_proto_rawDescData = protoimpl.X.CompressGZIP(file_hexes_v1_road_proto_rawDescData)
	})
	return file_hexes_v1_road_proto_rawDescData
}

var file_hexes_v1_road_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hexes_v1_road_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hexes_v1_road_proto_goTypes = []interface{}{
	(Road_Type)(0), // 0: hexes.v1.Road.Type
	(*Road)(nil),   // 1: hexes.v1.Road
}
var file_hexes_v1_road_proto_depIdxs = []int32{
	0, // 0: hexes.v1.Road.type:type_name -> hexes.v1.Road.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hexes_v1_road_proto_init() }
func file_hexes_v1_road_proto_init() {
	if File_hexes_v1_road_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hexes_v1_road_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Road); i {
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
			RawDescriptor: file_hexes_v1_road_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hexes_v1_road_proto_goTypes,
		DependencyIndexes: file_hexes_v1_road_proto_depIdxs,
		EnumInfos:         file_hexes_v1_road_proto_enumTypes,
		MessageInfos:      file_hexes_v1_road_proto_msgTypes,
	}.Build()
	File_hexes_v1_road_proto = out.File
	file_hexes_v1_road_proto_rawDesc = nil
	file_hexes_v1_road_proto_goTypes = nil
	file_hexes_v1_road_proto_depIdxs = nil
}