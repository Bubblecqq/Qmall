package logic

import (
	"QMall/seckill/cmd/domain/convert"
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSecKillUserQuotaByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSecKillUserQuotaByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSecKillUserQuotaByIdLogic {
	return &UpdateSecKillUserQuotaByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSecKillUserQuotaByIdLogic) UpdateSecKillUserQuotaById(in *pb.UpdateSecKillUserQuotaByIdReq) (*pb.UpdateSecKillUserQuotaByIdResp, error) {

	id, err := l.svcCtx.SecKillRepository.UpdateSecKillUserQuotaById(in)

	if err != nil {
		return &pb.UpdateSecKillUserQuotaByIdResp{}, err
	}

	return &pb.UpdateSecKillUserQuotaByIdResp{
		SecKillUserQuota: convert.ModelSecKillUserQuotaConvertPb(id),
	}, nil
}
