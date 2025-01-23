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
	resp = new(types.AddActivityProductResp)
	if req.ActivityTimeId <= 0 || req.ProductId <= 0 {
		err = errors.New(fmt.Sprintf("当前传入的参数不合法，具体为活动时间表Id：%v，请求的商品Id：%v", req.ActivityTimeId, req.ProductId))
		return
	}
	l.Info("正在获取活动表信息>>>>>>请求的活动表Id信息：", req.ActivityTimeId)
	activityTimeById, err := l.svcCtx.ActivityRpcConf.GetActivityTimeById(l.ctx, &activity.GetActivityTimeByIdReq{
		Id: req.ActivityTimeId,
	})

	if err != nil {
		err = errors.New(fmt.Sprintf("当前请求的商品Id：%v,活动Id：%v暂未开启活动！", req.ProductId, req.ActivityTimeId))
		return
	}
	l.Info(fmt.Printf("成功获取到活动时间表：%v\n", activityTimeById.GetActivityTime()))
	isOnline := activityTimeById.ActivityTime.IsEnable == 1
	l.Info(fmt.Printf("正在添加活动商品信息>>>>,当前请求的活动表ID：%v，商品Id：%v，当前商品是否可上线：%v", req.ActivityTimeId, req.ProductId, isOnline))
	product, err := l.svcCtx.ActivityRpcConf.AddActivityProduct(l.ctx, &activity.AddActivityProductReq{
		ActivityTimeId: req.ActivityTimeId,
		ProductId:      req.ProductId,
		IsOnline:       isOnline,
	})
	if err != nil {
		fmt.Printf(fmt.Errorf("添加活动商品信息失败！原因见：%v", err).Error())
		return
	}

	resp.ActivityProduct = convert.PbActivityProductConvertTypes(product.ActivityProduct)
	return
}
