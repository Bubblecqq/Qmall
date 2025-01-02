package logic

import (
	"context"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOrderLogic {
	return &FindOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOrderLogic) FindOrder(in *pb.FindOrderReq) (*pb.FindOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FindOrderResp{}, nil
}
