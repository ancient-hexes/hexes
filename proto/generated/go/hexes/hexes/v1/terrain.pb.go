// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: hexes/v1/terrain.proto

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

type Terrain_Type int32

const (
	Terrain_TYPE_UNSPECIFIED  Terrain_Type = 0
	Terrain_TYPE_GRASS        Terrain_Type = 1
	Terrain_TYPE_HIGHLANDS    Terrain_Type = 2
	Terrain_TYPE_DIRT         Terrain_Type = 3
	Terrain_TYPE_LAVA         Terrain_Type = 4
	Terrain_TYPE_SUBTERRANEAN Terrain_Type = 5
	Terrain_TYPE_ROCK         Terrain_Type = 6
	Terrain_TYPE_ROUGH        Terrain_Type = 7
	Terrain_TYPE_WASTELAND    Terrain_Type = 8
	Terrain_TYPE_SAND         Terrain_Type = 9
	Terrain_TYPE_SNOW         Terrain_Type = 10
	Terrain_TYPE_SWAMP        Terrain_Type = 11
	Terrain_TYPE_WATER        Terrain_Type = 12
)

// Enum value maps for Terrain_Type.
var (
	Terrain_Type_name = map[int32]string{
		0:  "TYPE_UNSPECIFIED",
		1:  "TYPE_GRASS",
		2:  "TYPE_HIGHLANDS",
		3:  "TYPE_DIRT",
		4:  "TYPE_LAVA",
		5:  "TYPE_SUBTERRANEAN",
		6:  "TYPE_ROCK",
		7:  "TYPE_ROUGH",
		8:  "TYPE_WASTELAND",
		9:  "TYPE_SAND",
		10: "TYPE_SNOW",
		11: "TYPE_SWAMP",
		12: "TYPE_WATER",
	}
	Terrain_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":  0,
		"TYPE_GRASS":        1,
		"TYPE_HIGHLANDS":    2,
		"TYPE_DIRT":         3,
		"TYPE_LAVA":         4,
		"TYPE_SUBTERRANEAN": 5,
		"TYPE_ROCK":         6,
		"TYPE_ROUGH":        7,
		"TYPE_WASTELAND":    8,
		"TYPE_SAND":         9,
		"TYPE_SNOW":         10,
		"TYPE_SWAMP":        11,
		"TYPE_WATER":        12,
	}
)

func (x Terrain_Type) Enum() *Terrain_Type {
	p := new(Terrain_Type)
	*p = x
	return p
}

func (x Terrain_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Terrain_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_hexes_v1_terrain_proto_enumTypes[0].Descriptor()
}

func (Terrain_Type) Type() protoreflect.EnumType {
	return &file_hexes_v1_terrain_proto_enumTypes[0]
}

func (x Terrain_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Terrain_Type.Descriptor instead.
func (Terrain_Type) EnumDescriptor() ([]byte, []int) {
	return file_hexes_v1_terrain_proto_rawDescGZIP(), []int{0, 0}
}

type Terrain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Terrain_Type `protobuf:"varint,1,opt,name=type,proto3,enum=hexes.v1.Terrain_Type" json:"type,omitempty"`
	// A movement penalty for a given terrain without a road.
	MovementPenalty float64 `protobuf:"fixed64,2,opt,name=movement_penalty,json=movementPenalty,proto3" json:"movement_penalty,omitempty"`
}

func (x *Terrain) Reset() {
	*x = Terrain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hexes_v1_terrain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Terrain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Terrain) ProtoMessage() {}

func (x *Terrain) ProtoReflect() protoreflect.Message {
	mi := &file_hexes_v1_terrain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Terrain.ProtoReflect.Descriptor instead.
func (*Terrain) Descriptor() ([]byte, []int) {
	return file_hexes_v1_terrain_proto_rawDescGZIP(), []int{0}
}

func (x *Terrain) GetType() Terrain_Type {
	if x != nil {
		return x.Type
	}
	return Terrain_TYPE_UNSPECIFIED
}

func (x *Terrain) GetMovementPenalty() float64 {
	if x != nil {
		return x.MovementPenalty
	}
	return 0
}

var File_hexes_v1_terrain_proto protoreflect.FileDescriptor

var file_hexes_v1_terrain_proto_rawDesc = []byte{
	0x0a, 0x16, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x22, 0xc9, 0x02, 0x0a, 0x07, 0x54, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x2a,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x68,
	0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x6f,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65,
	0x6e, 0x61, 0x6c, 0x74, 0x79, 0x22, 0xe6, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x41,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x49, 0x47,
	0x48, 0x4c, 0x41, 0x4e, 0x44, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x49, 0x52, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4c, 0x41, 0x56, 0x41, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x55, 0x42, 0x54, 0x45, 0x52, 0x52, 0x41, 0x4e, 0x45, 0x41, 0x4e, 0x10, 0x05, 0x12, 0x0d, 0x0a,
	0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x4f, 0x43, 0x4b, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x4f, 0x55, 0x47, 0x48, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x41, 0x53, 0x54, 0x45, 0x4c, 0x41, 0x4e, 0x44, 0x10, 0x08,
	0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x41, 0x4e, 0x44, 0x10, 0x09, 0x12,
	0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4e, 0x4f, 0x57, 0x10, 0x0a, 0x12, 0x0e,
	0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x57, 0x41, 0x4d, 0x50, 0x10, 0x0b, 0x12, 0x0e,
	0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x41, 0x54, 0x45, 0x52, 0x10, 0x0c, 0x42, 0x75,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x54, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x16,
	0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x68,
	0x65, 0x78, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x48,
	0x65, 0x78, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x48, 0x65, 0x78, 0x65, 0x73, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x14, 0x48, 0x65, 0x78, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x48, 0x65, 0x78, 0x65,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hexes_v1_terrain_proto_rawDescOnce sync.Once
	file_hexes_v1_terrain_proto_rawDescData = file_hexes_v1_terrain_proto_rawDesc
)

func file_hexes_v1_terrain_proto_rawDescGZIP() []byte {
	file_hexes_v1_terrain_proto_rawDescOnce.Do(func() {
		file_hexes_v1_terrain_proto_rawDescData = protoimpl.X.CompressGZIP(file_hexes_v1_terrain_proto_rawDescData)
	})
	return file_hexes_v1_terrain_proto_rawDescData
}

var file_hexes_v1_terrain_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hexes_v1_terrain_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hexes_v1_terrain_proto_goTypes = []interface{}{
	(Terrain_Type)(0), // 0: hexes.v1.Terrain.Type
	(*Terrain)(nil),   // 1: hexes.v1.Terrain
}
var file_hexes_v1_terrain_proto_depIdxs = []int32{
	0, // 0: hexes.v1.Terrain.type:type_name -> hexes.v1.Terrain.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hexes_v1_terrain_proto_init() }
func file_hexes_v1_terrain_proto_init() {
	if File_hexes_v1_terrain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hexes_v1_terrain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Terrain); i {
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
			RawDescriptor: file_hexes_v1_terrain_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hexes_v1_terrain_proto_goTypes,
		DependencyIndexes: file_hexes_v1_terrain_proto_depIdxs,
		EnumInfos:         file_hexes_v1_terrain_proto_enumTypes,
		MessageInfos:      file_hexes_v1_terrain_proto_msgTypes,
	}.Build()
	File_hexes_v1_terrain_proto = out.File
	file_hexes_v1_terrain_proto_rawDesc = nil
	file_hexes_v1_terrain_proto_goTypes = nil
	file_hexes_v1_terrain_proto_depIdxs = nil
}