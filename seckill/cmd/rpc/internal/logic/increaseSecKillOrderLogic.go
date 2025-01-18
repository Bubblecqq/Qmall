package logic

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseSecKillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillOrderLogic {
	return &IncreaseSecKillOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseSecKillOrderLogic) IncreaseSecKillOrder(in *pb.IncreaseSecKillOrderReq) (*pb.IncreaseSecKillOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IncreaseSecKillOrderResp{}, nil
}
