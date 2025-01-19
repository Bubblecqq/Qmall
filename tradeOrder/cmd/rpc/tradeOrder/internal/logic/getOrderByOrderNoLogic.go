package logic

import (
	"context"
	"fmt"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByOrderNoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByOrderNoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByOrderNoLogic {
	return &GetOrderByOrderNoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByOrderNoLogic) GetOrderByOrderNo(in *pb.GetOrderByOrderNoReq) (*pb.GetOrderByOrderNoResp, error) {
	fmt.Printf("正在查询到订单信息！>>>>>>请求订单编号：%v\n", in.OrderNo)

	orderNo, err := l.svcCtx.TradeOrderRepository.GetTradeOrderByOrderNo(in.OrderNo)

	if err != nil {
		return &pb.GetOrderByOrderNoResp{}, err
	}
	fmt.Printf("已查询到订单信息！>>>>>>订单信息：%v\n", orderNo)
	return &pb.GetOrderByOrderNoResp{
		TradeOrder: pb.ModelTradeOrderConvertPb(orderNo),
	}, nil
}
