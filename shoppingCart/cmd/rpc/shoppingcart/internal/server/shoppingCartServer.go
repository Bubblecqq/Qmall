// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: cart.proto

package server

import (
	"context"

	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/logic"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/svc"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"
)

type ShoppingCartServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedShoppingCartServer
}

func NewShoppingCartServer(svcCtx *svc.ServiceContext) *ShoppingCartServer {
	return &ShoppingCartServer{
		svcCtx: svcCtx,
	}
}

func (s *ShoppingCartServer) AddCart(ctx context.Context, in *pb.AddCartReq) (*pb.AddCartResp, error) {
	l := logic.NewAddCartLogic(ctx, s.svcCtx)
	return l.AddCart(in)
}

func (s *ShoppingCartServer) UpdateCart(ctx context.Context, in *pb.UpdateCartReq) (*pb.UpdateCartResp, error) {
	l := logic.NewUpdateCartLogic(ctx, s.svcCtx)
	return l.UpdateCart(in)
}

func (s *ShoppingCartServer) FindCart(ctx context.Context, in *pb.FindCartReq) (*pb.FindCartResp, error) {
	l := logic.NewFindCartLogic(ctx, s.svcCtx)
	return l.FindCart(in)
}

func (s *ShoppingCartServer) GetCarts(ctx context.Context, in *pb.FindCartsReq) (*pb.FindCartsResp, error) {
	l := logic.NewGetCartsLogic(ctx, s.svcCtx)
	return l.GetCarts(in)
}
