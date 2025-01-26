package logic

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDecreaseQuantityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDecreaseQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDecreaseQuantityLogic {
	return &GetDecreaseQuantityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDecreaseQuantityLogic) GetDecreaseQuantity(in *pb.GetDecreaseQuantityReq) (*pb.GetDecreaseQuantityResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetDecreaseQuantityResp{}, nil
}
