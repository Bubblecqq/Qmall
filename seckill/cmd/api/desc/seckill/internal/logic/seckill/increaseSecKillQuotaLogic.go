package seckill

import (
	"QMall/seckill/cmd/api/desc/seckill/internal/types/convert"
	"QMall/seckill/cmd/rpc/seckill"
	"context"
	"fmt"

	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillQuotaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewIncreaseSecKillQuotaLogic 添加秒杀限额
func NewIncreaseSecKillQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillQuotaLogic {
	return &IncreaseSecKillQuotaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillQuotaLogic) IncreaseSecKillQuota(req *types.IncreaseSecKillQuotaReq) (resp *types.IncreaseSecKillQuotaResp, err error) {

	l.Info(fmt.Printf("[*] 正在添加秒杀限额信息>>>>>>当前访问的商品Id：%v，秒杀限额数量：%v\n", req.ProductId, req.LimitNumber))

	quota, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillQuota(l.ctx, &seckill.IncreaseSecKillQuotaReq{
		ProductId:   req.ProductId,
		LimitNumber: req.LimitNumber,
	})
	l.Info(fmt.Printf("[Info] 已添加秒杀限额信息！当前访问的商品Id：%v，秒杀限额数量：%v\n", req.ProductId, req.LimitNumber))

	resp = new(types.IncreaseSecKillQuotaResp)
	resp.SecKillQuota = convert.PbSecKillQuotaConvertTypes(quota.SecKillQuota)
	return
}
