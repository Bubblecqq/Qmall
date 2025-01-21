package logic

import (
	"QMall/marketing/cmd/domain/convert"
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
	id, err := l.svcCtx.ActivityRepository.GetActivityTimeById(in)

	if err != nil {
		return &pb.GetActivityTimeByIdResp{}, err
	}
	l.Info("成功提取活动表Id：", id)
	return &pb.GetActivityTimeByIdResp{
		ActivityTime: convert.ModelActivityTimeConvertPb(id),
	}, nil
}
