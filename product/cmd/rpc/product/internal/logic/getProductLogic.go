package logic

import (
	"QMall/product/cmd/domain/model"
	"context"
	"errors"
	"fmt"

	"QMall/product/cmd/rpc/product/internal/svc"
	"QMall/product/cmd/rpc/product/pb"

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
	if in.Id <= 0 {
		fmt.Printf("当前查询的Id：%v商品不存在！\n", in.Id)
		return &pb.GetProductByIdResp{}, errors.New(fmt.Sprintf("当前查询的Id：%v商品不存在!", in.Id))
	}
	getProduct, err := l.svcCtx.ProductRepository.FindProduct(in.Id)

	if err != nil {
		return &pb.GetProductByIdResp{}, err
	}

	return &pb.GetProductByIdResp{
		Product: model.ProductModelConvertPbProduct(getProduct),
	}, nil
}
