package activity

import (
	"QMall/marketing/cmd/rpc/activity"
	"context"
	"fmt"

	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewAddActivityTimeLogic 添加活动表
func NewAddActivityTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityTimeLogic {
	return &AddActivityTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddActivityTimeLogic) AddActivityTime(req *types.AddActivityTimeReq) (resp *types.AddActivityTimeResp, err error) {

	time, err := l.svcCtx.ActivityRpcConf.AddActivityTime(l.ctx, &activity.AddActivityTimeReq{
		ActivityId: req.ActivityId,
		Name:       req.Name,
	})
	if err != nil {
		fmt.Printf(fmt.Errorf("添加活动商品时间信息失败！原因见：%v", err).Error())
		return
	}
	resp = new(types.AddActivityTimeResp)
	resp.ActivityId = time.ActivityId
	resp.Name = time.Name

	return
}
