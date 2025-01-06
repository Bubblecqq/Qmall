package logic

import (
	"context"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductOrderLogic {
	return &AddProductOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProductOrderLogic) AddProductOrder(in *pb.AddProductOrderReq) (*pb.AddProductOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddProductOrderResp{}, nil
}
