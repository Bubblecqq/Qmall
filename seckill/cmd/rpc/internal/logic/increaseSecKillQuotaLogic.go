package logic

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseSecKillQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillQuotaLogic {
	return &IncreaseSecKillQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseSecKillQuotaLogic) IncreaseSecKillQuota(in *pb.IncreaseSecKillQuotaReq) (*pb.IncreaseSecKillQuotaResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IncreaseSecKillQuotaResp{}, nil
}
