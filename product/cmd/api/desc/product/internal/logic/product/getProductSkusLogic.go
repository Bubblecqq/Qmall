package product

import (
	"QMall/product/cmd/domain/model"
	"QMall/product/cmd/rpc/product/product"
	"context"

	"QMall/product/cmd/api/desc/product/internal/svc"
	"QMall/product/cmd/api/desc/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSkusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取商品关联列表
func NewGetProductSkusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSkusLogic {
	return &GetProductSkusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductSkusLogic) GetProductSkus(req *types.GetProductSkuListReq) (resp *types.GetProductSkuListResp, err error) {

	productSku, err := l.svcCtx.ProductRpcConf.GetProductListSku(l.ctx, &product.GetProductSkuListReq{})

	responseProductSkuList := make([]types.ProductSku, len(productSku.ProductSkuList))

	for i := 0; i < len(productSku.ProductSkuList); i++ {
		responseProductSkuList[i] = types.ConvertResponseProductSku(model.PbProductSkuModelConvert(productSku.ProductSkuList[i]))
	}
	resp = new(types.GetProductSkuListResp)
	resp.ProductSkuList = responseProductSkuList

	return
}
