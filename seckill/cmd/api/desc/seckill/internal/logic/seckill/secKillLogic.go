package seckill

import (
	"QMall/common"
	"QMall/marketing/cmd/rpc/activity"
	"QMall/seckill/cmd/api/desc/seckill/internal/types/convert"
	"QMall/seckill/cmd/rpc/seckill"
	"context"
	"errors"
	"fmt"
	"time"

	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SecKillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewSecKillLogic 添加秒杀订单V2
func NewSecKillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SecKillLogic {
	return &SecKillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SecKillLogic) SecKill(req *types.SecKillReq) (resp *types.SecKillResp, err error) {
	resp = new(types.SecKillResp)

	//先查秒杀订单商品
	productsId, err := l.svcCtx.SecKillRpcConf.GetSecKillProductsByProductsId(l.ctx, &seckill.GetSecKillProductsByProductsIdReq{
		ProductId: req.ProductsId,
	})
	if err != nil || productsId.SecKillProducts == nil {
		fmt.Printf("当前请求的秒杀商品不存在>>>>>>，请求的商品Id：%v\n", req.ProductsId)
		return
	}
	l.Info(fmt.Printf("[*] 当前商品存在秒杀活动>>>>>>%v\n", req.ProductsId))
	activityInfoResp, err := l.svcCtx.ActivityRPC.GetActivityInfoByProductsNum(l.ctx, &activity.GetActivityInfoByProductsNumReq{
		ProductsId:  req.ProductsId,
		ProductsNum: productsId.SecKillProducts.ProductsNum,
	})

	l.Info(fmt.Printf("当前秒杀活动商品信息：%v\n", activityInfoResp.ActivityInfo))
	if err != nil {
		return
	}
	// 判断当前时间是否存在于活动时间内
	secKillTime := time.Now()
	//const timeFormat = "15:04:05"

	startTime, err := time.Parse(time.RFC3339, activityInfoResp.ActivityInfo.StartTime)
	if err != nil {
		err = errors.New(fmt.Sprintf("解析开始时间错误: %v", err))
		// 处理解析错误
		fmt.Printf("解析开始时间错误: %v", err)
		return
	}
	endTime, err := time.Parse(time.RFC3339, activityInfoResp.ActivityInfo.EndTime)
	if err != nil {
		err = errors.New(fmt.Sprintf("解析结束时间错误: %v", err))
		// 处理解析错误
		fmt.Printf("解析结束时间错误: %v", err)
	}
	durationActivity := IsDurationActivity(secKillTime, startTime, endTime)
	if !durationActivity {
		err = errors.New("当前用户秒杀不在时间段")
		l.Info(fmt.Printf("当前用户秒杀不在时间段\n"))
		return
	}

	if err != nil {
		return
	}
	// 将全局限额 存储到redis
	secKillQuota, err := l.svcCtx.SecKillRpcConf.SaveSecKillQuota(l.ctx, &seckill.SaveSecKillQuotaReq{
		ProductsId: req.ProductsId,
	})
	if err != nil {
		return
	}
	// 添加秒杀用户限额
	// 查询限额
	//var killedUserQuota types.SecKillUserQuota
	//fmt.Printf("正在查询用户限额....\n")
	//killedUserQuota, err := l.svcCtx.SecKillRpcConf.SaveSecKillUserQuota(l.ctx, &seckill.SaveSecKillUserQuotaReq{
	//	UserId:       req.Buyer,
	//	ProductsId:   req.ProductsId,
	//	Num:          req.Quantity,
	//	KilledNum:    0,
	//	SecKillQuota: secKillQuota.SecKillQuota,
	//})
	//if err != nil {
	//	return
	//}
	// 将stock 存储到redis
	//
	fmt.Printf("正在准备库存预扣>>>>>>>\n")
	secKillStock, err := l.svcCtx.SecKillRpcConf.SaveSecKillStock(l.ctx, &seckill.SaveSecKillStockReq{
		ProductsId: req.ProductsId,
		Quantity:   req.Quantity,
		UserId:     req.Buyer,
		Limit:      secKillQuota.SecKillQuota.Num,
	})
	// 订单生成
	fmt.Printf("当前库存数：%v\n", secKillStock.Stock)
	l.Info(fmt.Printf("正在添加秒杀订单>>>>>>>>，当前买家信息：%v，卖家信息：%v，买家购买的商品Id：%v，购入%v商品的价格：%v\n", req.Buyer, req.Seller, req.ProductsId, req.ProductsId, productsId.SecKillProducts.Price))
	now := time.Now()
	OrderNum := common.GenerateOrderNum(now, req.Buyer)

	order, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillOrder(l.ctx, &seckill.IncreaseSecKillOrderReq{
		Seller:      req.Seller,
		Buyer:       req.Buyer,
		ProductsId:  req.ProductsId,
		Price:       productsId.SecKillProducts.Price,
		ProductsNum: productsId.SecKillProducts.ProductsNum,
		OrderNum:    OrderNum,
		Quantity:    req.Quantity,
	})

	killRecord, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillRecord(l.ctx, &seckill.IncreaseSecKillRecordReq{
		UserId:     req.Buyer,
		Price:      productsId.SecKillProducts.Price,
		ProductsId: req.ProductsId,
		OrderNo:    OrderNum,
		SecNo:      common.GenerateSecKillOrderNo(now, req.Buyer),
		SkuId:      activityInfoResp.ActivityInfo.SkuId,
		Quantity:   req.Quantity,
	})
	if err != nil {
		return
	}

	resp.SecKillOrder = convert.PbSecKillOrderConvertTypes(order.SecKillOrder)
	//resp.SecKillUserQuota = convert.PbSecKillUserQuotaConvertTypes(killedUserQuota.SecKillUserQuota)
	resp.SecKillRecord = convert.PbSecKillRecordConvertTypes(killRecord.SecKillRecord)
	//resp.RemindStock = secKillStock.Stock
	return
}
