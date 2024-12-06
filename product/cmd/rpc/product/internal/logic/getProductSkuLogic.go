package logic

import (
	"context"
	"goZeroDemo4/product/cmd/domain/model"

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

	productSku, err := l.svcCtx.ProductRepository.FindProductSku(in.Id)

	if err != nil {
		return &pb.GetProductSkuByIdResp{}, err
	}

	return &pb.GetProductSkuByIdResp{
		ProductSku: model.ProductSkuModelConvertPbProductSku(productSku),
	}, nil
}
