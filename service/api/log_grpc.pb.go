// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: resources/log.proto

package api

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

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	GetAllLogs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (LogService_GetAllLogsClient, error)
	GetLogsFromApplicant(ctx context.Context, in *Request, opts ...grpc.CallOption) (LogService_GetLogsFromApplicantClient, error)
	WriteLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Empty, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) GetAllLogs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (LogService_GetAllLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogService_ServiceDesc.Streams[0], "/loadscheduler.log.service.LogService/GetAllLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceGetAllLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_GetAllLogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type logServiceGetAllLogsClient struct {
	grpc.ClientStream
}

func (x *logServiceGetAllLogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logServiceClient) GetLogsFromApplicant(ctx context.Context, in *Request, opts ...grpc.CallOption) (LogService_GetLogsFromApplicantClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogService_ServiceDesc.Streams[1], "/loadscheduler.log.service.LogService/GetLogsFromApplicant", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceGetLogsFromApplicantClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_GetLogsFromApplicantClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type logServiceGetLogsFromApplicantClient struct {
	grpc.ClientStream
}

func (x *logServiceGetLogsFromApplicantClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logServiceClient) WriteLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/loadscheduler.log.service.LogService/WriteLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility
type LogServiceServer interface {
	GetAllLogs(*Empty, LogService_GetAllLogsServer) error
	GetLogsFromApplicant(*Request, LogService_GetLogsFromApplicantServer) error
	WriteLog(context.Context, *Log) (*Empty, error)
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (UnimplementedLogServiceServer) GetAllLogs(*Empty, LogService_GetAllLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllLogs not implemented")
}
func (UnimplementedLogServiceServer) GetLogsFromApplicant(*Request, LogService_GetLogsFromApplicantServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLogsFromApplicant not implemented")
}
func (UnimplementedLogServiceServer) WriteLog(context.Context, *Log) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteLog not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_GetAllLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).GetAllLogs(m, &logServiceGetAllLogsServer{stream})
}

type LogService_GetAllLogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type logServiceGetAllLogsServer struct {
	grpc.ServerStream
}

func (x *logServiceGetAllLogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

func _LogService_GetLogsFromApplicant_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).GetLogsFromApplicant(m, &logServiceGetLogsFromApplicantServer{stream})
}

type LogService_GetLogsFromApplicantServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type logServiceGetLogsFromApplicantServer struct {
	grpc.ServerStream
}

func (x *logServiceGetLogsFromApplicantServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

func _LogService_WriteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).WriteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loadscheduler.log.service.LogService/WriteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).WriteLog(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loadscheduler.log.service.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteLog",
			Handler:    _LogService_WriteLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllLogs",
			Handler:       _LogService_GetAllLogs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetLogsFromApplicant",
			Handler:       _LogService_GetLogsFromApplicant_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "resources/log.proto",
}