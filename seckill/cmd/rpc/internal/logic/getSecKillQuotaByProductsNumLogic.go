package logic

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecKillQuotaByProductsNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSecKillQuotaByProductsNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSecKillQuotaByProductsNumLogic {
	return &GetSecKillQuotaByProductsNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSecKillQuotaByProductsNumLogic) GetSecKillQuotaByProductsNum(in *pb.GetSecKillQuotaByProductsIdReq) (*pb.GetSecKillQuotaByProductsIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetSecKillQuotaByProductsIdResp{}, nil
}
