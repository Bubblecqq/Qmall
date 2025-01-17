package activity

import (
	"context"

	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加活动表
func NewAddActivityTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityTimeLogic {
	return &AddActivityTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddActivityTimeLogic) AddActivityTime(req *types.AddActivityTimeReq) (resp *types.AddActivityTimeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
