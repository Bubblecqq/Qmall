package seckill

import (
	"QMall/marketing/cmd/rpc/activity"
	product2 "QMall/product/cmd/rpc/product/product"
	"QMall/seckill/cmd/api/desc/seckill/internal/types/convert"
	"QMall/seckill/cmd/rpc/seckill"
	"context"
	"errors"
	"fmt"

	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewIncreaseSecKillProductsLogic 添加秒杀商品
func NewIncreaseSecKillProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillProductsLogic {
	return &IncreaseSecKillProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillProductsLogic) IncreaseSecKillProducts(req *types.IncreaseSecKillProductsReq) (resp *types.IncreaseSecKillProductsResp, err error) {
	resp = new(types.IncreaseSecKillProductsResp)
	l.Info(fmt.Printf("[*] 正在查询当前请求的秒杀商品是否存在活动....\n"))

	activityProductByIdResp, err := l.svcCtx.ActivityRPC.GetActivityProductById(l.ctx, &activity.GetActivityProductByIdReq{
		ProductId: req.ProductId,
	})
	if activityProductByIdResp.ActivityProduct == nil {
		err = errors.New(fmt.Sprintf("当前请求的秒杀商品不存在！请求的商品Id：%v", req.ProductId))
		return
	}
	l.Info(fmt.Printf("当前活动商品信息：%v\n", activityProductByIdResp.GetActivityProduct()))
	product, err := l.svcCtx.ProductRPC.GetProduct(l.ctx, &product2.GetProductByIdReq{
		Id: req.ProductId,
	})

	// 检查秒杀价格
	if req.Price >= product.Product.StartingPrice {
		err = errors.New(fmt.Sprintf("秒杀价格应低于正常价格，商品ID：%v", req.ProductId))
		return
	}

	l.Info(fmt.Printf("[*] 正在添加秒杀商品信息>>>>>>当前访问的商品Id：%v，秒杀价格：%v，卖家信息：%v\n", req.ProductId, req.Price, req.Seller))

	products, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillProducts(l.ctx, &seckill.IncreaseSecKillProductsReq{
		ProductId:   req.ProductId,
		Seller:      req.Seller,
		Price:       req.Price,
		PictureUrl:  product.Product.MainPicture,
		ProductName: product.Product.Name,
		ProductsNum: activityProductByIdResp.ActivityProduct.ProductsNum,
	})

	l.Info(fmt.Printf("[INFO] 已添加秒杀商品信息！当前访问的商品Id：%v，秒杀价格：%v，卖家信息：%v\n", req.ProductId, req.Price, req.Seller))
	resp.SecKillProducts = convert.PbSecKillProductsConvertTypes(products.SecKillProducts)
	return
}
