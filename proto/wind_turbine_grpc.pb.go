// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: wind_turbine.proto

package proto

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

// WindTurbineServiceClient is the client API for WindTurbineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WindTurbineServiceClient interface {
	// 接口 1：写入风机数据
	SendData(ctx context.Context, in *WindTurbineData, opts ...grpc.CallOption) (*WriteResponse, error)
	// 接口 2：查询指定风机的平均值
	GetWindTurbineAverage(ctx context.Context, in *WindTurbineAverageRequest, opts ...grpc.CallOption) (*WindTurbineAverageResponse, error)
	// 接口 3：查询全场风机的平均值
	GetAllWindTurbinesAverage(ctx context.Context, in *AllWindTurbinesAverageRequest, opts ...grpc.CallOption) (*WindTurbinesAverageResponse, error)
}

type windTurbineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWindTurbineServiceClient(cc grpc.ClientConnInterface) WindTurbineServiceClient {
	return &windTurbineServiceClient{cc}
}

func (c *windTurbineServiceClient) SendData(ctx context.Context, in *WindTurbineData, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/proto.WindTurbineService/SendData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *windTurbineServiceClient) GetWindTurbineAverage(ctx context.Context, in *WindTurbineAverageRequest, opts ...grpc.CallOption) (*WindTurbineAverageResponse, error) {
	out := new(WindTurbineAverageResponse)
	err := c.cc.Invoke(ctx, "/proto.WindTurbineService/GetWindTurbineAverage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *windTurbineServiceClient) GetAllWindTurbinesAverage(ctx context.Context, in *AllWindTurbinesAverageRequest, opts ...grpc.CallOption) (*WindTurbinesAverageResponse, error) {
	out := new(WindTurbinesAverageResponse)
	err := c.cc.Invoke(ctx, "/proto.WindTurbineService/GetAllWindTurbinesAverage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WindTurbineServiceServer is the server API for WindTurbineService service.
// All implementations must embed UnimplementedWindTurbineServiceServer
// for forward compatibility
type WindTurbineServiceServer interface {
	// 接口 1：写入风机数据
	SendData(context.Context, *WindTurbineData) (*WriteResponse, error)
	// 接口 2：查询指定风机的平均值
	GetWindTurbineAverage(context.Context, *WindTurbineAverageRequest) (*WindTurbineAverageResponse, error)
	// 接口 3：查询全场风机的平均值
	GetAllWindTurbinesAverage(context.Context, *AllWindTurbinesAverageRequest) (*WindTurbinesAverageResponse, error)
	mustEmbedUnimplementedWindTurbineServiceServer()
}

// UnimplementedWindTurbineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWindTurbineServiceServer struct {
}

func (UnimplementedWindTurbineServiceServer) SendData(context.Context, *WindTurbineData) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendData not implemented")
}
func (UnimplementedWindTurbineServiceServer) GetWindTurbineAverage(context.Context, *WindTurbineAverageRequest) (*WindTurbineAverageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWindTurbineAverage not implemented")
}
func (UnimplementedWindTurbineServiceServer) GetAllWindTurbinesAverage(context.Context, *AllWindTurbinesAverageRequest) (*WindTurbinesAverageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllWindTurbinesAverage not implemented")
}
func (UnimplementedWindTurbineServiceServer) mustEmbedUnimplementedWindTurbineServiceServer() {}

// UnsafeWindTurbineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WindTurbineServiceServer will
// result in compilation errors.
type UnsafeWindTurbineServiceServer interface {
	mustEmbedUnimplementedWindTurbineServiceServer()
}

func RegisterWindTurbineServiceServer(s grpc.ServiceRegistrar, srv WindTurbineServiceServer) {
	s.RegisterService(&WindTurbineService_ServiceDesc, srv)
}

func _WindTurbineService_SendData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WindTurbineData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WindTurbineServiceServer).SendData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WindTurbineService/SendData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WindTurbineServiceServer).SendData(ctx, req.(*WindTurbineData))
	}
	return interceptor(ctx, in, info, handler)
}

func _WindTurbineService_GetWindTurbineAverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WindTurbineAverageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WindTurbineServiceServer).GetWindTurbineAverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WindTurbineService/GetWindTurbineAverage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WindTurbineServiceServer).GetWindTurbineAverage(ctx, req.(*WindTurbineAverageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WindTurbineService_GetAllWindTurbinesAverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllWindTurbinesAverageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WindTurbineServiceServer).GetAllWindTurbinesAverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WindTurbineService/GetAllWindTurbinesAverage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WindTurbineServiceServer).GetAllWindTurbinesAverage(ctx, req.(*AllWindTurbinesAverageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WindTurbineService_ServiceDesc is the grpc.ServiceDesc for WindTurbineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WindTurbineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WindTurbineService",
	HandlerType: (*WindTurbineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendData",
			Handler:    _WindTurbineService_SendData_Handler,
		},
		{
			MethodName: "GetWindTurbineAverage",
			Handler:    _WindTurbineService_GetWindTurbineAverage_Handler,
		},
		{
			MethodName: "GetAllWindTurbinesAverage",
			Handler:    _WindTurbineService_GetAllWindTurbinesAverage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wind_turbine.proto",
}