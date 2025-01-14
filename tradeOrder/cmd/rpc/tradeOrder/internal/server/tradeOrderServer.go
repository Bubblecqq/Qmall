// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: tradeOrder.proto

package server

import (
	"context"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/logic"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"
)

type TradeOrderServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTradeOrderServer
}

func NewTradeOrderServer(svcCtx *svc.ServiceContext) *TradeOrderServer {
	return &TradeOrderServer{
		svcCtx: svcCtx,
	}
}

func (s *TradeOrderServer) AddTradeOrder(ctx context.Context, in *pb.AddTradeOrderReq) (*pb.AddTradeOrderResp, error) {
	l := logic.NewAddTradeOrderLogic(ctx, s.svcCtx)
	return l.AddTradeOrder(in)
}

func (s *TradeOrderServer) UpdateTradeOrder(ctx context.Context, in *pb.AddTradeOrderReq) (*pb.AddTradeOrderResp, error) {
	l := logic.NewUpdateTradeOrderLogic(ctx, s.svcCtx)
	return l.UpdateTradeOrder(in)
}

func (s *TradeOrderServer) GetOrderTotal(ctx context.Context, in *pb.OrderTotalReq) (*pb.OrderTotalResp, error) {
	l := logic.NewGetOrderTotalLogic(ctx, s.svcCtx)
	return l.GetOrderTotal(in)
}

func (s *TradeOrderServer) FindOrder(ctx context.Context, in *pb.FindOrderReq) (*pb.FindOrderResp, error) {
	l := logic.NewFindOrderLogic(ctx, s.svcCtx)
	return l.FindOrder(in)
}

func (s *TradeOrderServer) GetOrders(ctx context.Context, in *pb.GetTradeOrderListReq) (*pb.GetTradeOrderListResp, error) {
	l := logic.NewGetOrdersLogic(ctx, s.svcCtx)
	return l.GetOrders(in)
}

func (s *TradeOrderServer) GetTradeOrdersByPage(ctx context.Context, in *pb.PageTradeOrderReq) (*pb.PageTradeOrderResp, error) {
	l := logic.NewGetTradeOrdersByPageLogic(ctx, s.svcCtx)
	return l.GetTradeOrdersByPage(in)
}

func (s *TradeOrderServer) BatchCreateOrderProduct(ctx context.Context, in *pb.AddProductOrderReq) (*pb.AddProductOrderResp, error) {
	l := logic.NewBatchCreateOrderProductLogic(ctx, s.svcCtx)
	return l.BatchCreateOrderProduct(in)
}

func (s *TradeOrderServer) GetProductOrderList(ctx context.Context, in *pb.GetProductOrderListReq) (*pb.GetProductOrderListResp, error) {
	l := logic.NewGetProductOrderListLogic(ctx, s.svcCtx)
	return l.GetProductOrderList(in)
}

func (s *TradeOrderServer) GetProductOrderByUserId(ctx context.Context, in *pb.GetProductOrderByUserIdReq) (*pb.GetProductOrderByUserIdResp, error) {
	l := logic.NewGetProductOrderByUserIdLogic(ctx, s.svcCtx)
	return l.GetProductOrderByUserId(in)
}
