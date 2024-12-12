package shoppingCart

import (
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"context"

	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/svc"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查看购物车
func NewFindCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCartLogic {
	return &FindCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCartLogic) FindCart(req *types.FindShoppingCartReq) (resp *types.FindShoppingCartResp, err error) {

	_, err = l.svcCtx.ShoppingCartRpcConf.FindCart(l.ctx, &shoppingcart.FindCartReq{
		Id:     req.Id,
		UserId: req.UserId,
	})
	resp = new(types.FindShoppingCartResp)

	return
}
