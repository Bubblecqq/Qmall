package logic

import (
	"context"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTradeOrdersByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTradeOrdersByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTradeOrdersByPageLogic {
	return &GetTradeOrdersByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTradeOrdersByPageLogic) GetTradeOrdersByPage(in *pb.PageTradeOrderReq) (*pb.PageTradeOrderResp, error) {

	page, err := l.svcCtx.TradeOrderRepository.Page(in.Length, in.PageIndex)

	pbPage := make([]*pb.TradeOrder, 0)
	for i := 0; i < len(page); i++ {
		pbPage = append(pbPage, pb.ModelTradeOrderConvertPb(&page[i]))
	}
	return &pb.PageTradeOrderResp{
		TradeOrders: pbPage,
	}, err
}
