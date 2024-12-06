package product

import (
	"QMall/product/cmd/rpc/product/product"
	"context"

	"QMall/product/cmd/api/desc/product/internal/svc"
	"QMall/product/cmd/api/desc/product/internal/types"

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

	_, err = l.svcCtx.ProductRpcConf.DeleteProductSku(l.ctx, &product.DeleteProductSkuReq{
		Id: req.Id,
	})

	return
}
