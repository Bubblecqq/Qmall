package logic

import (
	"context"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTradeOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTradeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTradeOrderLogic {
	return &UpdateTradeOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTradeOrderLogic) UpdateTradeOrder(in *pb.UpdateTradeOrderReq) (*pb.UpdateTradeOrderResp, error) {
	//order, err := l.svcCtx.TradeOrderRepository.GetTradeOrderByUserId(in.OrderId, in.UserId)

	order, err := l.svcCtx.TradeOrderRepository.UpdateTradeOrder2(in)
	return &pb.UpdateTradeOrderResp{
		TradeOrder: pb.ModelTradeOrderConvertPb(order),
	}, err
}
