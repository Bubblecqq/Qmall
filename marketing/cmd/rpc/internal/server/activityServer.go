// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: marketing.proto

package server

import (
	"context"

	"QMall/marketing/cmd/rpc/internal/logic"
	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"
)

type ActivityServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedActivityServer
}

func NewActivityServer(svcCtx *svc.ServiceContext) *ActivityServer {
	return &ActivityServer{
		svcCtx: svcCtx,
	}
}

func (s *ActivityServer) AddActivity(ctx context.Context, in *pb.AddActivityReq) (*pb.AddActivityResp, error) {
	l := logic.NewAddActivityLogic(ctx, s.svcCtx)
	return l.AddActivity(in)
}

func (s *ActivityServer) AddActivityTime(ctx context.Context, in *pb.AddActivityTimeReq) (*pb.AddActivityTimeResp, error) {
	l := logic.NewAddActivityTimeLogic(ctx, s.svcCtx)
	return l.AddActivityTime(in)
}

func (s *ActivityServer) AddActivityProduct(ctx context.Context, in *pb.AddActivityProductReq) (*pb.AddActivityProductResp, error) {
	l := logic.NewAddActivityProductLogic(ctx, s.svcCtx)
	return l.AddActivityProduct(in)
}

func (s *ActivityServer) AddActivityProductSku(ctx context.Context, in *pb.AddActivityProductSkuReq) (*pb.AddActivityProductSkuResp, error) {
	l := logic.NewAddActivityProductSkuLogic(ctx, s.svcCtx)
	return l.AddActivityProductSku(in)
}
