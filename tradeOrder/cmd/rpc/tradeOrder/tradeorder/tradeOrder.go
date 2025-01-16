// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: tradeOrder.proto

package tradeorder

import (
	"context"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddProductOrderReq             = pb.AddProductOrderReq
	AddProductOrderResp            = pb.AddProductOrderResp
	AddTradeOrderReq               = pb.AddTradeOrderReq
	AddTradeOrderResp              = pb.AddTradeOrderResp
	FindOrderReq                   = pb.FindOrderReq
	FindOrderResp                  = pb.FindOrderResp
	GetProductOrderByOrderIdReq    = pb.GetProductOrderByOrderIdReq
	GetProductOrderByOrderIdResp   = pb.GetProductOrderByOrderIdResp
	GetProductOrderByUserIdReq     = pb.GetProductOrderByUserIdReq
	GetProductOrderByUserIdResp    = pb.GetProductOrderByUserIdResp
	GetProductOrderListReq         = pb.GetProductOrderListReq
	GetProductOrderListResp        = pb.GetProductOrderListResp
	GetTradeOrderListReq           = pb.GetTradeOrderListReq
	GetTradeOrderListResp          = pb.GetTradeOrderListResp
	OrderTotalReq                  = pb.OrderTotalReq
	OrderTotalResp                 = pb.OrderTotalResp
	PageTradeOrderReq              = pb.PageTradeOrderReq
	PageTradeOrderResp             = pb.PageTradeOrderResp
	ProductOrder                   = pb.ProductOrder
	ShoppingCartSimple             = pb.ShoppingCartSimple
	ShoppingCartsProductInfoSimple = pb.ShoppingCartsProductInfoSimple
	TradeOrder                     = pb.TradeOrder
	UpdateTradeOrderReq            = pb.UpdateTradeOrderReq
	UpdateTradeOrderResp           = pb.UpdateTradeOrderResp

	TradeOrderZrpcClient interface {
		AddTradeOrder(ctx context.Context, in *AddTradeOrderReq, opts ...grpc.CallOption) (*AddTradeOrderResp, error)
		UpdateTradeOrder(ctx context.Context, in *UpdateTradeOrderReq, opts ...grpc.CallOption) (*UpdateTradeOrderResp, error)
		GetOrderTotal(ctx context.Context, in *OrderTotalReq, opts ...grpc.CallOption) (*OrderTotalResp, error)
		FindOrder(ctx context.Context, in *FindOrderReq, opts ...grpc.CallOption) (*FindOrderResp, error)
		GetOrders(ctx context.Context, in *GetTradeOrderListReq, opts ...grpc.CallOption) (*GetTradeOrderListResp, error)
		GetTradeOrdersByPage(ctx context.Context, in *PageTradeOrderReq, opts ...grpc.CallOption) (*PageTradeOrderResp, error)
		BatchCreateOrderProduct(ctx context.Context, in *AddProductOrderReq, opts ...grpc.CallOption) (*AddProductOrderResp, error)
		GetProductOrderList(ctx context.Context, in *GetProductOrderListReq, opts ...grpc.CallOption) (*GetProductOrderListResp, error)
		GetProductOrderByUserId(ctx context.Context, in *GetProductOrderByUserIdReq, opts ...grpc.CallOption) (*GetProductOrderByUserIdResp, error)
	}

	defaultTradeOrderZrpcClient struct {
		cli zrpc.Client
	}
)

func NewTradeOrderZrpcClient(cli zrpc.Client) TradeOrderZrpcClient {
	return &defaultTradeOrderZrpcClient{
		cli: cli,
	}
}

func (m *defaultTradeOrderZrpcClient) AddTradeOrder(ctx context.Context, in *AddTradeOrderReq, opts ...grpc.CallOption) (*AddTradeOrderResp, error) {
	client := pb.NewTradeOrderClient(m.cli.Conn())
	return client.AddTradeOrder(ctx, in, opts...)
}

func (m *defaultTradeOrderZrpcClient) UpdateTradeOrder(ctx context.Context, in *UpdateTradeOrderReq, opts ...grpc.CallOption) (*UpdateTradeOrderResp, error) {
	client := pb.NewTradeOrderClient(m.cli.Conn())
	return client.UpdateTradeOrder(ctx, in, opts...)
}

func (m *defaultTradeOrderZrpcClient) GetOrderTotal(ctx context.Context, in *OrderTotalReq, opts ...grpc.CallOption) (*OrderTotalResp, error) {
	client := pb.NewTradeOrderClient(m.cli.Conn())
	return client.GetOrderTotal(ctx, in, opts...)
}

func (m *defaultTradeOrderZrpcClient) FindOrder(ctx context.Context, in *FindOrderReq, opts ...grpc.CallOption) (*FindOrderResp, error) {
	client := pb.NewTradeOrderClient(m.cli.Conn())
	return client.FindOrder(ctx, in, opts...)
}

func (m *defaultTradeOrderZrpcClient) GetOrders(ctx context.Context, in *GetTradeOrderListReq, opts ...grpc.CallOption) (*GetTradeOrderListResp, error) {
	client := pb.NewTradeOrderClient(m.cli.Conn())
	return client.GetOrders(ctx, in, opts...)
}

func (m *defaultTradeOrderZrpcClient) GetTradeOrdersByPage(ctx context.Context, in *PageTradeOrderReq, opts ...grpc.CallOption) (*PageTradeOrderResp, error) {
	client := pb.NewTradeOrderClient(m.cli.Conn())
	return client.GetTradeOrdersByPage(ctx, in, opts...)
}

func (m *defaultTradeOrderZrpcClient) BatchCreateOrderProduct(ctx context.Context, in *AddProductOrderReq, opts ...grpc.CallOption) (*AddProductOrderResp, error) {
	client := pb.NewTradeOrderClient(m.cli.Conn())
	return client.BatchCreateOrderProduct(ctx, in, opts...)
}

func (m *defaultTradeOrderZrpcClient) GetProductOrderList(ctx context.Context, in *GetProductOrderListReq, opts ...grpc.CallOption) (*GetProductOrderListResp, error) {
	client := pb.NewTradeOrderClient(m.cli.Conn())
	return client.GetProductOrderList(ctx, in, opts...)
}

func (m *defaultTradeOrderZrpcClient) GetProductOrderByUserId(ctx context.Context, in *GetProductOrderByUserIdReq, opts ...grpc.CallOption) (*GetProductOrderByUserIdResp, error) {
	client := pb.NewTradeOrderClient(m.cli.Conn())
	return client.GetProductOrderByUserId(ctx, in, opts...)
}
