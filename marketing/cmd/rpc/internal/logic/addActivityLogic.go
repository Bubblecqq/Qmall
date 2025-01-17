package logic

import (
	"QMall/marketing/cmd/domain/convert"
	"context"
	"fmt"

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
	fmt.Printf("正在添加活动>>>>>>>>%v\n", in.ActivityName)
	activity, err := l.svcCtx.ActivityRepository.AddActivity(in.ActivityName, in.IsOnline, in.ActivityStartTime.AsTime(), in.ActivityEndTime.AsTime())

	if err != nil {
		return &pb.AddActivityResp{}, err
	}
	fmt.Printf("活动%v已添加！>>>>>>>>\n", in.ActivityName)
	return &pb.AddActivityResp{
		Activity: convert.ModelActivityConvertPb(activity),
	}, nil
}
