// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/contract.proto

package BAUMAN_HACK_IU5

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MyService_BidirectionalStreaming_FullMethodName = "/example.MyService/BidirectionalStreaming"
)

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (MyService_BidirectionalStreamingClient, error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (MyService_BidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &MyService_ServiceDesc.Streams[0], MyService_BidirectionalStreaming_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &myServiceBidirectionalStreamingClient{stream}
	return x, nil
}

type MyService_BidirectionalStreamingClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type myServiceBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *myServiceBidirectionalStreamingClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myServiceBidirectionalStreamingClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility
type MyServiceServer interface {
	BidirectionalStreaming(MyService_BidirectionalStreamingServer) error
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (UnimplementedMyServiceServer) BidirectionalStreaming(MyService_BidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreaming not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_BidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyServiceServer).BidirectionalStreaming(&myServiceBidirectionalStreamingServer{stream})
}

type MyService_BidirectionalStreamingServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type myServiceBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *myServiceBidirectionalStreamingServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myServiceBidirectionalStreamingServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BidirectionalStreaming",
			Handler:       _MyService_BidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/contract.proto",
}
