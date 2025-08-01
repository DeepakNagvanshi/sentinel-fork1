// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: proto/proto/log_ingestor.proto

package logging

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LogIngestor_SendLogParser_FullMethodName = "/logging.LogIngestor/SendLogParser"
)

// LogIngestorClient is the client API for LogIngestor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogIngestorClient interface {
	SendLogParser(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
}

type logIngestorClient struct {
	cc grpc.ClientConnInterface
}

func NewLogIngestorClient(cc grpc.ClientConnInterface) LogIngestorClient {
	return &logIngestorClient{cc}
}

func (c *logIngestorClient) SendLogParser(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, LogIngestor_SendLogParser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogIngestorServer is the server API for LogIngestor service.
// All implementations must embed UnimplementedLogIngestorServer
// for forward compatibility.
type LogIngestorServer interface {
	SendLogParser(context.Context, *LogRequest) (*LogResponse, error)
	mustEmbedUnimplementedLogIngestorServer()
}

// UnimplementedLogIngestorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogIngestorServer struct{}

func (UnimplementedLogIngestorServer) SendLogParser(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLogParser not implemented")
}
func (UnimplementedLogIngestorServer) mustEmbedUnimplementedLogIngestorServer() {}
func (UnimplementedLogIngestorServer) testEmbeddedByValue()                     {}

// UnsafeLogIngestorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogIngestorServer will
// result in compilation errors.
type UnsafeLogIngestorServer interface {
	mustEmbedUnimplementedLogIngestorServer()
}

func RegisterLogIngestorServer(s grpc.ServiceRegistrar, srv LogIngestorServer) {
	// If the following call pancis, it indicates UnimplementedLogIngestorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LogIngestor_ServiceDesc, srv)
}

func _LogIngestor_SendLogParser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogIngestorServer).SendLogParser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogIngestor_SendLogParser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogIngestorServer).SendLogParser(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogIngestor_ServiceDesc is the grpc.ServiceDesc for LogIngestor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogIngestor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logging.LogIngestor",
	HandlerType: (*LogIngestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLogParser",
			Handler:    _LogIngestor_SendLogParser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto/log_ingestor.proto",
}
