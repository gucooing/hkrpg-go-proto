// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ArchiveData.proto

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

type ArchiveData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArchiveAvatarIdList    []uint32 `protobuf:"varint,4,rep,packed,name=archive_avatar_id_list,json=archiveAvatarIdList,proto3" json:"archive_avatar_id_list,omitempty"`
	ArchiveEquipmentIdList []uint32 `protobuf:"varint,3,rep,packed,name=archive_equipment_id_list,json=archiveEquipmentIdList,proto3" json:"archive_equipment_id_list,omitempty"`
}

func (x *ArchiveData) Reset() {
	*x = ArchiveData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ArchiveData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveData) ProtoMessage() {}

func (x *ArchiveData) ProtoReflect() protoreflect.Message {
	mi := &file_ArchiveData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveData.ProtoReflect.Descriptor instead.
func (*ArchiveData) Descriptor() ([]byte, []int) {
	return file_ArchiveData_proto_rawDescGZIP(), []int{0}
}

func (x *ArchiveData) GetArchiveAvatarIdList() []uint32 {
	if x != nil {
		return x.ArchiveAvatarIdList
	}
	return nil
}

func (x *ArchiveData) GetArchiveEquipmentIdList() []uint32 {
	if x != nil {
		return x.ArchiveEquipmentIdList
	}
	return nil
}

var File_ArchiveData_proto protoreflect.FileDescriptor

var file_ArchiveData_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0b, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x33, 0x0a, 0x16, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x13, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x5f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x16, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ArchiveData_proto_rawDescOnce sync.Once
	file_ArchiveData_proto_rawDescData = file_ArchiveData_proto_rawDesc
)

func file_ArchiveData_proto_rawDescGZIP() []byte {
	file_ArchiveData_proto_rawDescOnce.Do(func() {
		file_ArchiveData_proto_rawDescData = protoimpl.X.CompressGZIP(file_ArchiveData_proto_rawDescData)
	})
	return file_ArchiveData_proto_rawDescData
}

var file_ArchiveData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ArchiveData_proto_goTypes = []interface{}{
	(*ArchiveData)(nil), // 0: ArchiveData
}
var file_ArchiveData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ArchiveData_proto_init() }
func file_ArchiveData_proto_init() {
	if File_ArchiveData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ArchiveData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveData); i {
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
			RawDescriptor: file_ArchiveData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ArchiveData_proto_goTypes,
		DependencyIndexes: file_ArchiveData_proto_depIdxs,
		MessageInfos:      file_ArchiveData_proto_msgTypes,
	}.Build()
	File_ArchiveData_proto = out.File
	file_ArchiveData_proto_rawDesc = nil
	file_ArchiveData_proto_goTypes = nil
	file_ArchiveData_proto_depIdxs = nil
}
