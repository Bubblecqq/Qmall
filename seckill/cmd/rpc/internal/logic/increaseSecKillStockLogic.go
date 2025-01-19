package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"
	"fmt"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IncreaseSecKillStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncreaseSecKillStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncreaseSecKillStockLogic {
	return &IncreaseSecKillStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IncreaseSecKillStockLogic) IncreaseSecKillStock(in *pb.IncreaseSecKillStockReq) (*pb.IncreaseSecKillStockResp, error) {
	fmt.Printf("[IncreaseSecKillStockLogic] 正在调用持久化数据层.....\n")
	stock, err := l.svcCtx.SecKillRepository.IncreaseSecKillStock(in)
	if err != nil {
		return &pb.IncreaseSecKillStockResp{}, err
	}
	fmt.Printf("[IncreaseSecKillStockLogic] 秒杀库存信息添加完成！请求的商品Id：%v，库存数：%v\n", in.ProductsId, in.Stock)
	return &pb.IncreaseSecKillStockResp{
		SecKillStock: convert.ModelSecKillStockConvertPb(stock),
	}, nil
}
