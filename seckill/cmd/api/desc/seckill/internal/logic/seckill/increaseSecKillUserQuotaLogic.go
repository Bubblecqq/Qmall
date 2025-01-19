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

type IncreaseSecKillUserQuotaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewIncreaseSecKillUserQuotaLogic 添加秒杀用户限额
func NewIncreaseSecKillUserQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillUserQuotaLogic {
	return &IncreaseSecKillUserQuotaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillUserQuotaLogic) IncreaseSecKillUserQuota(req *types.IncreaseSecKillUserQuotaReq) (resp *types.IncreaseSecKillUserQuotaResp, err error) {

	l.Info(fmt.Printf("[*] 正在添加秒杀限额信息>>>>>>当前访问的用户Id：%v，商品Id：%v，秒杀数量：%v\n", req.UserId, req.ProductsId, req.Num))

	quota, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillUserQuota(l.ctx, &seckill.IncreaseSecKillUserQuotaReq{
		UserId:     req.UserId,
		ProductsId: req.ProductsId,
		Num:        req.Num,
	})
	l.Info(fmt.Printf("[Info] 已添加秒杀用户限额信息！当前访问的用户Id：%v，商品Id：%v，秒杀数量：%v\n", req.UserId, req.ProductsId, req.Num))

	resp = new(types.IncreaseSecKillUserQuotaResp)
	resp.SecKillUserQuota = convert.PbSecKillUserQuotaConvertTypes(quota.SecKillUserQuota)
	return
}
