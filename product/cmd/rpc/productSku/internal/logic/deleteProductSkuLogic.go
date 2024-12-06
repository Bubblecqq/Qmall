package logic

import (
	"context"

	"goZeroDemo4/product/cmd/rpc/productSku/internal/svc"
	"goZeroDemo4/product/cmd/rpc/productSku/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSkuLogic {
	return &DeleteProductSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductSkuLogic) DeleteProductSku(in *pb.DeleteProductSkuReq) (*pb.DeleteProductSkuResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteProductSkuResp{}, nil
}
