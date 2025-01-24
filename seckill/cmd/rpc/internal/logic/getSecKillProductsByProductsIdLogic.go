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

type GetSecKillProductsByProductsIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSecKillProductsByProductsIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSecKillProductsByProductsIdLogic {
	return &GetSecKillProductsByProductsIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSecKillProductsByProductsIdLogic) GetSecKillProductsByProductsId(in *pb.GetSecKillProductsByProductsIdReq) (*pb.GetSecKillProductsByProductsIdResp, error) {

	fmt.Printf("正在查询秒杀商品信息.......\n")
	// 查询redis
	// 生成缓存键
	cacheKey := fmt.Sprintf("secKillProductsId:%d", in.GetProductId())
	fmt.Printf("正在查询秒杀商品信息>>>>>当前请求商品Id：%v\n", in.ProductId)
	var secKillProducts *model.SecKillProducts
	secKillProducts, err := l.svcCtx.SecKillRepository.GetSecKillProductsByProductsIdWithCache(l.ctx, cacheKey)
	if err != nil {

		if err.Error() != "cache miss" {
			return &pb.GetSecKillProductsByProductsIdResp{}, err
		}
		// 缓存未命中，从数据库查询
		fmt.Printf("当前缓存未命中，正在查询MYSQL....\n")
		secKillProducts, err = l.svcCtx.SecKillRepository.GetSecKillProductsByProductsId(in.ProductId)

		if err != nil {
			return &pb.GetSecKillProductsByProductsIdResp{}, err
		}
		// 将从数据库查询到的结果存入缓存
		err = l.svcCtx.SecKillRepository.SetSecKillProductsByProductsIdWithCache(l.ctx, cacheKey, secKillProducts)
		if err != nil {
			// 记录缓存设置失败的日志，但不影响返回结果
			fmt.Printf("设置秒杀商品信息缓存失败：%v\n", err)
		}
		fmt.Printf("设置秒杀商品信息缓存成功！\n")
	}

	// // 原只在mysql层
	//secKillProducts, err := l.svcCtx.SecKillRepository.GetSecKillProductsByProductsId(in.ProductId)
	//
	//if err != nil {
	//	return &pb.GetSecKillProductsByProductsIdResp{}, err
	//}

	fmt.Printf("已经查询到秒杀商品！商品信息如下：%v\n", secKillProducts)

	return &pb.GetSecKillProductsByProductsIdResp{
		SecKillProducts: convert.ModelSecKillProductsConvertPb(secKillProducts),
	}, nil
}
