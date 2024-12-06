package logic

import (
	"context"

	"goZeroDemo4/product/cmd/rpc/product/internal/svc"
	"goZeroDemo4/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductSkuLogic {
	return &CreateProductSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductSkuLogic) CreateProductSku(in *pb.CreateProductSkuReq) (*pb.CreateProductSkuResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateProductSkuResp{}, nil
}
