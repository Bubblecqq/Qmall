package logic

import (
	"context"

	"QMall/product/cmd/rpc/product/internal/svc"
	"QMall/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSkuLogic {
	return &UpdateProductSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductSkuLogic) UpdateProductSku(in *pb.UpdateProductSkuReq) (*pb.UpdateProductSkuResp, error) {

	err := l.svcCtx.ProductRepository.UpdateProductSku(in)

	return &pb.UpdateProductSkuResp{
		IsSuccess: err == nil,
	}, nil
}
