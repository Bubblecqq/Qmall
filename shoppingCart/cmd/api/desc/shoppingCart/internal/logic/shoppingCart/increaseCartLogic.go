package shoppingCart

import (
	"QMall/product/cmd/rpc/product/pb"
	"QMall/product/cmd/rpc/product/product"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"context"

	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/svc"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加购物车
func NewIncreaseCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseCartLogic {
	return &IncreaseCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseCartLogic) IncreaseCart(req *types.IncreaseShoppingCartReq) (resp *types.IncreaseShoppingCartResp, err error) {
	l.Info(">>>>>>>>>>>>>>>>>>>InCreaseCart>>>>>>>>>>>>>>>>>>>")
	productDetail, err := l.svcCtx.Target.ShowProductDetail(l.ctx, &product.ShowProductDetailReq{
		Id: req.ProductId,
	})
	l.Info(">>>>>>>>>>>>>>>>>>>>>>>>ProductDetail>>>>>>>>>>>>>>>>>>>")
	l.Info(productDetail.ProductDetail)
	detailProduct := productDetail.ProductDetail
	productSku, err := l.svcCtx.Target.GetProductSku(l.ctx, &product.GetProductSkuByIdReq{
		Id: req.ProductSkuId,
	})
	addCartReq := &shoppingcart.AddCartReq{
		Number:       req.Number,
		ProductId:    req.ProductId,
		ProductSkuId: req.ProductSkuId,
		UserId:       req.UserId,
		//ProductName:        req.ProductName,
		//ProductMainPicture: req.ProductMainPicture,
		//Id:                 req.Id,
	}
	addCartReq.ProductName = productDetail.ProductDetail.Name
	addCartReq.ProductMainPicture = productDetail.ProductDetail.MainPicture

	_, err = l.svcCtx.ShoppingCartRpcConf.AddCart(l.ctx, addCartReq)

	resp = new(types.IncreaseShoppingCartResp)
	resp.ProductSimple = ProductDetailConvertTypeProductSimple(detailProduct)
	resp.ProductSkuSimple = ProductSkuConvertTypeProductSimple(productSku.ProductSku)
	return
}

func ProductDetailConvertTypeProductSimple(productDetail *pb.DetailProduct) types.Product {
	return types.Product{
		Id:             productDetail.Id,
		Name:           productDetail.Name,
		MainPicture:    productDetail.MainPicture,
		SingleBuyLimit: productDetail.SingleBuyLimit,
		IsEnabled:      productDetail.IsEnable == 1,
		ProductType:    productDetail.ProductType,
	}
}

func ProductSkuConvertTypeProductSimple(productSku *pb.ProductSku) types.ProductSku {
	return types.ProductSku{
		Id:                  productSku.Id,
		Name:                productSku.Name,
		AttributeSymbolList: productSku.AttributeSymbolList,
		ProductId:           productSku.ProductId,
		SellPrice:           productSku.SellPrice,
		CostPrice:           productSku.CostPrice,
		Stock:               productSku.Stock,
		IsEnabled:           productSku.IsEnable,
	}
}
