package logic

import (
	"QMall/tradeOrder/cmd/domain/model"
	"context"
	"time"

	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductOrderLogic {
	return &AddProductOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProductOrderLogic) AddProductOrder(in *pb.AddProductOrderReq) (*pb.AddProductOrderResp, error) {
	var productOrder *model.TradeOrderProduct
	productOrder.OrderId = in.OrderId
	productOrder.ProductId = in.ProductId
	productOrder.UserId = in.UserId
	productOrder.ProductSkuId = in.ProductSkuId
	productOrder.ProductName = in.ProductName
	productOrder.ProductImageUrl = in.ProductImageUrl
	productOrder.SkuDescribe = in.SkuDescribe
	productOrder.Quantity = in.Quantity
	productOrder.ProductPrice = in.ProductPrice
	productOrder.RealPrice = in.RealPrice
	productOrder.RealAmount = in.RealAmount
	productOrder.CreateUser = in.UserId
	productOrder.ActivityId = in.ActivityId
	productOrder.ActivityType = in.ActivityType
	productOrder.CommentStatus = in.CommentStatus
	productOrder.DetailStatus = in.DetailStatus
	productOrder.CreateTime = time.Now()
	_, err := l.svcCtx.TradeOrderProductRepository.AddOrderProduct(productOrder)
	return &pb.AddProductOrderResp{}, err
}
