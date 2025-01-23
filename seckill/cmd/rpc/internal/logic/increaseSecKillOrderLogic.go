package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseSecKillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillOrderLogic {
	return &IncreaseSecKillOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseSecKillOrderLogic) IncreaseSecKillOrder(in *pb.IncreaseSecKillOrderReq) (*pb.IncreaseSecKillOrderResp, error) {

	//order, err := l.svcCtx.SecKillRepository.IncreaseSecKillOrder(in)

	order, err := l.svcCtx.KqPusherClient2.PusherWaitForPaymentSecKillOrder(l.ctx, in)

	if err != nil {
		return &pb.IncreaseSecKillOrderResp{}, err
	}

	return &pb.IncreaseSecKillOrderResp{
		SecKillOrder: convert.ModelSecKillOrderConvertPb(order),
	}, nil
}
