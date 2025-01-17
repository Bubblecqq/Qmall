package activity

import (
	"QMall/marketing/cmd/api/desc/marketing/internal/types/convert"
	"QMall/marketing/cmd/rpc/activity"
	"context"
	"fmt"

	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewAddActivityProductLogic 添加活动商品信息
func NewAddActivityProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityProductLogic {
	return &AddActivityProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddActivityProductLogic) AddActivityProduct(req *types.AddActivityProductReq) (resp *types.AddActivityProductResp, err error) {

	product, err := l.svcCtx.ActivityRpcConf.AddActivityProduct(l.ctx, &activity.AddActivityProductReq{
		ActivityTimeId: req.ActivityTimeId,
		ProductId:      req.ProductId,
	})
	if err != nil {
		fmt.Printf(fmt.Errorf("添加活动商品信息失败！原因见：%v", err).Error())
		return
	}
	resp = new(types.AddActivityProductResp)
	resp.ActivityProduct = convert.PbActivityProductConvertTypes(product.ActivityProduct)
	return
}
