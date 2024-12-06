package logic

import (
	"context"

	"goZeroDemo4/product/cmd/rpc/product/internal/svc"
	"goZeroDemo4/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *pb.CreateProductReq) (*pb.CreateProductResp, error) {

	err := l.svcCtx.ProductRepository.CreateProduct(in.Name, in.ProductType, in.CategoryId, in.StartingPrice, in.TotalStock, in.MainPicture, in.RemoteAreaPostage, in.SingleBuyLimit, in.Remark)
	if err != nil {
		return &pb.CreateProductResp{}, err
	}

	return &pb.CreateProductResp{}, nil
}
