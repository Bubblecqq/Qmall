package logic

import (
	"QMall/shoppingCart/cmd/domain/model"
	"context"

	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/svc"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartsByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartsByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartsByUserIdLogic {
	return &GetCartsByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartsByUserIdLogic) GetCartsByUserId(in *pb.FindCartsByUserIdReq) (*pb.FindCartsByUserIdResp, error) {
	carts := make([]*pb.ShoppingCart, 0)

	shoppingCarts, err := l.svcCtx.ShoppingCartRepository.GetShoppingCartsByUserId(in.UserId)

	for i := 0; i < len(shoppingCarts); i++ {
		carts = append(carts, model.ShoppingCartModelConvertPb(&shoppingCarts[i]))
	}

	return &pb.FindCartsByUserIdResp{
		Carts: carts,
	}, err
}
