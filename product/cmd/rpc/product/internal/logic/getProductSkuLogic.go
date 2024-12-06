package logic

import (
	"context"

	"goZeroDemo4/product/cmd/rpc/product/internal/svc"
	"goZeroDemo4/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductSkuLogic {
	return &GetProductSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductSkuLogic) GetProductSku(in *pb.GetProductSkuByIdReq) (*pb.GetProductSkuByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetProductSkuByIdResp{}, nil
}
