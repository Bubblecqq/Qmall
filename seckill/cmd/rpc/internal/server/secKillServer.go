// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: seckill.proto

package server

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/logic"
	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"
)

type SecKillServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSecKillServer
}

func NewSecKillServer(svcCtx *svc.ServiceContext) *SecKillServer {
	return &SecKillServer{
		svcCtx: svcCtx,
	}
}

func (s *SecKillServer) IncreaseSecKillOrder(ctx context.Context, in *pb.IncreaseSecKillOrderReq) (*pb.IncreaseSecKillOrderResp, error) {
	l := logic.NewIncreaseSecKillOrderLogic(ctx, s.svcCtx)
	return l.IncreaseSecKillOrder(in)
}

func (s *SecKillServer) IncreaseSecKillProducts(ctx context.Context, in *pb.IncreaseSecKillProductsReq) (*pb.IncreaseSecKillProductsResp, error) {
	l := logic.NewIncreaseSecKillProductsLogic(ctx, s.svcCtx)
	return l.IncreaseSecKillProducts(in)
}

func (s *SecKillServer) IncreaseSecKillQuota(ctx context.Context, in *pb.IncreaseSecKillQuotaReq) (*pb.IncreaseSecKillQuotaResp, error) {
	l := logic.NewIncreaseSecKillQuotaLogic(ctx, s.svcCtx)
	return l.IncreaseSecKillQuota(in)
}

func (s *SecKillServer) IncreaseSecKillUserQuota(ctx context.Context, in *pb.IncreaseSecKillUserQuotaReq) (*pb.IncreaseSecKillUserQuotaResp, error) {
	l := logic.NewIncreaseSecKillUserQuotaLogic(ctx, s.svcCtx)
	return l.IncreaseSecKillUserQuota(in)
}

func (s *SecKillServer) IncreaseSecKillStock(ctx context.Context, in *pb.IncreaseSecKillStockReq) (*pb.IncreaseSecKillStockResp, error) {
	l := logic.NewIncreaseSecKillStockLogic(ctx, s.svcCtx)
	return l.IncreaseSecKillStock(in)
}

func (s *SecKillServer) IncreaseSecKillRecord(ctx context.Context, in *pb.IncreaseSecKillRecordReq) (*pb.IncreaseSecKillRecordResp, error) {
	l := logic.NewIncreaseSecKillRecordLogic(ctx, s.svcCtx)
	return l.IncreaseSecKillRecord(in)
}
