package logic

import (
	"QMall/marketing/cmd/domain/convert"
	"context"
	"fmt"

	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityProductByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityProductByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityProductByIdLogic {
	return &GetActivityProductByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetActivityProductById 查询接口
func (l *GetActivityProductByIdLogic) GetActivityProductById(in *pb.GetActivityProductByIdReq) (*pb.GetActivityProductByIdResp, error) {

	fmt.Printf("[GetActivityProductByIdLogic] 正在查询活动商品>>>>>>>，请求的商品Id：%v\n", in.ProductId)

	activityProductsId, err := l.svcCtx.ActivityRepository.GetActivityProductByProductsId(in.ProductId)

	if err != nil {
		return &pb.GetActivityProductByIdResp{}, err
	}
	fmt.Printf("[GetActivityProductByIdLogic] 查询到活动商品记录！>>>>>>>，活动商品信息：%v\n", activityProductsId)

	return &pb.GetActivityProductByIdResp{
		ActivityProduct: convert.ModelActivityProductConvertPb(activityProductsId),
	}, nil
}
