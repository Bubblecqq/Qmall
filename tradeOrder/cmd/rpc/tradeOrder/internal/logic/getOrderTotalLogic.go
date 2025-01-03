package logic

import (
	"context"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderTotalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderTotalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderTotalLogic {
	return &GetOrderTotalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderTotalLogic) GetOrderTotal(in *pb.OrderTotalReq) (*pb.OrderTotalResp, error) {

	return &pb.OrderTotalResp{}, nil
}
