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

func (l *UpdateTradeOrderLogic) UpdateTradeOrder(in *pb.AddTradeOrderReq) (*pb.AddTradeOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddTradeOrderResp{}, nil
}
