package logic

import (
	"context"

	"goZeroDemo4/product/cmd/rpc/product/internal/svc"
	"goZeroDemo4/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *pb.DeleteProductReq) (*pb.DeleteProductResp, error) {

	product, err := l.svcCtx.ProductRepository.FindProduct(in.Id)

	if err != nil {
		return &pb.DeleteProductResp{}, err
	}
	l.Info("[ProductRPC] The current user ID to be deleted is", product.Id)
	err = l.svcCtx.ProductRepository.DeleteProduct(in.Id)

	if err != nil {
		return &pb.DeleteProductResp{}, err
	}

	return &pb.DeleteProductResp{}, nil
}
