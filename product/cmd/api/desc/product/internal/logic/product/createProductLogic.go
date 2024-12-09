package product

import (
	"QMall/product/cmd/rpc/product/product"
	"context"

	"QMall/product/cmd/api/desc/product/internal/svc"
	"QMall/product/cmd/api/desc/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建商品
func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProductLogic) CreateProduct(req *types.CreateProductReq) (resp *types.CreateProductResp, err error) {

	_, err = l.svcCtx.ProductRpcConf.CreateProduct(l.ctx, &product.CreateProductReq{
		Name:              req.Name,
		ProductType:       req.ProductType,
		CategoryId:        req.CategoryId,
		StartingPrice:     req.StartingPrice,
		TotalStock:        req.TotalStock,
		MainPicture:       req.MainPicture,
		RemoteAreaPostage: req.RemoteAreaPostage,
		SingleBuyLimit:    req.SingleBuyLimit,
		Remark:            req.Remark,
	})
	resp = new(types.CreateProductResp)
	return
}
