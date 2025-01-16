package tradeOrder

import (
	"QMall/tradeOrder/cmd/rpc/tradeOrder/tradeorder"
	"context"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTradeOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUpdateTradeOrderLogic 修改订单
func NewUpdateTradeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTradeOrderLogic {
	return &UpdateTradeOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTradeOrderLogic) UpdateTradeOrder(req *types.UpdateTradeOrderReq) (resp *types.UpdateTradeOrderResp, err error) {

	order, err := l.svcCtx.TradeOrderRpcConf.UpdateTradeOrder(l.ctx, &tradeorder.UpdateTradeOrderReq{
		UserId:       req.UserId,
		PayType:      req.PayType,
		IsAfterSale:  req.IsAfterSale,
		CancelReason: req.CancelReason,
		OrderId:      req.OrderId,
		IsRefund:     req.IsRefund,
	})
	resp = new(types.UpdateTradeOrderResp)
	resp.TradeOrder = types.TradeOrderPbConvertTypes(order.TradeOrder)
	return
}
