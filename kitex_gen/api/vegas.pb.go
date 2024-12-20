// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: vegas.proto

package api

import (
	context "context"
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

type DrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Eh   string `protobuf:"bytes,2,opt,name=eh,proto3" json:"eh,omitempty"`
}

func (x *DrawRequest) Reset() {
	*x = DrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vegas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawRequest) ProtoMessage() {}

func (x *DrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vegas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawRequest.ProtoReflect.Descriptor instead.
func (*DrawRequest) Descriptor() ([]byte, []int) {
	return file_vegas_proto_rawDescGZIP(), []int{0}
}

func (x *DrawRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *DrawRequest) GetEh() string {
	if x != nil {
		return x.Eh
	}
	return ""
}

type DrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base   *BaseResponse `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Rights []*Right      `protobuf:"bytes,2,rep,name=rights,proto3" json:"rights,omitempty"`
}

func (x *DrawResponse) Reset() {
	*x = DrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vegas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawResponse) ProtoMessage() {}

func (x *DrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vegas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawResponse.ProtoReflect.Descriptor instead.
func (*DrawResponse) Descriptor() ([]byte, []int) {
	return file_vegas_proto_rawDescGZIP(), []int{1}
}

func (x *DrawResponse) GetBase() *BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *DrawResponse) GetRights() []*Right {
	if x != nil {
		return x.Rights
	}
	return nil
}

type ShowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Eh   string `protobuf:"bytes,2,opt,name=eh,proto3" json:"eh,omitempty"`
}

func (x *ShowRequest) Reset() {
	*x = ShowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vegas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowRequest) ProtoMessage() {}

func (x *ShowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vegas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowRequest.ProtoReflect.Descriptor instead.
func (*ShowRequest) Descriptor() ([]byte, []int) {
	return file_vegas_proto_rawDescGZIP(), []int{2}
}

func (x *ShowRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ShowRequest) GetEh() string {
	if x != nil {
		return x.Eh
	}
	return ""
}

type ShowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *BaseResponse `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *ShowResponse) Reset() {
	*x = ShowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vegas_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowResponse) ProtoMessage() {}

func (x *ShowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vegas_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowResponse.ProtoReflect.Descriptor instead.
func (*ShowResponse) Descriptor() ([]byte, []int) {
	return file_vegas_proto_rawDescGZIP(), []int{3}
}

func (x *ShowResponse) GetBase() *BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

var File_vegas_proto protoreflect.FileDescriptor

var file_vegas_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x65, 0x67, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c,
	0x0a, 0x0b, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x65, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x68, 0x22, 0x59, 0x0a, 0x0c,
	0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x06, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x3c, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x65, 0x68, 0x22, 0x35, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x32, 0x68, 0x0a, 0x0c,
	0x56, 0x65, 0x67, 0x61, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x44, 0x72, 0x61, 0x77, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x72, 0x61,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x53, 0x68, 0x6f,
	0x77, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6c, 0x6c, 0x6c, 0x67, 0x70, 0x2f, 0x76, 0x65, 0x67, 0x61,
	0x73, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vegas_proto_rawDescOnce sync.Once
	file_vegas_proto_rawDescData = file_vegas_proto_rawDesc
)

func file_vegas_proto_rawDescGZIP() []byte {
	file_vegas_proto_rawDescOnce.Do(func() {
		file_vegas_proto_rawDescData = protoimpl.X.CompressGZIP(file_vegas_proto_rawDescData)
	})
	return file_vegas_proto_rawDescData
}

var file_vegas_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_vegas_proto_goTypes = []interface{}{
	(*DrawRequest)(nil),  // 0: api.DrawRequest
	(*DrawResponse)(nil), // 1: api.DrawResponse
	(*ShowRequest)(nil),  // 2: api.ShowRequest
	(*ShowResponse)(nil), // 3: api.ShowResponse
	(*User)(nil),         // 4: api.User
	(*BaseResponse)(nil), // 5: api.BaseResponse
	(*Right)(nil),        // 6: api.Right
}
var file_vegas_proto_depIdxs = []int32{
	4, // 0: api.DrawRequest.user:type_name -> api.User
	5, // 1: api.DrawResponse.base:type_name -> api.BaseResponse
	6, // 2: api.DrawResponse.rights:type_name -> api.Right
	4, // 3: api.ShowRequest.user:type_name -> api.User
	5, // 4: api.ShowResponse.base:type_name -> api.BaseResponse
	0, // 5: api.VegasService.Draw:input_type -> api.DrawRequest
	2, // 6: api.VegasService.Show:input_type -> api.ShowRequest
	1, // 7: api.VegasService.Draw:output_type -> api.DrawResponse
	3, // 8: api.VegasService.Show:output_type -> api.ShowResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_vegas_proto_init() }
func file_vegas_proto_init() {
	if File_vegas_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vegas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawRequest); i {
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
		file_vegas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawResponse); i {
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
		file_vegas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowRequest); i {
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
		file_vegas_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowResponse); i {
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
			RawDescriptor: file_vegas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vegas_proto_goTypes,
		DependencyIndexes: file_vegas_proto_depIdxs,
		MessageInfos:      file_vegas_proto_msgTypes,
	}.Build()
	File_vegas_proto = out.File
	file_vegas_proto_rawDesc = nil
	file_vegas_proto_goTypes = nil
	file_vegas_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.10.1. DO NOT EDIT.

type VegasService interface {
	Draw(ctx context.Context, req *DrawRequest) (res *DrawResponse, err error)
	Show(ctx context.Context, req *ShowRequest) (res *ShowResponse, err error)
}
