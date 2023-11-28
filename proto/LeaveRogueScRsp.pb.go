// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: LeaveRogueScRsp.proto

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

type LeaveRogueScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode   uint32      `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	RogueInfo *RogueInfo  `protobuf:"bytes,5,opt,name=rogue_info,json=rogueInfo,proto3" json:"rogue_info,omitempty"`
	Scene     *SceneInfo  `protobuf:"bytes,3,opt,name=scene,proto3" json:"scene,omitempty"`
	Lineup    *LineupInfo `protobuf:"bytes,13,opt,name=lineup,proto3" json:"lineup,omitempty"`
}

func (x *LeaveRogueScRsp) Reset() {
	*x = LeaveRogueScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LeaveRogueScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRogueScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRogueScRsp) ProtoMessage() {}

func (x *LeaveRogueScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_LeaveRogueScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRogueScRsp.ProtoReflect.Descriptor instead.
func (*LeaveRogueScRsp) Descriptor() ([]byte, []int) {
	return file_LeaveRogueScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *LeaveRogueScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *LeaveRogueScRsp) GetRogueInfo() *RogueInfo {
	if x != nil {
		return x.RogueInfo
	}
	return nil
}

func (x *LeaveRogueScRsp) GetScene() *SceneInfo {
	if x != nil {
		return x.Scene
	}
	return nil
}

func (x *LeaveRogueScRsp) GetLineup() *LineupInfo {
	if x != nil {
		return x.Lineup
	}
	return nil
}

var File_LeaveRogueScRsp_proto protoreflect.FileDescriptor

var file_LeaveRogueScRsp_proto_rawDesc = []byte{
	0x0a, 0x15, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x0f,
	0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x72, 0x6f, 0x67, 0x75, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LeaveRogueScRsp_proto_rawDescOnce sync.Once
	file_LeaveRogueScRsp_proto_rawDescData = file_LeaveRogueScRsp_proto_rawDesc
)

func file_LeaveRogueScRsp_proto_rawDescGZIP() []byte {
	file_LeaveRogueScRsp_proto_rawDescOnce.Do(func() {
		file_LeaveRogueScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_LeaveRogueScRsp_proto_rawDescData)
	})
	return file_LeaveRogueScRsp_proto_rawDescData
}

var file_LeaveRogueScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LeaveRogueScRsp_proto_goTypes = []interface{}{
	(*LeaveRogueScRsp)(nil), // 0: LeaveRogueScRsp
	(*RogueInfo)(nil),       // 1: RogueInfo
	(*SceneInfo)(nil),       // 2: SceneInfo
	(*LineupInfo)(nil),      // 3: LineupInfo
}
var file_LeaveRogueScRsp_proto_depIdxs = []int32{
	1, // 0: LeaveRogueScRsp.rogue_info:type_name -> RogueInfo
	2, // 1: LeaveRogueScRsp.scene:type_name -> SceneInfo
	3, // 2: LeaveRogueScRsp.lineup:type_name -> LineupInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_LeaveRogueScRsp_proto_init() }
func file_LeaveRogueScRsp_proto_init() {
	if File_LeaveRogueScRsp_proto != nil {
		return
	}
	file_RogueInfo_proto_init()
	file_SceneInfo_proto_init()
	file_LineupInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LeaveRogueScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRogueScRsp); i {
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
			RawDescriptor: file_LeaveRogueScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LeaveRogueScRsp_proto_goTypes,
		DependencyIndexes: file_LeaveRogueScRsp_proto_depIdxs,
		MessageInfos:      file_LeaveRogueScRsp_proto_msgTypes,
	}.Build()
	File_LeaveRogueScRsp_proto = out.File
	file_LeaveRogueScRsp_proto_rawDesc = nil
	file_LeaveRogueScRsp_proto_goTypes = nil
	file_LeaveRogueScRsp_proto_depIdxs = nil
}
