package tradeOrder

import (
	"QMall/tradeOrder/cmd/rpc/tradeOrder/tradeorder"
	"context"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTradeOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewFindTradeOrderLogic 查询订单
func NewFindTradeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTradeOrderLogic {
	return &FindTradeOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTradeOrderLogic) FindTradeOrder(req *types.FindTradeOrderReq) (resp *types.FindTradeOrderResp, err error) {

	order, err := l.svcCtx.TradeOrderRpcConf.FindOrder(l.ctx, &tradeorder.FindOrderReq{
		Id:      req.Id,
		OrderNo: req.OrderNo,
	})
	l.Info(">>>>>>>>>>>>>>>>>>>>>>Result>>>>>>>>>>>>>>>>>>>>>")
	l.Info(order.TradeOrder)

	resp = new(types.FindTradeOrderResp)
	resp.TradeOrder = types.TradeOrderPbConvertTypes(order.TradeOrder)
	return
}
