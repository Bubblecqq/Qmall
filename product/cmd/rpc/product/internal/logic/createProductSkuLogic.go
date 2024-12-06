package logic

import (
	"context"

	"QMall/product/cmd/rpc/product/internal/svc"
	"QMall/product/cmd/rpc/product/pb"

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

	err := l.svcCtx.ProductRepository.CreateProductSku(in.ProductId, in.AttributeSymbolList, in.Name, in.SellPrice, in.CostPrice, in.Stock, in.StockWarn)

	if err != nil {
		return &pb.CreateProductSkuResp{}, err
	}

	return &pb.CreateProductSkuResp{}, nil
}
