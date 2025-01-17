package logic

import (
	"QMall/marketing/cmd/domain/convert"
	product2 "QMall/product/cmd/rpc/product/product"
	"context"
	"fmt"

	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityProductSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddActivityProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityProductSkuLogic {
	return &AddActivityProductSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddActivityProductSkuLogic) AddActivityProductSku(in *pb.AddActivityProductSkuReq) (*pb.AddActivityProductSkuResp, error) {

	productSku, err := l.svcCtx.ProductRPC.GetProductSku(l.ctx, &product2.GetProductSkuByIdReq{
		Id: in.ProductSkuId,
	})

	if err != nil {
		fmt.Printf("当前商品库存Id：%v不存在！原因见：%v\n", in.ProductSkuId, err)
		return &pb.AddActivityProductSkuResp{}, err
	}
	fmt.Printf("正在添加活动商品库存信息>>>>>>>>商品库存Id：%v，活动商品Id：%v\n", in.ProductSkuId, in.ActivityProductId)

	sku, err := l.svcCtx.ActivityRepository.AddActivityProductSku(in.ActivityProductId, in.ProductId, in.ProductSkuId, in.UserId, in.Price, productSku.ProductSku.SellPrice, in.Number, productSku.ProductSku.Stock)

	if err != nil {
		return &pb.AddActivityProductSkuResp{}, err
	}
	fmt.Printf("活动商品库存商品Id：%v已添加！>>>>>>>>\n", sku.ActivityProductID)
	return &pb.AddActivityProductSkuResp{
		ActivityProductSku: convert.ModelActivityProductSkuConvertPb(sku),
	}, nil
}
