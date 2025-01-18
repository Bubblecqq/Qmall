package seckill

import (
	"context"

	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加秒杀记录
func NewIncreaseSecKillRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillRecordLogic {
	return &IncreaseSecKillRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillRecordLogic) IncreaseSecKillRecord(req *types.IncreaseSecKillRecordReq) (resp *types.IncreaseSecKillRecordResp, err error) {
	// todo: add your logic here and delete this line

	return
}
