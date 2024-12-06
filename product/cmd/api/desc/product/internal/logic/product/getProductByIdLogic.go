package product

import (
	"QMall/product/cmd/domain/model"
	"QMall/product/cmd/rpc/product/product"
	"context"

	"QMall/product/cmd/api/desc/product/internal/svc"
	"QMall/product/cmd/api/desc/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据id获取商品
func NewGetProductByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductByIdLogic {
	return &GetProductByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductByIdLogic) GetProductById(req *types.GetProductByIdReq) (resp *types.GetProductByIdResp, err error) {
	getProduct, err := l.svcCtx.ProductRpcConf.GetProduct(l.ctx, &product.GetProductByIdReq{
		Id: req.Id,
	})
	resp.Product = types.ConvertResponseProduct(model.PbProductModelConvert(getProduct.Product))
	return
}
