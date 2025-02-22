// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: fetcher.proto

package fetcher

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FetcherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
}

func (x *FetcherRequest) Reset() {
	*x = FetcherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetcherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetcherRequest) ProtoMessage() {}

func (x *FetcherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fetcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetcherRequest.ProtoReflect.Descriptor instead.
func (*FetcherRequest) Descriptor() ([]byte, []int) {
	return file_fetcher_proto_rawDescGZIP(), []int{0}
}

func (x *FetcherRequest) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

type FetcherReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *FetcherReply) Reset() {
	*x = FetcherReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetcherReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetcherReply) ProtoMessage() {}

func (x *FetcherReply) ProtoReflect() protoreflect.Message {
	mi := &file_fetcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetcherReply.ProtoReflect.Descriptor instead.
func (*FetcherReply) Descriptor() ([]byte, []int) {
	return file_fetcher_proto_rawDescGZIP(), []int{1}
}

func (x *FetcherReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_fetcher_proto protoreflect.FileDescriptor

var file_fetcher_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x22, 0x24, 0x0a, 0x0c,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x32, 0x42, 0x0a, 0x07, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x37, 0x0a,
	0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x66, 0x75, 0x77, 0x61, 0x66, 0x75, 0x2f, 0x48, 0x61,
	0x74, 0x65, 0x6e, 0x61, 0x2d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x2d, 0x32, 0x30, 0x32, 0x30,
	0x2f, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fetcher_proto_rawDescOnce sync.Once
	file_fetcher_proto_rawDescData = file_fetcher_proto_rawDesc
)

func file_fetcher_proto_rawDescGZIP() []byte {
	file_fetcher_proto_rawDescOnce.Do(func() {
		file_fetcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_fetcher_proto_rawDescData)
	})
	return file_fetcher_proto_rawDescData
}

var file_fetcher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fetcher_proto_goTypes = []interface{}{
	(*FetcherRequest)(nil), // 0: fetcher.FetcherRequest
	(*FetcherReply)(nil),   // 1: fetcher.FetcherReply
}
var file_fetcher_proto_depIdxs = []int32{
	0, // 0: fetcher.Fetcher.Fetch:input_type -> fetcher.FetcherRequest
	1, // 1: fetcher.Fetcher.Fetch:output_type -> fetcher.FetcherReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fetcher_proto_init() }
func file_fetcher_proto_init() {
	if File_fetcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fetcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetcherRequest); i {
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
		file_fetcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetcherReply); i {
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
			RawDescriptor: file_fetcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fetcher_proto_goTypes,
		DependencyIndexes: file_fetcher_proto_depIdxs,
		MessageInfos:      file_fetcher_proto_msgTypes,
	}.Build()
	File_fetcher_proto = out.File
	file_fetcher_proto_rawDesc = nil
	file_fetcher_proto_goTypes = nil
	file_fetcher_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FetcherClient is the client API for Fetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FetcherClient interface {
	Fetch(ctx context.Context, in *FetcherRequest, opts ...grpc.CallOption) (*FetcherReply, error)
}

type fetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewFetcherClient(cc grpc.ClientConnInterface) FetcherClient {
	return &fetcherClient{cc}
}

func (c *fetcherClient) Fetch(ctx context.Context, in *FetcherRequest, opts ...grpc.CallOption) (*FetcherReply, error) {
	out := new(FetcherReply)
	err := c.cc.Invoke(ctx, "/fetcher.Fetcher/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetcherServer is the server API for Fetcher service.
type FetcherServer interface {
	Fetch(context.Context, *FetcherRequest) (*FetcherReply, error)
}

// UnimplementedFetcherServer can be embedded to have forward compatible implementations.
type UnimplementedFetcherServer struct {
}

func (*UnimplementedFetcherServer) Fetch(context.Context, *FetcherRequest) (*FetcherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}

func RegisterFetcherServer(s *grpc.Server, srv FetcherServer) {
	s.RegisterService(&_Fetcher_serviceDesc, srv)
}

func _Fetcher_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetcherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetcherServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetcher.Fetcher/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetcherServer).Fetch(ctx, req.(*FetcherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fetcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fetcher.Fetcher",
	HandlerType: (*FetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Fetcher_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fetcher.proto",
}
