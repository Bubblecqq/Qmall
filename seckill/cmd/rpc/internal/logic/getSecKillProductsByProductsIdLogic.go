package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"
	"fmt"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecKillProductsByProductsIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSecKillProductsByProductsIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSecKillProductsByProductsIdLogic {
	return &GetSecKillProductsByProductsIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSecKillProductsByProductsIdLogic) GetSecKillProductsByProductsId(in *pb.GetSecKillProductsByProductsIdReq) (*pb.GetSecKillProductsByProductsIdResp, error) {

	fmt.Printf("正在查询秒杀商品信息.......\n")
	productsId, err := l.svcCtx.SecKillRepository.GetSecKillProductsByProductsId(in.ProductId)

	if err != nil {
		return &pb.GetSecKillProductsByProductsIdResp{}, err
	}

	fmt.Printf("已经查询到秒杀商品！商品信息如下：%v\n", productsId)

	return &pb.GetSecKillProductsByProductsIdResp{
		SecKillProducts: convert.ModelSecKillProductsConvertPb(productsId),
	}, nil
}
