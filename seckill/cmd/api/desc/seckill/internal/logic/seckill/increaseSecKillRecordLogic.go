package seckill

import (
	"QMall/seckill/cmd/api/desc/seckill/internal/types/convert"
	"QMall/seckill/cmd/rpc/seckill"
	"context"
	"fmt"

	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewIncreaseSecKillRecordLogic 添加秒杀记录
func NewIncreaseSecKillRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillRecordLogic {
	return &IncreaseSecKillRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IncreaseSecKillRecordLogic) IncreaseSecKillRecord(req *types.IncreaseSecKillRecordReq) (resp *types.IncreaseSecKillRecordResp, err error) {
	l.Info(fmt.Printf("[*] 正在添加用户秒杀记录信息>>>>>>当前访问的用户Id：%v，商品Id：%v\n", req.UserId, req.ProductsId))

	record, err := l.svcCtx.SecKillRpcConf.IncreaseSecKillRecord(l.ctx, &seckill.IncreaseSecKillRecordReq{
		UserId:     req.UserId,
		ProductsId: req.ProductsId,
	})
	l.Info(fmt.Printf("[Info] 已在添加用户秒杀记录信息>>>>>>当前访问的用户Id：%v，商品Id：%v\n", req.UserId, req.ProductsId))

	resp = new(types.IncreaseSecKillRecordResp)
	resp.SecKillRecord = convert.PbSecKillRecordConvertTypes(record.SecKillRecord)
	return
}
