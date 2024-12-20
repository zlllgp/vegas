// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: right.proto

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

type RightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *User              `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Eh      string             `protobuf:"bytes,2,opt,name=eh,proto3" json:"eh,omitempty"`
	PageReq *PaginationRequest `protobuf:"bytes,3,opt,name=page_req,json=pageReq,proto3" json:"page_req,omitempty"`
	Sfs     []*SearchField     `protobuf:"bytes,4,rep,name=sfs,proto3" json:"sfs,omitempty"`
}

func (x *RightRequest) Reset() {
	*x = RightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_right_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RightRequest) ProtoMessage() {}

func (x *RightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_right_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RightRequest.ProtoReflect.Descriptor instead.
func (*RightRequest) Descriptor() ([]byte, []int) {
	return file_right_proto_rawDescGZIP(), []int{0}
}

func (x *RightRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RightRequest) GetEh() string {
	if x != nil {
		return x.Eh
	}
	return ""
}

func (x *RightRequest) GetPageReq() *PaginationRequest {
	if x != nil {
		return x.PageReq
	}
	return nil
}

func (x *RightRequest) GetSfs() []*SearchField {
	if x != nil {
		return x.Sfs
	}
	return nil
}

type RightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base     *BaseResponse       `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	PageResp *PaginationResponse `protobuf:"bytes,2,opt,name=page_resp,json=pageResp,proto3" json:"page_resp,omitempty"`
	Rights   []*Right            `protobuf:"bytes,3,rep,name=rights,proto3" json:"rights,omitempty"`
}

func (x *RightResponse) Reset() {
	*x = RightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_right_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RightResponse) ProtoMessage() {}

func (x *RightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_right_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RightResponse.ProtoReflect.Descriptor instead.
func (*RightResponse) Descriptor() ([]byte, []int) {
	return file_right_proto_rawDescGZIP(), []int{1}
}

func (x *RightResponse) GetBase() *BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *RightResponse) GetPageResp() *PaginationResponse {
	if x != nil {
		return x.PageResp
	}
	return nil
}

func (x *RightResponse) GetRights() []*Right {
	if x != nil {
		return x.Rights
	}
	return nil
}

var File_right_proto protoreflect.FileDescriptor

var file_right_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x01, 0x0a, 0x0c, 0x52, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x65, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x68, 0x12, 0x31,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x22, 0x0a, 0x03, 0x73, 0x66, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x03, 0x73, 0x66, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x0d, 0x52, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x69, 0x67, 0x68, 0x74,
	0x52, 0x06, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x32, 0x43, 0x0a, 0x0c, 0x52, 0x69, 0x67, 0x68,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6c, 0x6c, 0x6c,
	0x67, 0x70, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x73, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_right_proto_rawDescOnce sync.Once
	file_right_proto_rawDescData = file_right_proto_rawDesc
)

func file_right_proto_rawDescGZIP() []byte {
	file_right_proto_rawDescOnce.Do(func() {
		file_right_proto_rawDescData = protoimpl.X.CompressGZIP(file_right_proto_rawDescData)
	})
	return file_right_proto_rawDescData
}

var file_right_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_right_proto_goTypes = []interface{}{
	(*RightRequest)(nil),       // 0: api.RightRequest
	(*RightResponse)(nil),      // 1: api.RightResponse
	(*User)(nil),               // 2: api.User
	(*PaginationRequest)(nil),  // 3: api.PaginationRequest
	(*SearchField)(nil),        // 4: api.SearchField
	(*BaseResponse)(nil),       // 5: api.BaseResponse
	(*PaginationResponse)(nil), // 6: api.PaginationResponse
	(*Right)(nil),              // 7: api.Right
}
var file_right_proto_depIdxs = []int32{
	2, // 0: api.RightRequest.user:type_name -> api.User
	3, // 1: api.RightRequest.page_req:type_name -> api.PaginationRequest
	4, // 2: api.RightRequest.sfs:type_name -> api.SearchField
	5, // 3: api.RightResponse.base:type_name -> api.BaseResponse
	6, // 4: api.RightResponse.page_resp:type_name -> api.PaginationResponse
	7, // 5: api.RightResponse.rights:type_name -> api.Right
	0, // 6: api.RightService.QueryRight:input_type -> api.RightRequest
	1, // 7: api.RightService.QueryRight:output_type -> api.RightResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_right_proto_init() }
func file_right_proto_init() {
	if File_right_proto != nil {
		return
	}
	file_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_right_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RightRequest); i {
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
		file_right_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RightResponse); i {
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
			RawDescriptor: file_right_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_right_proto_goTypes,
		DependencyIndexes: file_right_proto_depIdxs,
		MessageInfos:      file_right_proto_msgTypes,
	}.Build()
	File_right_proto = out.File
	file_right_proto_rawDesc = nil
	file_right_proto_goTypes = nil
	file_right_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.10.1. DO NOT EDIT.

type RightService interface {
	QueryRight(ctx context.Context, req *RightRequest) (res *RightResponse, err error)
}
