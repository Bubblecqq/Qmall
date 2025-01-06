package logic

import (
	"context"
	"fmt"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductOrderByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductOrderByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductOrderByUserIdLogic {
	return &GetProductOrderByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductOrderByUserIdLogic) GetProductOrderByUserId(in *pb.GetProductOrderByUserIdReq) (*pb.GetProductOrderByUserIdResp, error) {

	id, err := l.svcCtx.TradeOrderProductRepository.GetOrderProductByUserId(in.UserId)

	if err != nil {
		l.Info(fmt.Sprintf("Current Select userId:%v,Not Exist orderProduct!\n", in.UserId))
	}

	return &pb.GetProductOrderByUserIdResp{
		ProductOrder: pb.ModelProductOrderConvertPb(id),
	}, err
}
