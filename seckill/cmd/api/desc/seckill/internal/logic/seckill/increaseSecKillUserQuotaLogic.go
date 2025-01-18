package seckill

import (
	"context"

	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillUserQuotaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加秒杀用户限额
func NewIncreaseSecKillUserQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillUserQuotaLogic {
	return &IncreaseSecKillUserQuotaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillUserQuotaLogic) IncreaseSecKillUserQuota(req *types.IncreaseSecKillUserQuotaReq) (resp *types.IncreaseSecKillUserQuotaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
