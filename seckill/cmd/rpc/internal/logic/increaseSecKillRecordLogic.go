package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"
	"fmt"

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
	fmt.Printf("[IncreaseSecKillRecordLogic] 正在调用持久化数据层.....\n")

	record, err := l.svcCtx.SecKillRepository.IncreaseSecKillRecord(in)

	if err != nil {
		return &pb.IncreaseSecKillRecordResp{}, err
	}
	fmt.Printf("[IncreaseSecKillRecordLogic] 用户秒杀记录添加完成！请求的用户Id：%v，商品Id：%v，秒杀价格：%v\n", in.UserId, in.ProductsId, in.Price)

	return &pb.IncreaseSecKillRecordResp{
		SecKillRecord: convert.ModelSecKillRecordConvertPb(record),
	}, nil
}
