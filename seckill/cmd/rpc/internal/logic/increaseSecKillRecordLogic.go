package logic

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseSecKillRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillRecordLogic {
	return &IncreaseSecKillRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseSecKillRecordLogic) IncreaseSecKillRecord(in *pb.IncreaseSecKillRecordReq) (*pb.IncreaseSecKillRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IncreaseSecKillRecordResp{}, nil
}
