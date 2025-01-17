package activity

import (
	"context"

	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加活动商品信息
func NewAddActivityProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityProductLogic {
	return &AddActivityProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddActivityProductLogic) AddActivityProduct(req *types.AddActivityProductReq) (resp *types.AddActivityProductResp, err error) {
	// todo: add your logic here and delete this line

	return
}
