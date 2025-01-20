package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"
	"fmt"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecKillUserQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSecKillUserQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSecKillUserQuotaLogic {
	return &GetSecKillUserQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSecKillUserQuotaLogic) GetSecKillUserQuota(in *pb.GetSecKillUserQuotaReq) (*pb.GetSecKillUserQuotaResp, error) {

	fmt.Printf("正在查询当前用户是否存在限额>>>>>当前请求的用户Id：%v，商品Id：%v\n", in.UserId, in.ProductsId)

	userQuota, err := l.svcCtx.SecKillRepository.GetSecKillUserQuota(in.UserId, in.ProductsId)

	if err != nil {
		return &pb.GetSecKillUserQuotaResp{}, err
	}
	fmt.Println("已经查询到当前用户限额>>>>>", userQuota)
	return &pb.GetSecKillUserQuotaResp{
		SecKillUserQuota: convert.ModelSecKillUserQuotaConvertPb(userQuota),
	}, nil
}
