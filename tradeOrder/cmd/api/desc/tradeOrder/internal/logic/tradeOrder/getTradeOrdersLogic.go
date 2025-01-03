package tradeOrder

import (
	"QMall/tradeOrder/cmd/rpc/tradeOrder/tradeorder"
	"context"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTradeOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetTradeOrdersLogic 获取订单列表
func NewGetTradeOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTradeOrdersLogic {
	return &GetTradeOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTradeOrdersLogic) GetTradeOrders(req *types.GetTradeOrderListReq) (resp *types.GetTradeOrderListResp, err error) {
	orders := make([]types.TradeOrder, 0)
	getOrders, err := l.svcCtx.TradeOrderRpcConf.GetOrders(l.ctx, &tradeorder.GetTradeOrderListReq{})
	l.Info(">>>>>>>>>>>>>>>>>>>>>>Result>>>>>>>>>>>>>>>>>>>>>")
	for i := 0; i < len(getOrders.TradeOrders); i++ {
		orders = append(orders, types.TradeOrderPbConvertTypes(getOrders.TradeOrders[i]))
	}
	l.Info(orders)
	resp = new(types.GetTradeOrderListResp)
	resp.TradeOrders = orders
	return
}
