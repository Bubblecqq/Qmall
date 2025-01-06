package logic

import (
	"context"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductOrderListLogic {
	return &GetProductOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductOrderListLogic) GetProductOrderList(in *pb.GetProductOrderListReq) (*pb.GetProductOrderListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetProductOrderListResp{}, nil
}
