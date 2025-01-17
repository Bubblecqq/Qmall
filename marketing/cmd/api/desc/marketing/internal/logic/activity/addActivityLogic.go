package activity

import (
	"QMall/marketing/cmd/api/desc/marketing/internal/types/convert"
	"QMall/marketing/cmd/rpc/activity"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewAddActivityLogic 添加活动表
func NewAddActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityLogic {
	return &AddActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddActivityLogic) AddActivity(req *types.AddActivityReq) (resp *types.AddActivityResp, err error) {
	startTime, err := time.Parse(time.RFC3339, req.ActivityStartTime)
	if err != nil {
		fmt.Printf(fmt.Errorf("invaild start time format：%v", err).Error())
		return
	}
	endTime, err := time.Parse(time.RFC3339, req.ActivityEndTime)
	if err != nil {
		fmt.Printf(fmt.Errorf("invaild end time format：%v", err).Error())
		return
	}
	addActivity, err := l.svcCtx.ActivityRpcConf.AddActivity(l.ctx, &activity.AddActivityReq{
		ActivityName:      req.ActivityName,
		IsOnline:          req.IsOnline,
		ActivityStartTime: timestamppb.New(startTime),
		ActivityEndTime:   timestamppb.New(endTime),
	})

	resp = new(types.AddActivityResp)
	resp.Activity = convert.PbActivityConvertTypes(addActivity.Activity)
	return
}
