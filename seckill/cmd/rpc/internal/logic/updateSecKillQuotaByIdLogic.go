package logic

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSecKillQuotaByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSecKillQuotaByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSecKillQuotaByIdLogic {
	return &UpdateSecKillQuotaByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSecKillQuotaById 更新接口
func (l *UpdateSecKillQuotaByIdLogic) UpdateSecKillQuotaById(in *pb.UpdateSecKillQuotaByIdReq) (*pb.UpdateSecKillQuotaByIdResp, error) {

	return &pb.UpdateSecKillQuotaByIdResp{}, nil
}
