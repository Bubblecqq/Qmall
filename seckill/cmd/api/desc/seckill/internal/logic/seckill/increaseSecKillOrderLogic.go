package seckill

import (
	"QMall/common"
	"QMall/marketing/cmd/rpc/activity"
	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"QMall/seckill/cmd/api/desc/seckill/internal/types/convert"
	"QMall/seckill/cmd/rpc/seckill"
	"context"
	"errors"
	"fmt"
	"time"

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
	// 添加秒杀用户限额
	// 查询限额
	//var killedUserQuota types.SecKillUserQuota
	killedUserQuota, err := l.svcCtx.SecKillRpcConf.SaveSecKillUserQuota(l.ctx, &seckill.SaveSecKillUserQuotaReq{
		UserId:     req.Buyer,
		ProductsId: req.ProductsId,
		Num:        req.Quantity,
		KilledNum:  0,
	})
	if err != nil {
		return
	}
	//userQuota, err := l.svcCtx.SecKillRpcConf.GetSecKillUserQuota(l.ctx, &seckill.GetSecKillUserQuotaReq{
	//	UserId:     req.Buyer,
	//	ProductsId: req.ProductsId,
	//})
	//// 当前不存在限额
	//if err != nil || userQuota.SecKillUserQuota.Id <= 0 {
	//	// 创建 需要判断秒杀数量是否大于限额
	//	killUserQuota, err2 := l.svcCtx.SecKillRpcConf.IncreaseSecKillUserQuota(l.ctx, &seckill.IncreaseSecKillUserQuotaReq{
	//		UserId:     req.Buyer,
	//		ProductsId: req.ProductsId,
	//		Num:        req.Quantity,
	//		KilledNum:  0,
	//	})
	//	if err2 != nil {
	//		return
	//	}
	//	killedUserQuota = convert.PbSecKillUserQuotaConvertTypes(killUserQuota.SecKillUserQuota)
	//} else {
	//	// 存在则修改限额
	//	updateSecKillUserQuotaById, err2 := l.svcCtx.SecKillRpcConf.UpdateSecKillUserQuotaById(l.ctx, &seckill.UpdateSecKillUserQuotaByIdReq{
	//		ProductsId:  req.ProductsId,
	//		Num:         req.Quantity,
	//		UserId:      req.Buyer,
	//		ProductsNum: productsId.SecKillProducts.ProductsNum,
	//	})
	//	if err2 != nil {
	//		return
	//	}
	//	killedUserQuota = convert.PbSecKillUserQuotaConvertTypes(updateSecKillUserQuotaById.SecKillUserQuota)
	//}

	//
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
	// 创建秒杀记录

	// 库存预扣在这里
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

	//userQuota, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillUserQuota(l.ctx, &seckill.IncreaseSecKillUserQuotaReq{
	//	UserId:     req.Seller,
	//	ProductsId: req.ProductsId,
	//	//Num: req.,
	//})

	if err != nil {
		return
	}
	resp.SecKillOrder = convert.PbSecKillOrderConvertTypes(order.SecKillOrder)
	resp.SecKillProducts = convert.PbSecKillProductsConvertTypes(productsId.SecKillProducts)
	resp.SecKillUserQuota = convert.PbSecKillUserQuotaConvertTypes(killedUserQuota.SecKillUserQuota)
	resp.SecKillRecord = convert.PbSecKillRecordConvertTypes(killRecord.SecKillRecord)
	return
}

func IsDurationActivity(secKillTime, startTime, endTime time.Time) bool {
	fmt.Printf("当前秒杀时间：%v\n", secKillTime.String())
	fmt.Printf("活动时间段，开始时间：%v，结束时间：%v\n", startTime, endTime)
	//if secKillTime.Hour() > startTime.Hour() && secKillTime.Hour() < endTime.Hour() {
	//	return true
	//}
	//if secKillTime.Hour() == startTime.Hour() {
	//	if secKillTime.Minute() > startTime.Minute() {
	//		return true
	//	}
	//	if secKillTime.Minute() == startTime.Minute() && secKillTime.Second() > startTime.Second() {
	//		return true
	//	}
	//}
	//
	//if secKillTime.Hour() == endTime.Hour() {
	//	if secKillTime.Minute() < endTime.Minute() {
	//		return true
	//	}
	//	if secKillTime.Minute() == endTime.Minute() && secKillTime.Second() < endTime.Second() {
	//		return true
	//	}
	//}
	return secKillTime.After(startTime) && secKillTime.Before(endTime)
}
