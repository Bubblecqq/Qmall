package tradeOrder

import (
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"context"
	"fmt"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTradeOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成订单
func NewAddTradeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTradeOrderLogic {
	return &AddTradeOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTradeOrderLogic) AddTradeOrder(req *types.AddTradeOrderReq) (resp *types.AddTradeOrderResp, err error) {

	cartsByUserId, err := l.svcCtx.ShoppingCartRPC.GetCartsByUserId(l.ctx, &shoppingcart.FindCartsByUserIdReq{
		UserId: req.UserId,
	})
	carts := cartsByUserId.Carts
	l.Info(fmt.Printf("当前用户ID：%v，购物车信息如下：%v\n", req.UserId, carts))
	//计算订单总价

	resp = new(types.AddTradeOrderResp)
	return
}
