package logic

import (
	"context"
	"goZeroDemo4/product/cmd/domain/model"

	"goZeroDemo4/product/cmd/rpc/product/internal/svc"
	"goZeroDemo4/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductListLogic) GetProductList(in *pb.GetProductListReq) (*pb.GetProductListResp, error) {

	productList := l.svcCtx.ProductRepository.GetProductList()

	pbProductList := make([]*pb.Product, len(productList))
	for i := 0; i < len(productList); i++ {
		pbProductList[i] = model.ProductModelConvertPbProduct(&productList[i])
	}

	return &pb.GetProductListResp{
		ProductList: pbProductList,
	}, nil
}
