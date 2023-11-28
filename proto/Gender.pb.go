// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: Gender.proto

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

type Gender int32

const (
	Gender_GenderNone  Gender = 0
	Gender_GenderMan   Gender = 1
	Gender_GenderWoman Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GenderNone",
		1: "GenderMan",
		2: "GenderWoman",
	}
	Gender_value = map[string]int32{
		"GenderNone":  0,
		"GenderMan":   1,
		"GenderWoman": 2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_Gender_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_Gender_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_Gender_proto_rawDescGZIP(), []int{0}
}

var File_Gender_proto protoreflect.FileDescriptor

var file_Gender_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x38,
	0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x57, 0x6f, 0x6d, 0x61, 0x6e, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Gender_proto_rawDescOnce sync.Once
	file_Gender_proto_rawDescData = file_Gender_proto_rawDesc
)

func file_Gender_proto_rawDescGZIP() []byte {
	file_Gender_proto_rawDescOnce.Do(func() {
		file_Gender_proto_rawDescData = protoimpl.X.CompressGZIP(file_Gender_proto_rawDescData)
	})
	return file_Gender_proto_rawDescData
}

var file_Gender_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Gender_proto_goTypes = []interface{}{
	(Gender)(0), // 0: Gender
}
var file_Gender_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Gender_proto_init() }
func file_Gender_proto_init() {
	if File_Gender_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Gender_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Gender_proto_goTypes,
		DependencyIndexes: file_Gender_proto_depIdxs,
		EnumInfos:         file_Gender_proto_enumTypes,
	}.Build()
	File_Gender_proto = out.File
	file_Gender_proto_rawDesc = nil
	file_Gender_proto_goTypes = nil
	file_Gender_proto_depIdxs = nil
}
