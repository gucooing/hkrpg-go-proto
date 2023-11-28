// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: EnterRogueMapRoomCsReq.proto

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

type EnterRogueMapRoomCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId uint32 `protobuf:"varint,11,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	SiteId uint32 `protobuf:"varint,7,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
}

func (x *EnterRogueMapRoomCsReq) Reset() {
	*x = EnterRogueMapRoomCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnterRogueMapRoomCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterRogueMapRoomCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterRogueMapRoomCsReq) ProtoMessage() {}

func (x *EnterRogueMapRoomCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_EnterRogueMapRoomCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterRogueMapRoomCsReq.ProtoReflect.Descriptor instead.
func (*EnterRogueMapRoomCsReq) Descriptor() ([]byte, []int) {
	return file_EnterRogueMapRoomCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *EnterRogueMapRoomCsReq) GetRoomId() uint32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *EnterRogueMapRoomCsReq) GetSiteId() uint32 {
	if x != nil {
		return x.SiteId
	}
	return 0
}

var File_EnterRogueMapRoomCsReq_proto protoreflect.FileDescriptor

var file_EnterRogueMapRoomCsReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x52,
	0x6f, 0x6f, 0x6d, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a,
	0x0a, 0x16, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x52,
	0x6f, 0x6f, 0x6d, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EnterRogueMapRoomCsReq_proto_rawDescOnce sync.Once
	file_EnterRogueMapRoomCsReq_proto_rawDescData = file_EnterRogueMapRoomCsReq_proto_rawDesc
)

func file_EnterRogueMapRoomCsReq_proto_rawDescGZIP() []byte {
	file_EnterRogueMapRoomCsReq_proto_rawDescOnce.Do(func() {
		file_EnterRogueMapRoomCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterRogueMapRoomCsReq_proto_rawDescData)
	})
	return file_EnterRogueMapRoomCsReq_proto_rawDescData
}

var file_EnterRogueMapRoomCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnterRogueMapRoomCsReq_proto_goTypes = []interface{}{
	(*EnterRogueMapRoomCsReq)(nil), // 0: EnterRogueMapRoomCsReq
}
var file_EnterRogueMapRoomCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EnterRogueMapRoomCsReq_proto_init() }
func file_EnterRogueMapRoomCsReq_proto_init() {
	if File_EnterRogueMapRoomCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EnterRogueMapRoomCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterRogueMapRoomCsReq); i {
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
			RawDescriptor: file_EnterRogueMapRoomCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterRogueMapRoomCsReq_proto_goTypes,
		DependencyIndexes: file_EnterRogueMapRoomCsReq_proto_depIdxs,
		MessageInfos:      file_EnterRogueMapRoomCsReq_proto_msgTypes,
	}.Build()
	File_EnterRogueMapRoomCsReq_proto = out.File
	file_EnterRogueMapRoomCsReq_proto_rawDesc = nil
	file_EnterRogueMapRoomCsReq_proto_goTypes = nil
	file_EnterRogueMapRoomCsReq_proto_depIdxs = nil
}
