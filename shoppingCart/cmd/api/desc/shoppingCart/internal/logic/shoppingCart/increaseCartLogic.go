package shoppingCart

import (
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"context"

	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/svc"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加购物车
func NewIncreaseCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseCartLogic {
	return &IncreaseCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseCartLogic) IncreaseCart(req *types.IncreaseShoppingCartRep) (resp *types.IncreaseShoppingCartResp, err error) {

	_, err = l.svcCtx.ShoppingCartRpcConf.AddCart(l.ctx, &shoppingcart.AddCartReq{
		Number:             req.Number,
		ProductId:          req.ProductId,
		ProductSkuId:       req.ProductSkuId,
		UserId:             req.UserId,
		ProductName:        req.ProductName,
		ProductMainPicture: req.ProductMainPicture,
		Id:                 req.Id,
	})
	resp = new(types.IncreaseShoppingCartResp)
	return
}
