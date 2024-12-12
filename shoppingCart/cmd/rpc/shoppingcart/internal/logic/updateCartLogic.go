package logic

import (
	"context"

	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/svc"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartLogic {
	return &UpdateCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCartLogic) UpdateCart(in *pb.UpdateCartReq) (*pb.UpdateCartResp, error) {

	_, err := l.svcCtx.ShoppingCartRepository.UpdateShoppingCart(in)

	if err != nil {
		return &pb.UpdateCartResp{}, err
	}

	return &pb.UpdateCartResp{}, nil
}
