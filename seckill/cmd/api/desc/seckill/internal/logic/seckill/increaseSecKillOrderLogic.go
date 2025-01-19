package seckill

import (
	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"QMall/seckill/cmd/api/desc/seckill/internal/types/convert"
	"QMall/seckill/cmd/rpc/seckill"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewIncreaseSecKillOrderLogic 添加秒杀订单
func NewIncreaseSecKillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillOrderLogic {
	return &IncreaseSecKillOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillOrderLogic) IncreaseSecKillOrder(req *types.IncreaseSecKillOrderReq) (resp *types.IncreaseSecKillOrderResp, err error) {
	resp = new(types.IncreaseSecKillOrderResp)
	// 查看商品是否上线
	secKillProductsByProductsId, err := l.svcCtx.SecKillRpcConf.GetSecKillProductsByProductsId(l.ctx, &seckill.GetSecKillProductsByProductsIdReq{
		ProductId: req.ProductsId,
	})
	if err != nil || secKillProductsByProductsId.SecKillProducts == nil {
		fmt.Printf("当前请求的秒杀商品不存在>>>>>>，请求的商品Id：%v\n", req.ProductsId)
		return
	}
	// 判断当前时间是否存在于活动时间内
	//secKillTime := time.Now()

	//
	l.Info(fmt.Printf("正在添加秒杀订单>>>>>>>>，当前买家信息：%v，卖家信息：%v，买家购买的商品Id：%v，购入%v商品的价格：%v\n", req.Buyer, req.Seller, req.ProductsId, req.ProductsId, req.Price))

	order, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillOrder(l.ctx, &seckill.IncreaseSecKillOrderReq{})

	// 添加秒杀用户限额
	// 查询限额
	//userQuota, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillUserQuota(l.ctx, &seckill.IncreaseSecKillUserQuotaReq{
	//	UserId:     req.Seller,
	//	ProductsId: req.ProductsId,
	//	//Num: req.,
	//})

	if err != nil {
		return
	}

	if err != nil {
		return
	}
	resp.SecKillOrder = convert.PbSecKillOrderConvertTypes(order.SecKillOrder)
	resp.SecKillProducts = convert.PbSecKillProductsConvertTypes(secKillProductsByProductsId.SecKillProducts)

	return
}
