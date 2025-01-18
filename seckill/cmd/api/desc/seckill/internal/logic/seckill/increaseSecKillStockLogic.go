package seckill

import (
	"context"

	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加秒杀库存
func NewIncreaseSecKillStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillStockLogic {
	return &IncreaseSecKillStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillStockLogic) IncreaseSecKillStock(req *types.IncreaseSecKillStockReq) (resp *types.IncreaseSecKillStockResp, err error) {
	// todo: add your logic here and delete this line

	return
}
