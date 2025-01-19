package logic

import (
	"context"

	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityTimeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityTimeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityTimeByIdLogic {
	return &GetActivityTimeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityTimeByIdLogic) GetActivityTimeById(in *pb.GetActivityTimeByIdReq) (*pb.GetActivityTimeByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetActivityTimeByIdResp{}, nil
}
