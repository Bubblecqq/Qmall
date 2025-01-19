package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"
	"fmt"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseSecKillProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillProductsLogic {
	return &IncreaseSecKillProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseSecKillProductsLogic) IncreaseSecKillProducts(in *pb.IncreaseSecKillProductsReq) (*pb.IncreaseSecKillProductsResp, error) {

	fmt.Printf("[IncreaseSecKillProductsLogic] 正在调用持久化数据层.....\n")

	products, err := l.svcCtx.SecKillRepository.IncreaseSecKillProducts(in)

	if err != nil {
		return &pb.IncreaseSecKillProductsResp{}, err
	}

	return &pb.IncreaseSecKillProductsResp{
		SecKillProducts: convert.ModelSecKillProductsConvertPb(products),
	}, nil
}
