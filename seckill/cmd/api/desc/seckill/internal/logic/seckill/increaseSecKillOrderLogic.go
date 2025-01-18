package seckill

import (
	"context"

	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加秒杀订单
func NewIncreaseSecKillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillOrderLogic {
	return &IncreaseSecKillOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillOrderLogic) IncreaseSecKillOrder(req *types.IncreaseSecKillOrderReq) (resp *types.IncreaseSecKillOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
