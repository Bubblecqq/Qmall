package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"
	"fmt"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveSecKillUserQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveSecKillUserQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveSecKillUserQuotaLogic {
	return &SaveSecKillUserQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveSecKillUserQuotaLogic) SaveSecKillUserQuota(in *pb.SaveSecKillUserQuotaReq) (*pb.SaveSecKillUserQuotaResp, error) {
	fmt.Printf("[SaveSecKillUserQuotaLogic] 正在调用持久化数据层.....\n")
	quota, err := l.svcCtx.SecKillRepository.SaveSecKillUserQuota(in)
	if err != nil {
		return &pb.SaveSecKillUserQuotaResp{}, err
	}
	fmt.Printf("[SaveSecKillUserQuotaLogic] 持久化数据执行成功！\n")
	return &pb.SaveSecKillUserQuotaResp{
		SecKillUserQuota: convert.ModelSecKillUserQuotaConvertPb(quota),
	}, nil
}
