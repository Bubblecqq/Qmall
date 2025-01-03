package logic

import (
	"QMall/shoppingCart/cmd/domain/model"
	"context"

	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/svc"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartsLogic {
	return &GetCartsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartsLogic) GetCarts(in *pb.FindCartsReq) (*pb.FindCartsResp, error) {

	carts := make([]*pb.ShoppingCart, 0)

	shoppingCarts, err := l.svcCtx.ShoppingCartRepository.GetShoppingCarts()

	for i := 0; i < len(shoppingCarts); i++ {
		carts = append(carts, model.ShoppingCartModelConvertPb(&shoppingCarts[i]))
	}

	return &pb.FindCartsResp{
		Carts: carts,
	}, err
}
