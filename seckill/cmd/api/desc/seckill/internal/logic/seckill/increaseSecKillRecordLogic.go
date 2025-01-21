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

type IncreaseSecKillRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewIncreaseSecKillRecordLogic 添加秒杀记录
func NewIncreaseSecKillRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillRecordLogic {
	return &IncreaseSecKillRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillRecordLogic) IncreaseSecKillRecord(req *types.IncreaseSecKillRecordReq) (resp *types.IncreaseSecKillRecordResp, err error) {
	// 判断是否存在活动商品
	l.Info(fmt.Printf("[*] 正在查询当前请求的秒杀商品是否存在活动....\n"))

	activityProductByIdResp, err := l.svcCtx.ActivityRPC.GetActivityProductById(l.ctx, &activity.GetActivityProductByIdReq{
		ProductId: req.ProductsId,
	})
	if activityProductByIdResp.ActivityProduct == nil {
		err = errors.New(fmt.Sprintf("当前请求的秒杀商品不存在！请求的商品Id：%v", req.ProductsId))
		return
	}

	//设置订单为秒杀

	l.Info(fmt.Printf("[*] 正在添加用户秒杀记录信息>>>>>>当前访问的用户Id：%v，商品Id：%v\n", req.UserId, req.ProductsId))

	record, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillRecord(l.ctx, &seckill.IncreaseSecKillRecordReq{
		UserId:     req.UserId,
		ProductsId: req.ProductsId,
		OrderNo:    req.OrderNum,
		Price:      req.Price,
		Quantity:   req.Quantity,
		SkuId:      req.SkuId,
	})
	l.Info(fmt.Printf("[Info] 已在添加用户秒杀记录信息>>>>>>当前访问的用户Id：%v，商品Id：%v\n", req.UserId, req.ProductsId))

	resp = new(types.IncreaseSecKillRecordResp)
	resp.SecKillRecord = convert.PbSecKillRecordConvertTypes(record.SecKillRecord)
	return
}
