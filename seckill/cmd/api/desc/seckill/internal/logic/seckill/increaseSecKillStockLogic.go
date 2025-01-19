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

type IncreaseSecKillStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewIncreaseSecKillStockLogic 添加秒杀库存
func NewIncreaseSecKillStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillStockLogic {
	return &IncreaseSecKillStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillStockLogic) IncreaseSecKillStock(req *types.IncreaseSecKillStockReq) (resp *types.IncreaseSecKillStockResp, err error) {

	// 判断是否存在活动商品
	l.Info(fmt.Printf("[*] 正在查询当前请求的秒杀商品是否存在活动....\n"))

	activityProductByIdResp, err := l.svcCtx.ActivityRPC.GetActivityProductById(l.ctx, &activity.GetActivityProductByIdReq{
		ProductId: req.ProductsId,
	})
	if activityProductByIdResp.ActivityProduct == nil {
		err = errors.New(fmt.Sprintf("当前请求的秒杀商品不存在！请求的商品Id：%v", req.ProductsId))
		return
	}

	l.Info(fmt.Printf("[*] 正在添加秒杀库存信息>>>>>>当前需要添加的商品Id：%v，秒杀库存：%v\n", req.ProductsId, req.Stock))

	stock, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillStock(l.ctx, &seckill.IncreaseSecKillStockReq{
		ProductsId: req.ProductsId,
		Stock:      req.Stock,
	})
	l.Info(fmt.Printf("[Info] 已在添加秒杀库存信息>>>>>>当前需要添加的商品Id：%v，秒杀库存：%v\n", req.ProductsId, req.Stock))

	resp = new(types.IncreaseSecKillStockResp)
	resp.SecKillStock = convert.PbSecKillStockConvertTypes(stock.SecKillStock)
	return
}
