// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: hexes/v1/tile.proto

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

type Tile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A coordinate relative to map's bottom left corner.
	Coordinate *Coordinate `protobuf:"bytes,1,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	// Reference to `Terrain.type`.
	Terrain Terrain_Type `protobuf:"varint,2,opt,name=terrain,proto3,enum=hexes.v1.Terrain_Type" json:"terrain,omitempty"`
	// Reference to `Road.type`.
	Road Road_Type `protobuf:"varint,3,opt,name=road,proto3,enum=hexes.v1.Road_Type" json:"road,omitempty"`
	// A list of players who can view & inspect the tile.
	VisibleTo []Color `protobuf:"varint,4,rep,packed,name=visible_to,json=visibleTo,proto3,enum=hexes.v1.Color" json:"visible_to,omitempty"`
}

func (x *Tile) Reset() {
	*x = Tile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hexes_v1_tile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tile) ProtoMessage() {}

func (x *Tile) ProtoReflect() protoreflect.Message {
	mi := &file_hexes_v1_tile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tile.ProtoReflect.Descriptor instead.
func (*Tile) Descriptor() ([]byte, []int) {
	return file_hexes_v1_tile_proto_rawDescGZIP(), []int{0}
}

func (x *Tile) GetCoordinate() *Coordinate {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

func (x *Tile) GetTerrain() Terrain_Type {
	if x != nil {
		return x.Terrain
	}
	return Terrain_TYPE_UNSPECIFIED
}

func (x *Tile) GetRoad() Road_Type {
	if x != nil {
		return x.Road
	}
	return Road_TYPE_UNSPECIFIED
}

func (x *Tile) GetVisibleTo() []Color {
	if x != nil {
		return x.VisibleTo
	}
	return nil
}

var File_hexes_v1_tile_proto protoreflect.FileDescriptor

var file_hexes_v1_tile_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x14, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x61, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01,
	0x0a, 0x04, 0x54, 0x69, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x65, 0x78,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x07,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x68, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x74, 0x65, 0x72, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x27,
	0x0a, 0x04, 0x72, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x68,
	0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x72, 0x6f, 0x61, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x68, 0x65,
	0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x09, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x42, 0x8b, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x78, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x54, 0x69, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x78, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x68,
	0x65, 0x78, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x48,
	0x65, 0x78, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x48, 0x65, 0x78, 0x65, 0x73, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x14, 0x48, 0x65, 0x78, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x48, 0x65, 0x78, 0x65,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hexes_v1_tile_proto_rawDescOnce sync.Once
	file_hexes_v1_tile_proto_rawDescData = file_hexes_v1_tile_proto_rawDesc
)

func file_hexes_v1_tile_proto_rawDescGZIP() []byte {
	file_hexes_v1_tile_proto_rawDescOnce.Do(func() {
		file_hexes_v1_tile_proto_rawDescData = protoimpl.X.CompressGZIP(file_hexes_v1_tile_proto_rawDescData)
	})
	return file_hexes_v1_tile_proto_rawDescData
}

var file_hexes_v1_tile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hexes_v1_tile_proto_goTypes = []interface{}{
	(*Tile)(nil),       // 0: hexes.v1.Tile
	(*Coordinate)(nil), // 1: hexes.v1.Coordinate
	(Terrain_Type)(0),  // 2: hexes.v1.Terrain.Type
	(Road_Type)(0),     // 3: hexes.v1.Road.Type
	(Color)(0),         // 4: hexes.v1.Color
}
var file_hexes_v1_tile_proto_depIdxs = []int32{
	1, // 0: hexes.v1.Tile.coordinate:type_name -> hexes.v1.Coordinate
	2, // 1: hexes.v1.Tile.terrain:type_name -> hexes.v1.Terrain.Type
	3, // 2: hexes.v1.Tile.road:type_name -> hexes.v1.Road.Type
	4, // 3: hexes.v1.Tile.visible_to:type_name -> hexes.v1.Color
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_hexes_v1_tile_proto_init() }
func file_hexes_v1_tile_proto_init() {
	if File_hexes_v1_tile_proto != nil {
		return
	}
	file_hexes_v1_color_proto_init()
	file_hexes_v1_coordinate_proto_init()
	file_hexes_v1_road_proto_init()
	file_hexes_v1_terrain_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_hexes_v1_tile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tile); i {
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
			RawDescriptor: file_hexes_v1_tile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hexes_v1_tile_proto_goTypes,
		DependencyIndexes: file_hexes_v1_tile_proto_depIdxs,
		MessageInfos:      file_hexes_v1_tile_proto_msgTypes,
	}.Build()
	File_hexes_v1_tile_proto = out.File
	file_hexes_v1_tile_proto_rawDesc = nil
	file_hexes_v1_tile_proto_goTypes = nil
	file_hexes_v1_tile_proto_depIdxs = nil
}
