package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"
	"fmt"

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
	fmt.Printf("[IncreaseSecKillQuotaLogic] 正在调用持久化数据层.....\n")

	quota, err := l.svcCtx.SecKillRepository.IncreaseSecKillQuota(in)

	if err != nil {
		return &pb.IncreaseSecKillQuotaResp{}, err
	}
	fmt.Printf("[IncreaseSecKillQuotaLogic] 秒杀用户限额添加完成！请求的商品Id：%v，限额秒杀数：%v\n", in.ProductId, in.LimitNumber)

	return &pb.IncreaseSecKillQuotaResp{
		SecKillQuota: convert.ModelSecKillQuotaConvertPb(quota),
	}, nil
}
