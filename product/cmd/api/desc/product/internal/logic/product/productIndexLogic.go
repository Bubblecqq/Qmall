package product

import (
	"QMall/product/cmd/domain/model"
	"QMall/product/cmd/rpc/product/product"
	"context"

	"QMall/product/cmd/api/desc/product/internal/svc"
	"QMall/product/cmd/api/desc/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商品分页
func NewProductIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductIndexLogic {
	return &ProductIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductIndexLogic) ProductIndex(req *types.PageReq) (resp *types.PageResp, err error) {

	pageIndex, err := l.svcCtx.ProductRpcConf.PageIndex(l.ctx, &product.PageReq{
		Length:    req.Length,
		PageIndex: req.PageIndex,
	})

	responsePageIndex := make([]types.Product, len(pageIndex.PageProductList))

	for i := 0; i < len(responsePageIndex); i++ {
		responsePageIndex[i] = types.ConvertResponseProduct(model.PbProductModelConvert(pageIndex.PageProductList[i]))
	}
	resp = new(types.PageResp)
	resp.ResponsePageIndex = responsePageIndex
	return
}
