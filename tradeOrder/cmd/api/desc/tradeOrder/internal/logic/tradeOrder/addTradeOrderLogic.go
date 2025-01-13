package tradeOrder

import (
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"QMall/tradeOrder/cmd/domain/model"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/tradeorder"
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
	var pbCarts []*pb.ShoppingCartSimple
	for i := 0; i < len(carts); i++ {
		pbCarts = append(pbCarts, pb.ShoppingCartsConvert(carts[i]))
	}
	l.Info(fmt.Printf("当前用户ID：%v，购物车信息如下：%v\n", req.UserId, carts))
	// 计算购物车订单总价
	totalAmount := 0.0
	priceByUserId, err := l.svcCtx.ShoppingCartRPC.GetTotalPriceByUserId(l.ctx, &shoppingcart.GetTotalPriceByUserIdReq{
		UserId: req.UserId,
	})
	//total, err := l.svcCtx.TradeOrderRpcConf.GetOrderTotal(l.ctx, &tradeorder.OrderTotalReq{
	//	Carts: pbCarts,
	//})
	totalAmount = priceByUserId.TotalPrice
	// 计算购物车订单总价
	shippingAmount := 0.0
	discountAmount := 0.0
	order := &model.TradeOrder{TotalAmount: totalAmount, ShippingAmount: shippingAmount, DiscountAmount: discountAmount, PayAmount: totalAmount + shippingAmount - discountAmount}
	addTradeOrder := &tradeorder.AddTradeOrderReq{
		TradeOrder: pb.ModelTradeOrderConvertPb(order),
	}

	tradeOrder, err := l.svcCtx.TradeOrderRpcConf.AddTradeOrder(l.ctx, addTradeOrder)

	//addProductOrder:= &tradeorder.AddProductOrderReq{}
	//addProductOrder.OrderId = tradeOrder.TradeOrder.Id

	resp = new(types.AddTradeOrderResp)
	resp.TradeOrder = types.TradeOrderPbConvertTypes(tradeOrder.TradeOrder)
	return
}
