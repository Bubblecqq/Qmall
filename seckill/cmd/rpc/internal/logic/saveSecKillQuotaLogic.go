package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"QMall/seckill/cmd/domain/model"
	"context"
	"fmt"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveSecKillQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveSecKillQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveSecKillQuotaLogic {
	return &SaveSecKillQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveSecKillQuotaLogic) SaveSecKillQuota(in *pb.SaveSecKillQuotaReq) (*pb.SaveSecKillQuotaResp, error) {
	cacheKey := fmt.Sprintf("secKillQuota:%d", in.ProductsId)
	fmt.Printf("正在查询秒杀商品限额信息>>>>>当前请求商品Id：%v\n", in.ProductsId)
	var secKillQuota *model.SecKillQuota

	secKillQuota, err := l.svcCtx.SecKillRepository.GetSecKillQuotaWithCache(l.ctx, cacheKey)
	if err != nil {
		if err.Error() != "cache miss" {
			return &pb.SaveSecKillQuotaResp{}, err
		}
		// 缓存未命中，从数据库查询
		fmt.Printf("当前缓存未命中，正在查询MYSQL....\n")
		secKillQuota, err = l.svcCtx.SecKillRepository.GetSecKillQuotaByProductsId(in.ProductsId)
		if err != nil {
			return &pb.SaveSecKillQuotaResp{}, err
		}
		// 将从数据库查询到的结果存入缓存
		err = l.svcCtx.SecKillRepository.SetSecKillQuotaWithCache(l.ctx, cacheKey, secKillQuota)
		if err != nil {
			// 记录缓存设置失败的日志，但不影响返回结果
			fmt.Printf("设置活动商品信息缓存失败：%v\n", err)
		}
		fmt.Printf("设置秒杀商品限额信息缓存成功！\n")

	}
	fmt.Printf("当前商品限额信息：%v\n", secKillQuota)
	return &pb.SaveSecKillQuotaResp{
		SecKillQuota: convert.ModelSecKillQuotaConvertPb(secKillQuota),
	}, nil
}
