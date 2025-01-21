package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecKillQuotaByProductsIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSecKillQuotaByProductsIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSecKillQuotaByProductsIdLogic {
	return &GetSecKillQuotaByProductsIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSecKillQuotaByProductsIdLogic) GetSecKillQuotaByProductsId(in *pb.GetSecKillQuotaByProductsIdReq) (*pb.GetSecKillQuotaByProductsIdResp, error) {

	quotaByProductId, err := l.svcCtx.SecKillRepository.GetSecKillQuotaByProductsId(in.ProductId)

	if err != nil {
		return &pb.GetSecKillQuotaByProductsIdResp{}, err
	}

	return &pb.GetSecKillQuotaByProductsIdResp{
		SecKillQuota: convert.ModelSecKillQuotaConvertPb(quotaByProductId),
	}, nil
}
