package logic

import (
	"QMall/product/cmd/rpc/product/internal/svc"
	"QMall/product/cmd/rpc/product/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveProductSkuWithCacheLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveProductSkuWithCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveProductSkuWithCacheLogic {
	return &SaveProductSkuWithCacheLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveProductSkuWithCacheLogic) SaveProductSkuWithCache(in *pb.SaveProductSkuWithCacheReq) (*pb.SaveProductSkuWithCacheResp, error) {

	//cacheKey:= fmt.Sprintf("activity_info:%d:%s", in.SkuId,)

	return &pb.SaveProductSkuWithCacheResp{}, nil
}
