package product

import (
	"QMall/product/cmd/domain/model"
	"QMall/product/cmd/rpc/product/product"
	"context"

	"QMall/product/cmd/api/desc/product/internal/svc"
	"QMall/product/cmd/api/desc/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取商品列表
func NewGetProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductsLogic {
	return &GetProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductsLogic) GetProducts(req *types.GetProductListReq) (resp *types.GetProductListResp, err error) {

	products, err := l.svcCtx.ProductRpcConf.GetProductList(l.ctx, &product.GetProductListReq{})

	responseProductList := make([]types.Product, len(products.ProductList))

	for i := 0; i < len(products.ProductList); i++ {
		responseProductList[i] = types.ConvertResponseProduct(model.PbProductModelConvert(products.ProductList[i]))
	}

	resp.ProductList = responseProductList
	return
}
