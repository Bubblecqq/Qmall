package logic

import (
	"context"

	"goZeroDemo4/product/cmd/rpc/product/internal/svc"
	"goZeroDemo4/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListSkuLogic {
	return &GetProductListSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductListSkuLogic) GetProductListSku(in *pb.GetProductSkuListReq) (*pb.GetProductSkuListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetProductSkuListResp{}, nil
}
