package activity

import (
	"QMall/marketing/cmd/api/desc/marketing/internal/types/convert"
	"QMall/marketing/cmd/rpc/activity"
	"context"
	"errors"
	"fmt"

	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityProductSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewAddActivityProductSkuLogic 添加活动商品库存信息
func NewAddActivityProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityProductSkuLogic {
	return &AddActivityProductSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddActivityProductSkuLogic) AddActivityProductSku(req *types.AddActivityProductSkuReq) (resp *types.AddActivityProductSkuResp, err error) {
	if req.ActivityProductId <= 0 || req.ProductId <= 0 || req.ProductSkuId <= 0 {
		err = errors.New(fmt.Sprintf("当前传入的参数不合法，具体为活动商品Id：%v，请求的商品Id：%v，请求的商品库存Id：%v", req.ActivityProductId, req.ProductId, req.ProductSkuId))
		return
	}
	sku, err := l.svcCtx.ActivityRpcConf.AddActivityProductSku(l.ctx, &activity.AddActivityProductSkuReq{
		ActivityProductId: req.ActivityProductId,
		ProductId:         req.ProductId,
		Price:             req.Price,
		Number:            req.Number,
		ProductSkuId:      req.ProductSkuId,
		UserId:            req.UserId,
	})
	if err != nil {
		fmt.Printf(fmt.Errorf("添加活动商品库存信息失败！原因见：%v", err).Error())
		return
	}
	resp = new(types.AddActivityProductSkuResp)
	resp.ActivityProductSku = convert.PbActivityProductSkuConvertTypes(sku.ActivityProductSku)
	return
}
