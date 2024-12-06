package logic

import (
	"context"
	"goZeroDemo4/product/cmd/domain/model"

	"goZeroDemo4/product/cmd/rpc/product/internal/svc"
	"goZeroDemo4/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *pb.GetProductByIdReq) (*pb.GetProductByIdResp, error) {

	getProduct, err := l.svcCtx.ProductRepository.FindProduct(in.Id)

	if err != nil {
		return &pb.GetProductByIdResp{}, err
	}

	return &pb.GetProductByIdResp{
		Product: model.ProductModelConvertPbProduct(getProduct),
	}, nil
}
