package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"
	"fmt"

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
	fmt.Printf("[IncreaseSecKillUserQuotaLogic] 正在调用持久化数据层.....\n")

	quota, err := l.svcCtx.SecKillRepository.IncreaseSecKillUserQuota(in)

	if err != nil {
		return &pb.IncreaseSecKillUserQuotaResp{}, err
	}
	fmt.Printf("[IncreaseSecKillUserQuotaLogic] 秒杀用户限额添加完成！请求的用户Id：%v，请求的商品Id：%v，限额秒杀数：%v\n", in.UserId, in.ProductsId, in.Num)

	return &pb.IncreaseSecKillUserQuotaResp{
		SecKillUserQuota: convert.ModelSecKillUserQuotaConvertPb(quota),
	}, nil
}
