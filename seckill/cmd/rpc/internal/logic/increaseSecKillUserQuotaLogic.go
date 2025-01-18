package logic

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillUserQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseSecKillUserQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillUserQuotaLogic {
	return &IncreaseSecKillUserQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseSecKillUserQuotaLogic) IncreaseSecKillUserQuota(in *pb.IncreaseSecKillUserQuotaReq) (*pb.IncreaseSecKillUserQuotaResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IncreaseSecKillUserQuotaResp{}, nil
}
