// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SyncRogueMiracleSelectInfoScNotify.proto

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

type SyncRogueMiracleSelectInfoScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MiracleSelectInfo *RogueMiracleSelectInfo `protobuf:"bytes,7,opt,name=miracle_select_info,json=miracleSelectInfo,proto3" json:"miracle_select_info,omitempty"`
}

func (x *SyncRogueMiracleSelectInfoScNotify) Reset() {
	*x = SyncRogueMiracleSelectInfoScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SyncRogueMiracleSelectInfoScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRogueMiracleSelectInfoScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRogueMiracleSelectInfoScNotify) ProtoMessage() {}

func (x *SyncRogueMiracleSelectInfoScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SyncRogueMiracleSelectInfoScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRogueMiracleSelectInfoScNotify.ProtoReflect.Descriptor instead.
func (*SyncRogueMiracleSelectInfoScNotify) Descriptor() ([]byte, []int) {
	return file_SyncRogueMiracleSelectInfoScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRogueMiracleSelectInfoScNotify) GetMiracleSelectInfo() *RogueMiracleSelectInfo {
	if x != nil {
		return x.MiracleSelectInfo
	}
	return nil
}

var File_SyncRogueMiracleSelectInfoScNotify_proto protoreflect.FileDescriptor

var file_SyncRogueMiracleSelectInfoScNotify_proto_rawDesc = []byte{
	0x0a, 0x28, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x22, 0x53, 0x79, 0x6e, 0x63,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x47,
	0x0a, 0x13, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SyncRogueMiracleSelectInfoScNotify_proto_rawDescOnce sync.Once
	file_SyncRogueMiracleSelectInfoScNotify_proto_rawDescData = file_SyncRogueMiracleSelectInfoScNotify_proto_rawDesc
)

func file_SyncRogueMiracleSelectInfoScNotify_proto_rawDescGZIP() []byte {
	file_SyncRogueMiracleSelectInfoScNotify_proto_rawDescOnce.Do(func() {
		file_SyncRogueMiracleSelectInfoScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SyncRogueMiracleSelectInfoScNotify_proto_rawDescData)
	})
	return file_SyncRogueMiracleSelectInfoScNotify_proto_rawDescData
}

var file_SyncRogueMiracleSelectInfoScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SyncRogueMiracleSelectInfoScNotify_proto_goTypes = []interface{}{
	(*SyncRogueMiracleSelectInfoScNotify)(nil), // 0: SyncRogueMiracleSelectInfoScNotify
	(*RogueMiracleSelectInfo)(nil),             // 1: RogueMiracleSelectInfo
}
var file_SyncRogueMiracleSelectInfoScNotify_proto_depIdxs = []int32{
	1, // 0: SyncRogueMiracleSelectInfoScNotify.miracle_select_info:type_name -> RogueMiracleSelectInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SyncRogueMiracleSelectInfoScNotify_proto_init() }
func file_SyncRogueMiracleSelectInfoScNotify_proto_init() {
	if File_SyncRogueMiracleSelectInfoScNotify_proto != nil {
		return
	}
	file_RogueMiracleSelectInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SyncRogueMiracleSelectInfoScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRogueMiracleSelectInfoScNotify); i {
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
			RawDescriptor: file_SyncRogueMiracleSelectInfoScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SyncRogueMiracleSelectInfoScNotify_proto_goTypes,
		DependencyIndexes: file_SyncRogueMiracleSelectInfoScNotify_proto_depIdxs,
		MessageInfos:      file_SyncRogueMiracleSelectInfoScNotify_proto_msgTypes,
	}.Build()
	File_SyncRogueMiracleSelectInfoScNotify_proto = out.File
	file_SyncRogueMiracleSelectInfoScNotify_proto_rawDesc = nil
	file_SyncRogueMiracleSelectInfoScNotify_proto_goTypes = nil
	file_SyncRogueMiracleSelectInfoScNotify_proto_depIdxs = nil
}
