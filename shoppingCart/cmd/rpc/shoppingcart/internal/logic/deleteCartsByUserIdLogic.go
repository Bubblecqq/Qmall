package logic

import (
	"context"
	"fmt"

	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/svc"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCartsByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCartsByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCartsByUserIdLogic {
	return &DeleteCartsByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCartsByUserIdLogic) DeleteCartsByUserId(in *pb.DeleteCartsByUserIdReq) (*pb.DeleteCartsByUserIdResp, error) {

	l.Info(fmt.Sprintf("正在清空购物车，userId：%v\n", in.UserId))

	err := l.svcCtx.ShoppingCartRepository.DeleteShoppingCartByUserId(in.UserId)
	if err == nil {
		l.Info(fmt.Sprintf("用户Id：%v的购物车已清空！\n", in.UserId))
	}
	return &pb.DeleteCartsByUserIdResp{}, err
}
