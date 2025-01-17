package logic

import (
	"QMall/marketing/cmd/domain/convert"
	"context"
	"fmt"

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
	fmt.Printf("正在添加活动时间>>>>>>>>活动Id：%v，活动名：%v\n", in.ActivityId, in.Name)

	time, err := l.svcCtx.ActivityRepository.AddActivityTime(in.ActivityId, in.Name, 1)

	if err != nil {
		fmt.Printf("当前活动时间Id：%v不存在！原因见：%v\n", time.ID, err)
		return &pb.AddActivityTimeResp{}, err
	}
	activityById, err := l.svcCtx.ActivityRepository.GetActivityById(time.ActivityID)
	if err != nil {
		fmt.Printf("当前活动Id：%v不存在！原因见：%v\n", activityById.ID, err)
		return &pb.AddActivityTimeResp{}, err
	}
	fmt.Printf("当前活动Id：%v已添加！>>>>>>>>\n", activityById.ID)

	return &pb.AddActivityTimeResp{
		ActivityId: in.ActivityId,
		Name:       in.Name,
		Activity:   convert.ModelActivityConvertPb(activityById),
	}, nil
}
