package logic

import (
	"context"

	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/svc"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTotalPriceByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTotalPriceByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTotalPriceByUserIdLogic {
	return &GetTotalPriceByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTotalPriceByUserIdLogic) GetTotalPriceByUserId(in *pb.GetTotalPriceByUserIdReq) (*pb.GetTotalPriceByUserIdResp, error) {

	totalPrice := l.svcCtx.ShoppingCartRepository.GetTotalPriceByUserId(in.UserId)

	return &pb.GetTotalPriceByUserIdResp{
		TotalPrice: totalPrice,
	}, nil
}
