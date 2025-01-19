// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: marketing.proto

//import "product/cmd/rpc/product/pb/product.proto";

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
	Activity_AddActivity_FullMethodName            = "/pb.activity/AddActivity"
	Activity_AddActivityTime_FullMethodName        = "/pb.activity/AddActivityTime"
	Activity_AddActivityProduct_FullMethodName     = "/pb.activity/AddActivityProduct"
	Activity_AddActivityProductSku_FullMethodName  = "/pb.activity/AddActivityProductSku"
	Activity_GetActivityProductById_FullMethodName = "/pb.activity/GetActivityProductById"
)

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityClient interface {
	AddActivity(ctx context.Context, in *AddActivityReq, opts ...grpc.CallOption) (*AddActivityResp, error)
	AddActivityTime(ctx context.Context, in *AddActivityTimeReq, opts ...grpc.CallOption) (*AddActivityTimeResp, error)
	AddActivityProduct(ctx context.Context, in *AddActivityProductReq, opts ...grpc.CallOption) (*AddActivityProductResp, error)
	AddActivityProductSku(ctx context.Context, in *AddActivityProductSkuReq, opts ...grpc.CallOption) (*AddActivityProductSkuResp, error)
	// 查询接口
	GetActivityProductById(ctx context.Context, in *GetActivityProductByIdReq, opts ...grpc.CallOption) (*GetActivityProductByIdResp, error)
}

type activityClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityClient(cc grpc.ClientConnInterface) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) AddActivity(ctx context.Context, in *AddActivityReq, opts ...grpc.CallOption) (*AddActivityResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddActivityResp)
	err := c.cc.Invoke(ctx, Activity_AddActivity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) AddActivityTime(ctx context.Context, in *AddActivityTimeReq, opts ...grpc.CallOption) (*AddActivityTimeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddActivityTimeResp)
	err := c.cc.Invoke(ctx, Activity_AddActivityTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) AddActivityProduct(ctx context.Context, in *AddActivityProductReq, opts ...grpc.CallOption) (*AddActivityProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddActivityProductResp)
	err := c.cc.Invoke(ctx, Activity_AddActivityProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) AddActivityProductSku(ctx context.Context, in *AddActivityProductSkuReq, opts ...grpc.CallOption) (*AddActivityProductSkuResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddActivityProductSkuResp)
	err := c.cc.Invoke(ctx, Activity_AddActivityProductSku_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) GetActivityProductById(ctx context.Context, in *GetActivityProductByIdReq, opts ...grpc.CallOption) (*GetActivityProductByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActivityProductByIdResp)
	err := c.cc.Invoke(ctx, Activity_GetActivityProductById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
// All implementations must embed UnimplementedActivityServer
// for forward compatibility.
type ActivityServer interface {
	AddActivity(context.Context, *AddActivityReq) (*AddActivityResp, error)
	AddActivityTime(context.Context, *AddActivityTimeReq) (*AddActivityTimeResp, error)
	AddActivityProduct(context.Context, *AddActivityProductReq) (*AddActivityProductResp, error)
	AddActivityProductSku(context.Context, *AddActivityProductSkuReq) (*AddActivityProductSkuResp, error)
	// 查询接口
	GetActivityProductById(context.Context, *GetActivityProductByIdReq) (*GetActivityProductByIdResp, error)
	mustEmbedUnimplementedActivityServer()
}

// UnimplementedActivityServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActivityServer struct{}

func (UnimplementedActivityServer) AddActivity(context.Context, *AddActivityReq) (*AddActivityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddActivity not implemented")
}
func (UnimplementedActivityServer) AddActivityTime(context.Context, *AddActivityTimeReq) (*AddActivityTimeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddActivityTime not implemented")
}
func (UnimplementedActivityServer) AddActivityProduct(context.Context, *AddActivityProductReq) (*AddActivityProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddActivityProduct not implemented")
}
func (UnimplementedActivityServer) AddActivityProductSku(context.Context, *AddActivityProductSkuReq) (*AddActivityProductSkuResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddActivityProductSku not implemented")
}
func (UnimplementedActivityServer) GetActivityProductById(context.Context, *GetActivityProductByIdReq) (*GetActivityProductByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivityProductById not implemented")
}
func (UnimplementedActivityServer) mustEmbedUnimplementedActivityServer() {}
func (UnimplementedActivityServer) testEmbeddedByValue()                  {}

// UnsafeActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServer will
// result in compilation errors.
type UnsafeActivityServer interface {
	mustEmbedUnimplementedActivityServer()
}

func RegisterActivityServer(s grpc.ServiceRegistrar, srv ActivityServer) {
	// If the following call pancis, it indicates UnimplementedActivityServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Activity_ServiceDesc, srv)
}

func _Activity_AddActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).AddActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_AddActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).AddActivity(ctx, req.(*AddActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_AddActivityTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddActivityTimeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).AddActivityTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_AddActivityTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).AddActivityTime(ctx, req.(*AddActivityTimeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_AddActivityProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddActivityProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).AddActivityProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_AddActivityProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).AddActivityProduct(ctx, req.(*AddActivityProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_AddActivityProductSku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddActivityProductSkuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).AddActivityProductSku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_AddActivityProductSku_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).AddActivityProductSku(ctx, req.(*AddActivityProductSkuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_GetActivityProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityProductByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetActivityProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Activity_GetActivityProductById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetActivityProductById(ctx, req.(*GetActivityProductByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_ServiceDesc is the grpc.ServiceDesc for Activity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddActivity",
			Handler:    _Activity_AddActivity_Handler,
		},
		{
			MethodName: "AddActivityTime",
			Handler:    _Activity_AddActivityTime_Handler,
		},
		{
			MethodName: "AddActivityProduct",
			Handler:    _Activity_AddActivityProduct_Handler,
		},
		{
			MethodName: "AddActivityProductSku",
			Handler:    _Activity_AddActivityProductSku_Handler,
		},
		{
			MethodName: "GetActivityProductById",
			Handler:    _Activity_GetActivityProductById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketing.proto",
}
