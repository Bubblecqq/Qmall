package product

import (
	"QMall/product/cmd/rpc/product/pb"
	"QMall/product/cmd/rpc/product/product"
	"context"
	"fmt"

	"QMall/product/cmd/api/desc/product/internal/svc"
	"QMall/product/cmd/api/desc/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取商品详情
func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDetailLogic) ProductDetail(req *types.ShowProductDetailReq) (resp *types.ShowProductDetailResp, err error) {

	detail, err := l.svcCtx.ProductRpcConf.ShowProductDetail(l.ctx, &product.ShowProductDetailReq{
		Id: req.Id,
	})
	productDetail := detail.ProductDetail
	resp = new(types.ShowProductDetailResp)
	resp.DetailProduct = pbProductDetailConvertTypesModel(productDetail)
	fmt.Println(">>>>>>>>>>", resp.DetailProduct)
	l.Info(">>>>>>>>>", resp.DetailProduct)
	return
}

func pbProductDetailConvertTypesModel(productDetail *pb.DetailProduct) types.DetailProduct {
	return types.DetailProduct{
		Id:                productDetail.Id,
		Name:              productDetail.Name,
		ProductType:       productDetail.ProductType,
		CategoryId:        productDetail.CategoryId,
		StartingPrice:     productDetail.StartingPrice,
		TotalStock:        productDetail.TotalStock,
		MainPicture:       productDetail.MainPicture,
		RemoteAreaPostage: productDetail.RemoteAreaPostage,
		SingleBuyLimit:    productDetail.SingleBuyLimit,
		Remark:            productDetail.Remark,
		Detail:            productDetail.Detail,
		PictureList:       productDetail.PictureList,
		IsEnabled:         productDetail.IsEnable == 1,
	}
}
