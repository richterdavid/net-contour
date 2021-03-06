// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/eds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdsDummy) Reset()         { *m = EdsDummy{} }
func (m *EdsDummy) String() string { return proto.CompactTextString(m) }
func (*EdsDummy) ProtoMessage()    {}
func (*EdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c8760f38742c17f, []int{0}
}

func (m *EdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdsDummy.Unmarshal(m, b)
}
func (m *EdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdsDummy.Marshal(b, m, deterministic)
}
func (m *EdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdsDummy.Merge(m, src)
}
func (m *EdsDummy) XXX_Size() int {
	return xxx_messageInfo_EdsDummy.Size(m)
}
func (m *EdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_EdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_EdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EdsDummy)(nil), "envoy.api.v2.EdsDummy")
}

func init() { proto.RegisterFile("envoy/api/v2/eds.proto", fileDescriptor_3c8760f38742c17f) }

var fileDescriptor_3c8760f38742c17f = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbd, 0x4e, 0xe3, 0x40,
	0x14, 0x85, 0xd7, 0x1b, 0xed, 0x66, 0x35, 0x5a, 0x65, 0x25, 0x17, 0x9b, 0x5d, 0x6f, 0x94, 0x8d,
	0x4c, 0x8a, 0x28, 0x85, 0x8d, 0x92, 0x2e, 0x1d, 0x21, 0xa1, 0xa2, 0x08, 0xa4, 0xa1, 0x9d, 0x78,
	0x2e, 0xce, 0x48, 0xf6, 0xcc, 0x30, 0x33, 0x36, 0x58, 0x34, 0x88, 0x0a, 0xd1, 0x50, 0x20, 0x21,
	0x1e, 0x80, 0x37, 0xe2, 0x15, 0x78, 0x02, 0xe8, 0x11, 0x8a, 0x7f, 0x12, 0x0c, 0x0a, 0x15, 0x9d,
	0xed, 0xef, 0xdc, 0x73, 0x7f, 0x8e, 0xd1, 0x6f, 0x60, 0x31, 0x4f, 0x5c, 0x2c, 0xa8, 0x1b, 0xf7,
	0x5c, 0x20, 0xca, 0x11, 0x92, 0x6b, 0x6e, 0xfe, 0x4c, 0xbf, 0x3b, 0x58, 0x50, 0x27, 0xee, 0x59,
	0x8d, 0x92, 0x8a, 0x50, 0xe5, 0xf1, 0x18, 0x64, 0x92, 0x69, 0xad, 0x86, 0xcf, 0xb9, 0x1f, 0x40,
	0x8a, 0x31, 0x63, 0x5c, 0x63, 0x4d, 0x39, 0xcb, 0x9d, 0xac, 0x66, 0x4e, 0xd3, 0xb7, 0x59, 0x74,
	0xe8, 0x92, 0x48, 0xa6, 0x82, 0x75, 0xfc, 0x58, 0x62, 0x21, 0x40, 0x16, 0xf5, 0xad, 0xbc, 0xf7,
	0xca, 0xd8, 0x95, 0xa0, 0x78, 0x24, 0x3d, 0x28, 0x1c, 0x22, 0x22, 0x70, 0x49, 0x10, 0x52, 0x5f,
	0x62, 0x5d, 0xf0, 0x7a, 0x8c, 0x03, 0x4a, 0xb0, 0x06, 0xb7, 0x78, 0xc8, 0xc1, 0xbf, 0xf2, 0xf2,
	0x8c, 0x08, 0x4e, 0x99, 0xce, 0xa0, 0x8d, 0xd0, 0x8f, 0x31, 0x51, 0xa3, 0x28, 0x0c, 0x93, 0xde,
	0x59, 0x05, 0xfd, 0x19, 0xe7, 0x78, 0x54, 0x6c, 0x3f, 0x05, 0x19, 0x53, 0x0f, 0xcc, 0x03, 0xf4,
	0x6b, 0xaa, 0x25, 0xe0, 0xb0, 0x50, 0x28, 0xb3, 0xe9, 0xbc, 0x3e, 0x9f, 0xb3, 0x2c, 0xd9, 0x87,
	0xa3, 0x08, 0x94, 0xb6, 0xfe, 0xaf, 0xe5, 0x4a, 0x70, 0xa6, 0xc0, 0xfe, 0xd2, 0x31, 0x36, 0x0d,
	0x13, 0xa3, 0xda, 0x08, 0x02, 0x8d, 0x57, 0xc6, 0x1b, 0x6f, 0x0a, 0x17, 0xf4, 0x9d, 0x7b, 0xfb,
	0x63, 0x51, 0xa9, 0xc5, 0x29, 0xaa, 0xed, 0x80, 0xf6, 0xe6, 0x9f, 0x38, 0x7b, 0xe7, 0xfc, 0xfe,
	0xe1, 0xfa, 0xeb, 0x5f, 0xbb, 0x5e, 0xfa, 0x57, 0x06, 0xc5, 0x79, 0x55, 0x8a, 0x2b, 0x03, 0xa3,
	0x6b, 0x75, 0x2f, 0xef, 0x6e, 0x9e, 0xaa, 0x6d, 0x64, 0x97, 0x1c, 0xb7, 0x83, 0x48, 0x69, 0x90,
	0xbb, 0x1c, 0x93, 0x2d, 0xa5, 0xa8, 0xcf, 0x42, 0x60, 0x7a, 0xb8, 0xf7, 0x78, 0xfb, 0x7c, 0xf5,
	0xad, 0x65, 0x36, 0x33, 0xad, 0xca, 0x8e, 0xef, 0x2c, 0x43, 0x8b, 0xfb, 0x38, 0x10, 0x73, 0x8c,
	0x2c, 0xca, 0xb3, 0x01, 0x85, 0xe4, 0x27, 0x49, 0x69, 0xd6, 0xe1, 0x22, 0xd0, 0xc9, 0x22, 0xdc,
	0x89, 0x71, 0x61, 0x18, 0x93, 0xea, 0xec, 0x7b, 0x1a, 0x75, 0xff, 0x25, 0x00, 0x00, 0xff, 0xff,
	0xf2, 0x6a, 0x05, 0x16, 0x06, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointDiscoveryServiceClient is the client API for EndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointDiscoveryServiceClient interface {
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.EndpointDiscoveryService/DeltaEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceDeltaEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_DeltaEndpointsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceDeltaEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoveryServiceServer is the server API for EndpointDiscoveryService service.
type EndpointDiscoveryServiceServer interface {
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	DeltaEndpoints(EndpointDiscoveryService_DeltaEndpointsServer) error
	FetchEndpoints(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedEndpointDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEndpointDiscoveryServiceServer struct {
}

func (*UnimplementedEndpointDiscoveryServiceServer) StreamEndpoints(srv EndpointDiscoveryService_StreamEndpointsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEndpoints not implemented")
}
func (*UnimplementedEndpointDiscoveryServiceServer) DeltaEndpoints(srv EndpointDiscoveryService_DeltaEndpointsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaEndpoints not implemented")
}
func (*UnimplementedEndpointDiscoveryServiceServer) FetchEndpoints(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchEndpoints not implemented")
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_DeltaEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).DeltaEndpoints(&endpointDiscoveryServiceDeltaEndpointsServer{stream})
}

type EndpointDiscoveryService_DeltaEndpointsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceDeltaEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaEndpoints",
			Handler:       _EndpointDiscoveryService_DeltaEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/eds.proto",
}
