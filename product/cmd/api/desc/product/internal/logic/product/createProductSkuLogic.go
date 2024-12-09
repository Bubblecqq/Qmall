package product

import (
	"QMall/product/cmd/rpc/product/product"
	"context"

	"QMall/product/cmd/api/desc/product/internal/svc"
	"QMall/product/cmd/api/desc/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建商品关联
func NewCreateProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductSkuLogic {
	return &CreateProductSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProductSkuLogic) CreateProductSku(req *types.CreateProductSkuReq) (resp *types.CreateProductSkuResp, err error) {

	_, err = l.svcCtx.ProductRpcConf.CreateProductSku(l.ctx, &product.CreateProductSkuReq{
		Name:                req.Name,
		ProductId:           req.ProductId,
		AttributeSymbolList: req.AttributeSymbolList,
		SellPrice:           req.SellPrice,
		CostPrice:           req.CostPrice,
		Stock:               req.Stock,
		StockWarn:           req.StockWarn,
	})
	resp = new(types.CreateProductSkuResp)
	return
}
