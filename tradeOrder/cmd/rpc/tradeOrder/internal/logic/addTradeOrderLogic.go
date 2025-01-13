package logic

import (
	"context"
	"fmt"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTradeOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTradeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTradeOrderLogic {
	return &AddTradeOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTradeOrderLogic) AddTradeOrder(in *pb.AddTradeOrderReq) (*pb.AddTradeOrderResp, error) {

	l.Info(fmt.Sprintf("正在生成订单，用户Id为:%v.\n", in.TradeOrder.UserId))
	order, err := l.svcCtx.TradeOrderRepository.AddTradeOrder(in)

	if err != nil {
		l.Info(fmt.Sprintf("用户Id:%v,订单生成失败！,详情见：%v\n", in.TradeOrder.UserId, err))
	}
	l.Info(fmt.Sprintf("用户Id：%v，订单已生成！\n", in.TradeOrder.UserId))

	return &pb.AddTradeOrderResp{
		TradeOrder: pb.ModelTradeOrderConvertPb(order),
	}, nil
}
