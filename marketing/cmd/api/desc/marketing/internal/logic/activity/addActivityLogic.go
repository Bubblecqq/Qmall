package activity

import (
	"context"

	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加活动表
func NewAddActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityLogic {
	return &AddActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddActivityLogic) AddActivity(req *types.AddActivityReq) (resp *types.AddActivityResp, err error) {
	// todo: add your logic here and delete this line

	return
}
