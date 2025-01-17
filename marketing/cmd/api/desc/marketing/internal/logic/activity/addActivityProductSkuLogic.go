package activity

import (
	"context"

	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityProductSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加活动商品库存信息
func NewAddActivityProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityProductSkuLogic {
	return &AddActivityProductSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddActivityProductSkuLogic) AddActivityProductSku(req *types.AddActivityProductSkuReq) (resp *types.AddActivityProductSkuResp, err error) {
	// todo: add your logic here and delete this line

	return
}
