package seckill

import (
	"context"

	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillQuotaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加秒杀限额
func NewIncreaseSecKillQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillQuotaLogic {
	return &IncreaseSecKillQuotaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillQuotaLogic) IncreaseSecKillQuota(req *types.IncreaseSecKillQuotaReq) (resp *types.IncreaseSecKillQuotaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
