// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: tradeOrder.proto

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
	TradeOrder_AddTradeOrder_FullMethodName        = "/pb.tradeOrder/AddTradeOrder"
	TradeOrder_UpdateTradeOrder_FullMethodName     = "/pb.tradeOrder/UpdateTradeOrder"
	TradeOrder_GetOrderTotal_FullMethodName        = "/pb.tradeOrder/GetOrderTotal"
	TradeOrder_FindOrder_FullMethodName            = "/pb.tradeOrder/FindOrder"
	TradeOrder_GetOrders_FullMethodName            = "/pb.tradeOrder/GetOrders"
	TradeOrder_GetTradeOrdersByPage_FullMethodName = "/pb.tradeOrder/GetTradeOrdersByPage"
)

// TradeOrderClient is the client API for TradeOrder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradeOrderClient interface {
	AddTradeOrder(ctx context.Context, in *AddTradeOrderReq, opts ...grpc.CallOption) (*AddTradeOrderResp, error)
	UpdateTradeOrder(ctx context.Context, in *AddTradeOrderReq, opts ...grpc.CallOption) (*AddTradeOrderResp, error)
	GetOrderTotal(ctx context.Context, in *OrderTotalReq, opts ...grpc.CallOption) (*OrderTotalResp, error)
	FindOrder(ctx context.Context, in *FindOrderReq, opts ...grpc.CallOption) (*FindOrderResp, error)
	GetOrders(ctx context.Context, in *GetTradeOrderListReq, opts ...grpc.CallOption) (*GetTradeOrderListResp, error)
	GetTradeOrdersByPage(ctx context.Context, in *PageTradeOrderReq, opts ...grpc.CallOption) (*PageTradeOrderResp, error)
}

type tradeOrderClient struct {
	cc grpc.ClientConnInterface
}

func NewTradeOrderClient(cc grpc.ClientConnInterface) TradeOrderClient {
	return &tradeOrderClient{cc}
}

