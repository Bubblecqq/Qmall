package logic

import (
	"QMall/marketing/cmd/domain/convert"
	"context"
	"fmt"

	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityInfoByProductsNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityInfoByProductsNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityInfoByProductsNumLogic {
	return &GetActivityInfoByProductsNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityInfoByProductsNumLogic) GetActivityInfoByProductsNum(in *pb.GetActivityInfoByProductsNumReq) (*pb.GetActivityInfoByProductsNumResp, error) {

	fmt.Printf("正在查询活动商品信息>>>>>当前请求商品Id：%v，请求商品编号：%v\n", in.ProductsId, in.ProductsNum)

	num, err := l.svcCtx.ActivityRepository.GetActivityInfoByProductsNum(in.ProductsId, in.ProductsNum)

	if err != nil {
		return &pb.GetActivityInfoByProductsNumResp{}, err
	}
	fmt.Printf("查询活动商品成功！活动商品信息：%v\n", num)
	return &pb.GetActivityInfoByProductsNumResp{
		ActivityInfo: convert.ModelActivityInfoConvertPb(num),
	}, nil
}
