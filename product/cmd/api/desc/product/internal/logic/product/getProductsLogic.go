package product

import (
	"context"

	"goZeroDemo4/product/cmd/api/desc/product/internal/svc"
	"goZeroDemo4/product/cmd/api/desc/product/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
