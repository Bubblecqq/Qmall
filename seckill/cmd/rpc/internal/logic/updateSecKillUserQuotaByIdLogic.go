package logic

import (
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
	// todo: add your logic here and delete this line

	return &pb.UpdateSecKillUserQuotaByIdResp{}, nil
}
