package logic

import (
	"QMall/shoppingCart/cmd/domain/model"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/svc"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCartLogic {
	return &FindCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindCartLogic) FindCart(in *pb.FindCartReq) (*pb.FindCartResp, error) {

	shoppingCart, err := l.svcCtx.ShoppingCartRepository.FindShoppingCart(in.UserId)
	if err != nil {
		return &pb.FindCartResp{}, err
	}

	return &pb.FindCartResp{
		ShoppingCart: model.ShppingCartModelConvertPb(shoppingCart),
	}, nil
}
