package logic

import (
	"context"
	"goZeroDemo4/product/cmd/domain/model"

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

	productSkuList := l.svcCtx.ProductRepository.GetProductSkuList()

	pbProductSkuList := make([]*pb.ProductSku, len(productSkuList))
	for i := 0; i < len(productSkuList); i++ {
		pbProductSkuList[i] = model.ProductSkuModelConvertPbProductSku(&productSkuList[i])
	}

	return &pb.GetProductSkuListResp{
		ProductSkuList: pbProductSkuList,
	}, nil
}
