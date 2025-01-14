package logic

import (
	"context"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchCreateOrderProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchCreateOrderProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchCreateOrderProductLogic {
	return &BatchCreateOrderProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchCreateOrderProductLogic) BatchCreateOrderProduct(in *pb.AddProductOrderReq) (*pb.AddProductOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddProductOrderResp{}, nil
}
