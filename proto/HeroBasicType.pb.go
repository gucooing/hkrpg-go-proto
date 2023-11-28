// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: HeroBasicType.proto

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

type HeroBasicType int32

const (
	HeroBasicType_None        HeroBasicType = 0
	HeroBasicType_BoyWarrior  HeroBasicType = 8001
	HeroBasicType_GirlWarrior HeroBasicType = 8002
	HeroBasicType_BoyKnight   HeroBasicType = 8003
	HeroBasicType_GirlKnight  HeroBasicType = 8004
	HeroBasicType_BoyRogue    HeroBasicType = 8005
	HeroBasicType_GirlRogue   HeroBasicType = 8006
	HeroBasicType_BoyMage     HeroBasicType = 8007
	HeroBasicType_GirlMage    HeroBasicType = 8008
	HeroBasicType_BoyShaman   HeroBasicType = 8009
	HeroBasicType_GirlShaman  HeroBasicType = 8010
	HeroBasicType_BoyWarlock  HeroBasicType = 8011
	HeroBasicType_GirlWarlock HeroBasicType = 8012
	HeroBasicType_BoyPriest   HeroBasicType = 8013
	HeroBasicType_GirlPriest  HeroBasicType = 8014
)

// Enum value maps for HeroBasicType.
var (
	HeroBasicType_name = map[int32]string{
		0:    "None",
		8001: "BoyWarrior",
		8002: "GirlWarrior",
		8003: "BoyKnight",
		8004: "GirlKnight",
		8005: "BoyRogue",
		8006: "GirlRogue",
		8007: "BoyMage",
		8008: "GirlMage",
		8009: "BoyShaman",
		8010: "GirlShaman",
		8011: "BoyWarlock",
		8012: "GirlWarlock",
		8013: "BoyPriest",
		8014: "GirlPriest",
	}
	HeroBasicType_value = map[string]int32{
		"None":        0,
		"BoyWarrior":  8001,
		"GirlWarrior": 8002,
		"BoyKnight":   8003,
		"GirlKnight":  8004,
		"BoyRogue":    8005,
		"GirlRogue":   8006,
		"BoyMage":     8007,
		"GirlMage":    8008,
		"BoyShaman":   8009,
		"GirlShaman":  8010,
		"BoyWarlock":  8011,
		"GirlWarlock": 8012,
		"BoyPriest":   8013,
		"GirlPriest":  8014,
	}
)

func (x HeroBasicType) Enum() *HeroBasicType {
	p := new(HeroBasicType)
	*p = x
	return p
}

func (x HeroBasicType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HeroBasicType) Descriptor() protoreflect.EnumDescriptor {
	return file_HeroBasicType_proto_enumTypes[0].Descriptor()
}

func (HeroBasicType) Type() protoreflect.EnumType {
	return &file_HeroBasicType_proto_enumTypes[0]
}

func (x HeroBasicType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HeroBasicType.Descriptor instead.
func (HeroBasicType) EnumDescriptor() ([]byte, []int) {
	return file_HeroBasicType_proto_rawDescGZIP(), []int{0}
}

var File_HeroBasicType_proto protoreflect.FileDescriptor

var file_HeroBasicType_proto_rawDesc = []byte{
	0x0a, 0x13, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xfe, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0a, 0x42, 0x6f, 0x79, 0x57, 0x61, 0x72, 0x72, 0x69, 0x6f, 0x72, 0x10,
	0xc1, 0x3e, 0x12, 0x10, 0x0a, 0x0b, 0x47, 0x69, 0x72, 0x6c, 0x57, 0x61, 0x72, 0x72, 0x69, 0x6f,
	0x72, 0x10, 0xc2, 0x3e, 0x12, 0x0e, 0x0a, 0x09, 0x42, 0x6f, 0x79, 0x4b, 0x6e, 0x69, 0x67, 0x68,
	0x74, 0x10, 0xc3, 0x3e, 0x12, 0x0f, 0x0a, 0x0a, 0x47, 0x69, 0x72, 0x6c, 0x4b, 0x6e, 0x69, 0x67,
	0x68, 0x74, 0x10, 0xc4, 0x3e, 0x12, 0x0d, 0x0a, 0x08, 0x42, 0x6f, 0x79, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x10, 0xc5, 0x3e, 0x12, 0x0e, 0x0a, 0x09, 0x47, 0x69, 0x72, 0x6c, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x10, 0xc6, 0x3e, 0x12, 0x0c, 0x0a, 0x07, 0x42, 0x6f, 0x79, 0x4d, 0x61, 0x67, 0x65, 0x10,
	0xc7, 0x3e, 0x12, 0x0d, 0x0a, 0x08, 0x47, 0x69, 0x72, 0x6c, 0x4d, 0x61, 0x67, 0x65, 0x10, 0xc8,
	0x3e, 0x12, 0x0e, 0x0a, 0x09, 0x42, 0x6f, 0x79, 0x53, 0x68, 0x61, 0x6d, 0x61, 0x6e, 0x10, 0xc9,
	0x3e, 0x12, 0x0f, 0x0a, 0x0a, 0x47, 0x69, 0x72, 0x6c, 0x53, 0x68, 0x61, 0x6d, 0x61, 0x6e, 0x10,
	0xca, 0x3e, 0x12, 0x0f, 0x0a, 0x0a, 0x42, 0x6f, 0x79, 0x57, 0x61, 0x72, 0x6c, 0x6f, 0x63, 0x6b,
	0x10, 0xcb, 0x3e, 0x12, 0x10, 0x0a, 0x0b, 0x47, 0x69, 0x72, 0x6c, 0x57, 0x61, 0x72, 0x6c, 0x6f,
	0x63, 0x6b, 0x10, 0xcc, 0x3e, 0x12, 0x0e, 0x0a, 0x09, 0x42, 0x6f, 0x79, 0x50, 0x72, 0x69, 0x65,
	0x73, 0x74, 0x10, 0xcd, 0x3e, 0x12, 0x0f, 0x0a, 0x0a, 0x47, 0x69, 0x72, 0x6c, 0x50, 0x72, 0x69,
	0x65, 0x73, 0x74, 0x10, 0xce, 0x3e, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HeroBasicType_proto_rawDescOnce sync.Once
	file_HeroBasicType_proto_rawDescData = file_HeroBasicType_proto_rawDesc
)

func file_HeroBasicType_proto_rawDescGZIP() []byte {
	file_HeroBasicType_proto_rawDescOnce.Do(func() {
		file_HeroBasicType_proto_rawDescData = protoimpl.X.CompressGZIP(file_HeroBasicType_proto_rawDescData)
	})
	return file_HeroBasicType_proto_rawDescData
}

var file_HeroBasicType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_HeroBasicType_proto_goTypes = []interface{}{
	(HeroBasicType)(0), // 0: HeroBasicType
}
var file_HeroBasicType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HeroBasicType_proto_init() }
func file_HeroBasicType_proto_init() {
	if File_HeroBasicType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HeroBasicType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HeroBasicType_proto_goTypes,
		DependencyIndexes: file_HeroBasicType_proto_depIdxs,
		EnumInfos:         file_HeroBasicType_proto_enumTypes,
	}.Build()
	File_HeroBasicType_proto = out.File
	file_HeroBasicType_proto_rawDesc = nil
	file_HeroBasicType_proto_goTypes = nil
	file_HeroBasicType_proto_depIdxs = nil
}
