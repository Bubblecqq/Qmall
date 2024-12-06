package product

import (
	"context"
	"goZeroDemo4/product/cmd/domain/model"
	"goZeroDemo4/product/cmd/rpc/product/product"

	"goZeroDemo4/product/cmd/api/desc/product/internal/svc"
	"goZeroDemo4/product/cmd/api/desc/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSkuByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据id获取商品关联
func NewGetProductSkuByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSkuByIdLogic {
	return &GetProductSkuByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductSkuByIdLogic) GetProductSkuById(req *types.GetProductSkuByIdReq) (resp *types.GetProductSkuByIdResp, err error) {

	sku, err := l.svcCtx.ProductRpcConf.GetProductSku(l.ctx, &product.GetProductSkuByIdReq{
		Id: req.Id,
	})
	resp.ProductSku = types.ConvertResponseProductSku(model.PbProductSkuModelConvert(sku.ProductSku))
	return
}
