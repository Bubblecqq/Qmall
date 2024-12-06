package logic

import (
	"context"

	"QMall/product/cmd/rpc/product/internal/svc"
	"QMall/product/cmd/rpc/product/pb"

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

	productSku, err := l.svcCtx.ProductRepository.FindProductSku(in.Id)

	if err != nil {
		return &pb.DeleteProductSkuResp{}, err
	}
	l.Info("[ProductRPC] The product_sku ID to be deleted currently is", productSku.Id)

	err = l.svcCtx.ProductRepository.DeleteProductSku(productSku.Id)
	if err != nil {
		return &pb.DeleteProductSkuResp{}, err
	}
	return &pb.DeleteProductSkuResp{}, nil
}
