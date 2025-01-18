package logic

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseSecKillStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillStockLogic {
	return &IncreaseSecKillStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseSecKillStockLogic) IncreaseSecKillStock(in *pb.IncreaseSecKillStockReq) (*pb.IncreaseSecKillStockResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IncreaseSecKillStockResp{}, nil
}
