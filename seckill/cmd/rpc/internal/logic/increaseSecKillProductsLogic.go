package logic

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseSecKillProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillProductsLogic {
	return &IncreaseSecKillProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseSecKillProductsLogic) IncreaseSecKillProducts(in *pb.IncreaseSecKillProductsReq) (*pb.IncreaseSecKillProductsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IncreaseSecKillProductsResp{}, nil
}
