package logic

import (
	"QMall/marketing/cmd/domain/convert"
	"QMall/marketing/cmd/domain/model"
	"context"
	"fmt"

	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityInfoByProductsNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityInfoByProductsNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityInfoByProductsNumLogic {
	return &GetActivityInfoByProductsNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityInfoByProductsNumLogic) GetActivityInfoByProductsNum(in *pb.GetActivityInfoByProductsNumReq) (*pb.GetActivityInfoByProductsNumResp, error) {

	// 生成缓存键
	cacheKey := fmt.Sprintf("activity_info:%d:%s", in.ProductsId, in.ProductsNum)
	fmt.Printf("正在查询活动商品信息>>>>>当前请求商品Id：%v，请求商品编号：%v\n", in.ProductsId, in.ProductsNum)
	var activityInfo *model.ActivityInfo
	activityInfo, err := l.svcCtx.ActivityRepository.GetActivityInfoWithCache(l.ctx, cacheKey)
	if err != nil {

		if err.Error() != "cache miss" {
			return &pb.GetActivityInfoByProductsNumResp{}, err
		}
		// 缓存未命中，从数据库查询
		fmt.Printf("当前缓存未命中，正在查询MYSQL....\n")
		activityInfo, err = l.svcCtx.ActivityRepository.GetActivityInfoByProductsNum(in.ProductsId, in.ProductsNum)

		if err != nil {
			return &pb.GetActivityInfoByProductsNumResp{}, err
		}
		// 将从数据库查询到的结果存入缓存
		err = l.svcCtx.ActivityRepository.SetActivityInfoWithCache(l.ctx, cacheKey, activityInfo)
		if err != nil {
			// 记录缓存设置失败的日志，但不影响返回结果
			fmt.Printf("设置活动商品信息缓存失败：%v\n", err)
		}
		fmt.Printf("设置活动商品信息缓存成功！\n")
		// 将库存存入缓存

	}
	//activityInfo, err := l.svcCtx.ActivityRepository.GetActivityInfoByProductsNum(in.ProductsId, in.ProductsNum)
	//
	//if err != nil {
	//	return &pb.GetActivityInfoByProductsNumResp{}, err
	//}
	fmt.Printf("查询活动商品成功！活动商品信息：%v\n", activityInfo)
	return &pb.GetActivityInfoByProductsNumResp{
		ActivityInfo: convert.ModelActivityInfoConvertPb(activityInfo),
	}, nil
}
