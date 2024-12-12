package logic

import (
	"context"

	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/svc"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartLogic {
	return &AddCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCartLogic) AddCart(in *pb.AddCartReq) (*pb.AddCartResp, error) {

	_, err := l.svcCtx.ShoppingCartRepository.InCreaseShoppingCart(in.Number, in.ProductId, in.ProductSkuId, in.UserId, in.ProductName, in.ProductMainPicture)
	if err != nil {
		return &pb.AddCartResp{}, err
	}

	return &pb.AddCartResp{}, nil
}
