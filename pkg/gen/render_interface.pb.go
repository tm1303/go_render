// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: render_interface.proto

package gen

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

type SetBackgroundColourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Colour string `protobuf:"bytes,1,opt,name=colour,proto3" json:"colour,omitempty"`
}

func (x *SetBackgroundColourRequest) Reset() {
	*x = SetBackgroundColourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_render_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBackgroundColourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBackgroundColourRequest) ProtoMessage() {}

func (x *SetBackgroundColourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_render_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBackgroundColourRequest.ProtoReflect.Descriptor instead.
func (*SetBackgroundColourRequest) Descriptor() ([]byte, []int) {
	return file_render_interface_proto_rawDescGZIP(), []int{0}
}

func (x *SetBackgroundColourRequest) GetColour() string {
	if x != nil {
		return x.Colour
	}
	return ""
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_render_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_render_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_render_interface_proto_rawDescGZIP(), []int{1}
}

type RenderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RenderRequest) Reset() {
	*x = RenderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_render_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderRequest) ProtoMessage() {}

func (x *RenderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_render_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderRequest.ProtoReflect.Descriptor instead.
func (*RenderRequest) Descriptor() ([]byte, []int) {
	return file_render_interface_proto_rawDescGZIP(), []int{2}
}

type RenderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunk []byte `protobuf:"bytes,1,opt,name=fileChunk,proto3" json:"fileChunk,omitempty"`
}

func (x *RenderResponse) Reset() {
	*x = RenderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_render_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderResponse) ProtoMessage() {}

func (x *RenderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_render_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderResponse.ProtoReflect.Descriptor instead.
func (*RenderResponse) Descriptor() ([]byte, []int) {
	return file_render_interface_proto_rawDescGZIP(), []int{3}
}

func (x *RenderResponse) GetFileChunk() []byte {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

var File_render_interface_proto protoreflect.FileDescriptor

var file_render_interface_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x65, 0x6e, 0x22, 0x34, 0x0a,
	0x1a, 0x53, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f,
	0x6c, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c,
	0x6f, 0x75, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x0e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0x96, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x13, 0x53, 0x65, 0x74,
	0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72,
	0x12, 0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x52, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x06,
	0x5a, 0x04, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_render_interface_proto_rawDescOnce sync.Once
	file_render_interface_proto_rawDescData = file_render_interface_proto_rawDesc
)

func file_render_interface_proto_rawDescGZIP() []byte {
	file_render_interface_proto_rawDescOnce.Do(func() {
		file_render_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_render_interface_proto_rawDescData)
	})
	return file_render_interface_proto_rawDescData
}

var file_render_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_render_interface_proto_goTypes = []interface{}{
	(*SetBackgroundColourRequest)(nil), // 0: gen.SetBackgroundColourRequest
	(*EmptyResponse)(nil),              // 1: gen.EmptyResponse
	(*RenderRequest)(nil),              // 2: gen.RenderRequest
	(*RenderResponse)(nil),             // 3: gen.RenderResponse
}
var file_render_interface_proto_depIdxs = []int32{
	0, // 0: gen.RenderInterface.SetBackgroundColour:input_type -> gen.SetBackgroundColourRequest
	2, // 1: gen.RenderInterface.Render:input_type -> gen.RenderRequest
	1, // 2: gen.RenderInterface.SetBackgroundColour:output_type -> gen.EmptyResponse
	3, // 3: gen.RenderInterface.Render:output_type -> gen.RenderResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_render_interface_proto_init() }
func file_render_interface_proto_init() {
	if File_render_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_render_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBackgroundColourRequest); i {
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
		file_render_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_render_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderRequest); i {
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
		file_render_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderResponse); i {
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
			RawDescriptor: file_render_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_render_interface_proto_goTypes,
		DependencyIndexes: file_render_interface_proto_depIdxs,
		MessageInfos:      file_render_interface_proto_msgTypes,
	}.Build()
	File_render_interface_proto = out.File
	file_render_interface_proto_rawDesc = nil
	file_render_interface_proto_goTypes = nil
	file_render_interface_proto_depIdxs = nil
}
