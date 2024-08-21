// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0--rc2
// source: proto/math.proto

package pb

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
	MathService_Add_FullMethodName      = "/MathService/Add"
	MathService_Subtract_FullMethodName = "/MathService/Subtract"
	MathService_Multiply_FullMethodName = "/MathService/Multiply"
	MathService_Divide_FullMethodName   = "/MathService/Divide"
)

// MathServiceClient is the client API for MathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MathServiceClient interface {
	Add(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Subtract(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Multiply(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Divide(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
}

type mathServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMathServiceClient(cc grpc.ClientConnInterface) MathServiceClient {
	return &mathServiceClient{cc}
}

func (c *mathServiceClient) Add(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, MathService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Subtract(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, MathService_Subtract_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Multiply(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, MathService_Multiply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Divide(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, MathService_Divide_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServiceServer is the server API for MathService service.
// All implementations must embed UnimplementedMathServiceServer
// for forward compatibility.
type MathServiceServer interface {
	Add(context.Context, *MathRequest) (*MathResponse, error)
	Subtract(context.Context, *MathRequest) (*MathResponse, error)
	Multiply(context.Context, *MathRequest) (*MathResponse, error)
	Divide(context.Context, *MathRequest) (*MathResponse, error)
	mustEmbedUnimplementedMathServiceServer()
}

// UnimplementedMathServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMathServiceServer struct{}

func (UnimplementedMathServiceServer) Add(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedMathServiceServer) Subtract(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedMathServiceServer) Multiply(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedMathServiceServer) Divide(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedMathServiceServer) mustEmbedUnimplementedMathServiceServer() {}
func (UnimplementedMathServiceServer) testEmbeddedByValue()                     {}

// UnsafeMathServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MathServiceServer will
// result in compilation errors.
type UnsafeMathServiceServer interface {
	mustEmbedUnimplementedMathServiceServer()
}

func RegisterMathServiceServer(s grpc.ServiceRegistrar, srv MathServiceServer) {
	// If the following call pancis, it indicates UnimplementedMathServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MathService_ServiceDesc, srv)
}

func _MathService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MathService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Add(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MathService_Subtract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Subtract(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MathService_Multiply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Multiply(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MathService_Divide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Divide(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MathService_ServiceDesc is the grpc.ServiceDesc for MathService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MathService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MathService",
	HandlerType: (*MathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MathService_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _MathService_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _MathService_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _MathService_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/math.proto",
}
