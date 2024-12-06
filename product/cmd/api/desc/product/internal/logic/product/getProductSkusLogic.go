package product

import (
	"context"

	"goZeroDemo4/product/cmd/api/desc/product/internal/svc"
	"goZeroDemo4/product/cmd/api/desc/product/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
