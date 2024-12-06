package productSku

import (
	"context"

	"goZeroDemo4/product/cmd/api/desc/productSku/internal/svc"
	"goZeroDemo4/product/cmd/api/desc/productSku/internal/types"

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

func (l *GetProductSkuByIdLogic) GetProductSkuById(req *types.GetProductSkuByIdReq) (resp *types.GetProductSkuByIdReq, err error) {
	// todo: add your logic here and delete this line

	return
}
