package logic

import (
	"context"

	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddActivityTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityTimeLogic {
	return &AddActivityTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddActivityTimeLogic) AddActivityTime(in *pb.AddActivityTimeReq) (*pb.AddActivityTimeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddActivityTimeResp{}, nil
}
