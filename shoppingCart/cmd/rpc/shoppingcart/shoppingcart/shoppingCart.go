// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: cart.proto

package shoppingcart

import (
	"context"

	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCartReq            = pb.AddCartReq
	AddCartResp           = pb.AddCartResp
	FindCartReq           = pb.FindCartReq
	FindCartResp          = pb.FindCartResp
	FindCartsByUserIdReq  = pb.FindCartsByUserIdReq
	FindCartsByUserIdResp = pb.FindCartsByUserIdResp
	FindCartsReq          = pb.FindCartsReq
	FindCartsResp         = pb.FindCartsResp
	ProductSimple         = pb.ProductSimple
	ProductSkuSimple      = pb.ProductSkuSimple
	ShoppingCart          = pb.ShoppingCart
	UpdateCartReq         = pb.UpdateCartReq
	UpdateCartResp        = pb.UpdateCartResp

	ShoppingCartZrpcClient interface {
		AddCart(ctx context.Context, in *AddCartReq, opts ...grpc.CallOption) (*AddCartResp, error)
		UpdateCart(ctx context.Context, in *UpdateCartReq, opts ...grpc.CallOption) (*UpdateCartResp, error)
		FindCart(ctx context.Context, in *FindCartReq, opts ...grpc.CallOption) (*FindCartResp, error)
		GetCarts(ctx context.Context, in *FindCartsReq, opts ...grpc.CallOption) (*FindCartsResp, error)
		GetCartsByUserId(ctx context.Context, in *FindCartsByUserIdReq, opts ...grpc.CallOption) (*FindCartsByUserIdResp, error)
	}

	defaultShoppingCartZrpcClient struct {
		cli zrpc.Client
	}
)

func NewShoppingCartZrpcClient(cli zrpc.Client) ShoppingCartZrpcClient {
	return &defaultShoppingCartZrpcClient{
		cli: cli,
	}
}

func (m *defaultShoppingCartZrpcClient) AddCart(ctx context.Context, in *AddCartReq, opts ...grpc.CallOption) (*AddCartResp, error) {
	client := pb.NewShoppingCartClient(m.cli.Conn())
	return client.AddCart(ctx, in, opts...)
}

func (m *defaultShoppingCartZrpcClient) UpdateCart(ctx context.Context, in *UpdateCartReq, opts ...grpc.CallOption) (*UpdateCartResp, error) {
	client := pb.NewShoppingCartClient(m.cli.Conn())
	return client.UpdateCart(ctx, in, opts...)
}

func (m *defaultShoppingCartZrpcClient) FindCart(ctx context.Context, in *FindCartReq, opts ...grpc.CallOption) (*FindCartResp, error) {
	client := pb.NewShoppingCartClient(m.cli.Conn())
	return client.FindCart(ctx, in, opts...)
}

func (m *defaultShoppingCartZrpcClient) GetCarts(ctx context.Context, in *FindCartsReq, opts ...grpc.CallOption) (*FindCartsResp, error) {
	client := pb.NewShoppingCartClient(m.cli.Conn())
	return client.GetCarts(ctx, in, opts...)
}

func (m *defaultShoppingCartZrpcClient) GetCartsByUserId(ctx context.Context, in *FindCartsByUserIdReq, opts ...grpc.CallOption) (*FindCartsByUserIdResp, error) {
	client := pb.NewShoppingCartClient(m.cli.Conn())
	return client.GetCartsByUserId(ctx, in, opts...)
}
