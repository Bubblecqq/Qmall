package shoppingCart

import (
	"QMall/product/cmd/rpc/product/pb"
	"QMall/product/cmd/rpc/product/product"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"QMall/user/cmd/rpc/user"
	"context"
	"errors"
	"fmt"

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
	l.Info("Current UUID:", req.Token)
	resp = new(types.IncreaseShoppingCartResp)
	l.Info(">>>>>>>>>>>>>>>>>>>InCreaseCart>>>>>>>>>>>>>>>>>>>")
	productDetail, err := l.svcCtx.ProductRPC.ShowProductDetail(l.ctx, &product.ShowProductDetailReq{
		Id: req.ProductId,
	})

	// Token
	tokenResp, err := l.svcCtx.UserRPC.GetUserToken(l.ctx, &user.TokenReq{
		Uuid: req.Token,
	})
	token := tokenResp.Token
	if err != nil || tokenResp.IsLogin {
		return resp, errors.New("当前您未登录！")
	}
	l.Info("Current Login User Token:", token)

	l.Info(">>>>>>>>>>>>>>>>>>>>>>>>ProductDetail>>>>>>>>>>>>>>>>>>>")
	l.Info(productDetail.ProductDetail)
	detailProduct := productDetail.ProductDetail
	productSku, err := l.svcCtx.ProductRPC.GetProductSku(l.ctx, &product.GetProductSkuByIdReq{
		Id: req.ProductSkuId,
	})
	sku := productSku.ProductSku
	stock := sku.Stock
	sku.Stock -= req.Number
	resp.IsBeyondMaxLimit = false
	resp.CanSetShoppingCartNumber = stock
	updateStock := &product.UpdateProductSkuReq{
		ProductSku: sku,
	}
	updateProductSku, err := l.svcCtx.ProductRPC.UpdateProductSku(l.ctx, updateStock)
	// 添加数量大于库存时提示添加失败
	if req.Number > resp.CanSetShoppingCartNumber {
		resp.IsBeyondMaxLimit = true
		return resp, errors.New(fmt.Sprintf("超出库存设置范围！请重新添加，当前商品库存：%v", stock))
	}
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

	resp.ProductSimple = ProductDetailConvertTypeProductSimple(detailProduct)
	resp.ProductSkuSimple = ProductSkuConvertTypeProductSimple(productSku.ProductSku)
	resp.ShoppingCartNumber = req.Number
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
