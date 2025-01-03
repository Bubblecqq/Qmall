package tradeOrder

import (
	"QMall/tradeOrder/cmd/rpc/tradeOrder/tradeorder"
	"context"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTradeOrdersByPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取订单列表
func NewGetTradeOrdersByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTradeOrdersByPageLogic {
	return &GetTradeOrdersByPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTradeOrdersByPageLogic) GetTradeOrdersByPage(req *types.PageTradeOrderReq) (resp *types.PageTradeOrderResp, err error) {
	orders := make([]types.TradeOrder, 0)
	page, err := l.svcCtx.TradeOrderRpcConf.GetTradeOrdersByPage(l.ctx, &tradeorder.PageTradeOrderReq{
		Length:    req.Length,
		PageIndex: req.PageIndex,
	})
	l.Info(">>>>>>>>>>>>>>>>>>>>>>Page Result>>>>>>>>>>>>>>>>>>>>>")
	for i := 0; i < len(page.TradeOrders); i++ {
		orders = append(orders, types.TradeOrderPbConvertTypes(page.TradeOrders[i]))
	}
	l.Info(orders)
	resp = new(types.PageTradeOrderResp)
	resp.ResponsePageTradeOrderIndex = orders
	return
}
