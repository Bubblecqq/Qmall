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
	fmt.Printf("[IncreaseSecKillProductsLogic] 秒杀商品信息添加完成！商品信息Id：%v，秒杀价格：%v，卖家信息：%v\n", in.ProductId, in.Price, in.Seller)

	return &pb.IncreaseSecKillProductsResp{
		SecKillProducts: convert.ModelSecKillProductsConvertPb(products),
	}, nil
}
