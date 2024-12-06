package productSku

import (
	"context"

	"goZeroDemo4/product/cmd/api/desc/productSku/internal/svc"
	"goZeroDemo4/product/cmd/api/desc/productSku/internal/types"

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

func (l *CreateProductSkuLogic) CreateProductSku(req *types.CreateProductSkuReq) (resp *types.CreateProductSkuReq, err error) {
	// todo: add your logic here and delete this line

	return
}