func (c *tradeOrderClient) AddTradeOrder(ctx context.Context, in *AddTradeOrderReq, opts ...grpc.CallOption) (*AddTradeOrderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddTradeOrderResp)
	err := c.cc.Invoke(ctx, TradeOrder_AddTradeOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeOrderClient) UpdateTradeOrder(ctx context.Context, in *AddTradeOrderReq, opts ...grpc.CallOption) (*AddTradeOrderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddTradeOrderResp)
	err := c.cc.Invoke(ctx, TradeOrder_UpdateTradeOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeOrderClient) GetOrderTotal(ctx context.Context, in *OrderTotalReq, opts ...grpc.CallOption) (*OrderTotalResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderTotalResp)
	err := c.cc.Invoke(ctx, TradeOrder_GetOrderTotal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeOrderClient) FindOrder(ctx context.Context, in *FindOrderReq, opts ...grpc.CallOption) (*FindOrderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindOrderResp)
	err := c.cc.Invoke(ctx, TradeOrder_FindOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeOrderClient) GetOrders(ctx context.Context, in *GetTradeOrderListReq, opts ...grpc.CallOption) (*GetTradeOrderListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTradeOrderListResp)
	err := c.cc.Invoke(ctx, TradeOrder_GetOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeOrderClient) GetTradeOrdersByPage(ctx context.Context, in *PageTradeOrderReq, opts ...grpc.CallOption) (*PageTradeOrderResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PageTradeOrderResp)
	err := c.cc.Invoke(ctx, TradeOrder_GetTradeOrdersByPage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradeOrderServer is the server API for TradeOrder service.
// All implementations must embed UnimplementedTradeOrderServer
// for forward compatibility.
type TradeOrderServer interface {
	AddTradeOrder(context.Context, *AddTradeOrderReq) (*AddTradeOrderResp, error)
	UpdateTradeOrder(context.Context, *AddTradeOrderReq) (*AddTradeOrderResp, error)
	GetOrderTotal(context.Context, *OrderTotalReq) (*OrderTotalResp, error)
	FindOrder(context.Context, *FindOrderReq) (*FindOrderResp, error)
	GetOrders(context.Context, *GetTradeOrderListReq) (*GetTradeOrderListResp, error)
	GetTradeOrdersByPage(context.Context, *PageTradeOrderReq) (*PageTradeOrderResp, error)
	mustEmbedUnimplementedTradeOrderServer()
}

// UnimplementedTradeOrderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTradeOrderServer struct{}

func (UnimplementedTradeOrderServer) AddTradeOrder(context.Context, *AddTradeOrderReq) (*AddTradeOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTradeOrder not implemented")
}
func (UnimplementedTradeOrderServer) UpdateTradeOrder(context.Context, *AddTradeOrderReq) (*AddTradeOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTradeOrder not implemented")
}
func (UnimplementedTradeOrderServer) GetOrderTotal(context.Context, *OrderTotalReq) (*OrderTotalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderTotal not implemented")
}
func (UnimplementedTradeOrderServer) FindOrder(context.Context, *FindOrderReq) (*FindOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrder not implemented")
}
func (UnimplementedTradeOrderServer) GetOrders(context.Context, *GetTradeOrderListReq) (*GetTradeOrderListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedTradeOrderServer) GetTradeOrdersByPage(context.Context, *PageTradeOrderReq) (*PageTradeOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTradeOrdersByPage not implemented")
}
func (UnimplementedTradeOrderServer) mustEmbedUnimplementedTradeOrderServer() {}
func (UnimplementedTradeOrderServer) testEmbeddedByValue()                    {}

// UnsafeTradeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradeOrderServer will
// result in compilation errors.
type UnsafeTradeOrderServer interface {
	mustEmbedUnimplementedTradeOrderServer()
}

func RegisterTradeOrderServer(s grpc.ServiceRegistrar, srv TradeOrderServer) {
	// If the following call pancis, it indicates UnimplementedTradeOrderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TradeOrder_ServiceDesc, srv)
}

func _TradeOrder_AddTradeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTradeOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeOrderServer).AddTradeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeOrder_AddTradeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeOrderServer).AddTradeOrder(ctx, req.(*AddTradeOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeOrder_UpdateTradeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTradeOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeOrderServer).UpdateTradeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeOrder_UpdateTradeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeOrderServer).UpdateTradeOrder(ctx, req.(*AddTradeOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeOrder_GetOrderTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderTotalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeOrderServer).GetOrderTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeOrder_GetOrderTotal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeOrderServer).GetOrderTotal(ctx, req.(*OrderTotalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeOrder_FindOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeOrderServer).FindOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeOrder_FindOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeOrderServer).FindOrder(ctx, req.(*FindOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeOrder_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradeOrderListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeOrderServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeOrder_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeOrderServer).GetOrders(ctx, req.(*GetTradeOrderListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeOrder_GetTradeOrdersByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageTradeOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeOrderServer).GetTradeOrdersByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeOrder_GetTradeOrdersByPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeOrderServer).GetTradeOrdersByPage(ctx, req.(*PageTradeOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TradeOrder_ServiceDesc is the grpc.ServiceDesc for TradeOrder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TradeOrder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.tradeOrder",
	HandlerType: (*TradeOrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTradeOrder",
			Handler:    _TradeOrder_AddTradeOrder_Handler,
		},
		{
			MethodName: "UpdateTradeOrder",
			Handler:    _TradeOrder_UpdateTradeOrder_Handler,
		},
		{
			MethodName: "GetOrderTotal",
			Handler:    _TradeOrder_GetOrderTotal_Handler,
		},
		{
			MethodName: "FindOrder",
			Handler:    _TradeOrder_FindOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _TradeOrder_GetOrders_Handler,
		},
		{
			MethodName: "GetTradeOrdersByPage",
			Handler:    _TradeOrder_GetTradeOrdersByPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tradeOrder.proto",
}
