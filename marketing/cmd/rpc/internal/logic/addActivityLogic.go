package logic

import (
	"context"

	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityLogic {
	return &AddActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddActivityLogic) AddActivity(in *pb.AddActivityReq) (*pb.AddActivityResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddActivityResp{}, nil
}
