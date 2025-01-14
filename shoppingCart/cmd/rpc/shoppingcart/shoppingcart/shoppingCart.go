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
	AddCartReq                  = pb.AddCartReq
	AddCartResp                 = pb.AddCartResp
	DeleteCartsByUserIdReq      = pb.DeleteCartsByUserIdReq
	DeleteCartsByUserIdResp     = pb.DeleteCartsByUserIdResp
	FindCartReq                 = pb.FindCartReq
	FindCartResp                = pb.FindCartResp
	FindCartsByUserIdReq        = pb.FindCartsByUserIdReq
	FindCartsByUserIdResp       = pb.FindCartsByUserIdResp
	FindCartsReq                = pb.FindCartsReq
	FindCartsResp               = pb.FindCartsResp
	GetTotalPriceByUserIdReq    = pb.GetTotalPriceByUserIdReq
	GetTotalPriceByUserIdResp   = pb.GetTotalPriceByUserIdResp
	ProductSimple               = pb.ProductSimple
	ProductSkuSimple            = pb.ProductSkuSimple
	ShoppingCart                = pb.ShoppingCart
	ShoppingCartsProductInfo    = pb.ShoppingCartsProductInfo
	ShowDetailShoppingCartsReq  = pb.ShowDetailShoppingCartsReq
	ShowDetailShoppingCartsResp = pb.ShowDetailShoppingCartsResp
	UpdateCartReq               = pb.UpdateCartReq
	UpdateCartResp              = pb.UpdateCartResp

	ShoppingCartZrpcClient interface {
		AddCart(ctx context.Context, in *AddCartReq, opts ...grpc.CallOption) (*AddCartResp, error)
		UpdateCart(ctx context.Context, in *UpdateCartReq, opts ...grpc.CallOption) (*UpdateCartResp, error)
		FindCart(ctx context.Context, in *FindCartReq, opts ...grpc.CallOption) (*FindCartResp, error)
		GetCarts(ctx context.Context, in *FindCartsReq, opts ...grpc.CallOption) (*FindCartsResp, error)
		GetCartsByUserId(ctx context.Context, in *FindCartsByUserIdReq, opts ...grpc.CallOption) (*FindCartsByUserIdResp, error)
		DeleteCartsByUserId(ctx context.Context, in *DeleteCartsByUserIdReq, opts ...grpc.CallOption) (*DeleteCartsByUserIdResp, error)
		GetTotalPriceByUserId(ctx context.Context, in *GetTotalPriceByUserIdReq, opts ...grpc.CallOption) (*GetTotalPriceByUserIdResp, error)
		ShowDetailShoppingCarts(ctx context.Context, in *ShowDetailShoppingCartsReq, opts ...grpc.CallOption) (*ShowDetailShoppingCartsResp, error)
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

func (m *defaultShoppingCartZrpcClient) DeleteCartsByUserId(ctx context.Context, in *DeleteCartsByUserIdReq, opts ...grpc.CallOption) (*DeleteCartsByUserIdResp, error) {
	client := pb.NewShoppingCartClient(m.cli.Conn())
	return client.DeleteCartsByUserId(ctx, in, opts...)
}

func (m *defaultShoppingCartZrpcClient) GetTotalPriceByUserId(ctx context.Context, in *GetTotalPriceByUserIdReq, opts ...grpc.CallOption) (*GetTotalPriceByUserIdResp, error) {
	client := pb.NewShoppingCartClient(m.cli.Conn())
	return client.GetTotalPriceByUserId(ctx, in, opts...)
}

func (m *defaultShoppingCartZrpcClient) ShowDetailShoppingCarts(ctx context.Context, in *ShowDetailShoppingCartsReq, opts ...grpc.CallOption) (*ShowDetailShoppingCartsResp, error) {
	client := pb.NewShoppingCartClient(m.cli.Conn())
	return client.ShowDetailShoppingCarts(ctx, in, opts...)
}
