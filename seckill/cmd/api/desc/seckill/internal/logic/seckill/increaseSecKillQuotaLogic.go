package seckill

import (
	"QMall/marketing/cmd/rpc/activity"
	"QMall/seckill/cmd/api/desc/seckill/internal/types/convert"
	"QMall/seckill/cmd/rpc/seckill"
	"context"
	"errors"
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
	resp = new(types.IncreaseSecKillQuotaResp)
	// 判断是否存在活动商品
	l.Info(fmt.Printf("[*] 正在查询当前请求的秒杀商品是否存在活动....\n"))

	activityProductByIdResp, err := l.svcCtx.ActivityRPC.GetActivityProductById(l.ctx, &activity.GetActivityProductByIdReq{
		ProductId: req.ProductId,
	})
	if activityProductByIdResp.ActivityProduct == nil {
		return
	}

	if req.LimitNumber <= 0 {
		err = errors.New(fmt.Sprintf("当前秒杀商品输入的秒杀限额数量不对！，秒杀限额商品Id：%v，秒杀限额数量：%v\n", req.ProductId, req.LimitNumber))
		return
	}

	l.Info(fmt.Printf("[*] 正在添加秒杀限额信息>>>>>>当前访问的商品Id：%v，秒杀限额数量：%v\n", req.ProductId, req.LimitNumber))

	quota, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillQuota(l.ctx, &seckill.IncreaseSecKillQuotaReq{
		ProductId:   req.ProductId,
		LimitNumber: req.LimitNumber,
	})
	l.Info(fmt.Printf("[Info] 已添加秒杀限额信息！当前访问的商品Id：%v，秒杀限额数量：%v\n", req.ProductId, req.LimitNumber))

	resp.SecKillQuota = convert.PbSecKillQuotaConvertTypes(quota.SecKillQuota)
	return
}
