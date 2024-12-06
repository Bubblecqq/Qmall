package logic

import (
	"QMall/product/cmd/domain/model"
	"context"

	"QMall/product/cmd/rpc/product/internal/svc"
	"QMall/product/cmd/rpc/product/pb"

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
