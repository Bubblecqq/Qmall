package seckill

import (
	"context"

	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewIncreaseSecKillProductsLogic 添加秒杀商品
func NewIncreaseSecKillProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillProductsLogic {
	return &IncreaseSecKillProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillProductsLogic) IncreaseSecKillProducts(req *types.IncreaseSecKillProductsReq) (resp *types.IncreaseSecKillProductsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
