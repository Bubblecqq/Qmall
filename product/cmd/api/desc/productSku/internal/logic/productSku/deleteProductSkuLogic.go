package productSku

import (
	"context"

	"goZeroDemo4/product/cmd/api/desc/productSku/internal/svc"
	"goZeroDemo4/product/cmd/api/desc/productSku/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据id删除商品关联
func NewDeleteProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSkuLogic {
	return &DeleteProductSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductSkuLogic) DeleteProductSku(req *types.DeleteProductSkuReq) (resp *types.DeleteProductSkuResp, err error) {
	// todo: add your logic here and delete this line

	return
}
