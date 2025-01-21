package seckill

import (
	"QMall/seckill/cmd/api/desc/seckill/internal/types/convert"
	"QMall/seckill/cmd/rpc/seckill"
	"context"
	"errors"
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
	resp = new(types.IncreaseSecKillUserQuotaResp)
	// 先判断用户限额是否存在，存在则返回
	isExistUserQuota, err := l.svcCtx.SecKillRpcConf.GetSecKillUserQuota(l.ctx, &seckill.GetSecKillUserQuotaReq{
		UserId:     req.UserId,
		ProductsId: req.ProductsId,
	})
	if err != nil {
		err = errors.New(fmt.Sprintf("当前用户请求的秒杀商品不存在！用户请求的商品Id：%v，秒杀数量：%v", req.ProductsId, req.Num))
		return
	}
	if isExistUserQuota.SecKillUserQuota != nil {
		err = errors.New(fmt.Sprintf("当前用户Id：%v，商品Id：%v存在限额，请勿重复添加！", req.UserId, req.ProductsId))
		resp.SecKillUserQuota = convert.PbSecKillUserQuotaConvertTypes(isExistUserQuota.SecKillUserQuota)
		return
	}
	l.Info(fmt.Printf("[*] 正在添加用户秒杀限额信息>>>>>>当前访问的用户Id：%v，商品Id：%v，秒杀数量：%v\n", req.UserId, req.ProductsId, req.Num))

	// 判断Num是否大于限额
	quotaByProductsIdResp, err := l.svcCtx.SecKillRpcConf.GetSecKillQuotaByProductsId(l.ctx, &seckill.GetSecKillQuotaByProductsIdReq{
		ProductId: req.ProductsId,
	})
	if err != nil {
		err = errors.New(fmt.Sprintf("当前用户请求的秒杀商品不存在！用户请求的商品Id：%v，秒杀数量：%v", req.ProductsId, req.Num))
		return
	}

	if req.Num <= 0 {
		err = errors.New(fmt.Sprintf("当前用户秒杀限额设置参数不正确！当前输入的用户限额数：%v", req.Num))
		return
	}

	if req.Num > quotaByProductsIdResp.SecKillQuota.Num {
		err = errors.New(fmt.Sprintf("当前用户已经达到了限额！用户请求的秒杀数量：%v，最大限额：%v", req.Num, quotaByProductsIdResp.SecKillQuota.Num))
		return
	}

	quota, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillUserQuota(l.ctx, &seckill.IncreaseSecKillUserQuotaReq{
		UserId:     req.UserId,
		ProductsId: req.ProductsId,
		Num:        req.Num,
	})
	l.Info(fmt.Printf("[Info] 已添加秒杀用户限额信息！当前访问的用户Id：%v，商品Id：%v，秒杀数量：%v\n", req.UserId, req.ProductsId, req.Num))

	resp.SecKillUserQuota = convert.PbSecKillUserQuotaConvertTypes(quota.SecKillUserQuota)
	return
}
