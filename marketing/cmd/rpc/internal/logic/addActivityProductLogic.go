package logic

import (
	"QMall/common"
	"QMall/marketing/cmd/domain/convert"
	"QMall/marketing/cmd/domain/model"
	product2 "QMall/product/cmd/rpc/product/product"
	"context"
	"fmt"
	"time"

	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddActivityProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityProductLogic {
	return &AddActivityProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddActivityProductLogic) AddActivityProduct(in *pb.AddActivityProductReq) (*pb.AddActivityProductResp, error) {
	fmt.Printf("正在添加活动商品>>>>>>>>商品Id：%v，活动Id：%v\n", in.ProductId, in.ActivityTimeId)
	activityProduct := &model.ActivityProduct{}
	getProduct, err := l.svcCtx.ProductRPC.GetProduct(l.ctx, &product2.GetProductByIdReq{
		Id: in.ProductId,
	})
	if err != nil {
		fmt.Printf("当前商品Id：%v不存在！原因见：%v\n", in.ProductId, err)
		return &pb.AddActivityProductResp{}, err
	}
	// 是否上线
	if !in.IsOnline {
		fmt.Printf("当前商品Id：%v暂未开启活动！原因见：%v\n", in.ProductId, err)
		return &pb.AddActivityProductResp{}, err
	}
	now := time.Now()
	activityProduct.ProductID = in.ProductId
	activityProduct.ActivityTimeID = in.ActivityTimeId
	activityProduct.IsDeleted = 0
	activityProduct.ProductName = getProduct.Product.Name
	activityProduct.ProductStartingPrice = getProduct.Product.StartingPrice
	activityProduct.ProductMainPic = getProduct.Product.MainPicture
	activityProduct.CategoryID = int64(getProduct.Product.CategoryId)
	activityProduct.CreateTime = now
	activityProduct.ProductsNum = common.GenerateProductsNum(now, in.ProductId)
	product, err := l.svcCtx.ActivityRepository.AddActivityProduct(activityProduct)
	if err != nil {
		return &pb.AddActivityProductResp{}, err
	}
	fmt.Printf("活动商品名%v已添加！>>>>>>>>\n", product.ProductName)
	return &pb.AddActivityProductResp{
		ActivityProduct: convert.ModelActivityProductConvertPb(product),
	}, nil
}
