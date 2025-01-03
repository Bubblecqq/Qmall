package logic

import (
	"context"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrdersLogic {
	return &GetOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrdersLogic) GetOrders(in *pb.GetTradeOrderListReq) (*pb.GetTradeOrderListResp, error) {

	orders, err := l.svcCtx.TradeOrderRepository.GetTradeOrders()

	pbOrders := make([]*pb.TradeOrder, 0)
	for i := 0; i < len(orders); i++ {
		pbOrders = append(pbOrders, pb.ModelTradeOrderConvertPb(&orders[i]))
	}

	return &pb.GetTradeOrderListResp{
		TradeOrders: pbOrders,
	}, err
}
