package logic

import (
	"context"

	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/svc"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowDetailShoppingCartsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShowDetailShoppingCartsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowDetailShoppingCartsLogic {
	return &ShowDetailShoppingCartsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShowDetailShoppingCartsLogic) ShowDetailShoppingCarts(in *pb.ShowDetailShoppingCartsReq) (*pb.ShowDetailShoppingCartsResp, error) {

	detailList, err := l.svcCtx.ShoppingCartRepository.ShowShoppingCartsDetailList(in.UserId)
	var pbDetailList []*pb.ShoppingCartsProductInfo
	for i := 0; i < len(detailList); i++ {
		pbDetailList = append(pbDetailList, pb.ShoppingCartsDetailConvertPb(detailList[i]))
	}
	return &pb.ShowDetailShoppingCartsResp{
		UserId:                   in.UserId,
		ShoppingCartsNumber:      int64(len(detailList)),
		ShoppingCartsProductInfo: pbDetailList,
	}, err
}
