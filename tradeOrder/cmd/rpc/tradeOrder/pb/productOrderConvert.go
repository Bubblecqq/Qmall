package pb

import (
	"QMall/tradeOrder/cmd/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelProductOrderConvertPb(productOrder *model.TradeOrderProduct) *ProductOrder {
	return &ProductOrder{
		Id:              productOrder.Id,
		ProductId:       productOrder.ProductId,
		ProductSkuId:    productOrder.ProductSkuId,
		UserId:          productOrder.UserId,
		ProductName:     productOrder.ProductName,
		ProductImageUrl: productOrder.ProductImageUrl,
		SkuDescribe:     productOrder.SkuDescribe,
		Quantity:        productOrder.Quantity,
		ProductPrice:    productOrder.ProductPrice,
		RealPrice:       productOrder.RealPrice,
		RealAmount:      productOrder.RealAmount,
		CreateUser:      productOrder.CreateUser,
		UpdateUser:      productOrder.UpdateUser,
		CreateTime:      timestamppb.New(productOrder.CreateTime),
		//UpdateTime:      timestamppb.New(*productOrder.UpdateTime),
		ActivityId:    productOrder.ActivityId,
		DetailStatus:  productOrder.DetailStatus,
		ActivityType:  productOrder.ActivityType,
		CommentStatus: productOrder.CommentStatus,
	}
}

func PbProductOrderConvertModel(productOrder *ProductOrder) *model.TradeOrderProduct {
	return &model.TradeOrderProduct{
		Id:              productOrder.Id,
		ProductId:       productOrder.ProductId,
		ProductSkuId:    productOrder.ProductSkuId,
		UserId:          productOrder.UserId,
		ProductName:     productOrder.ProductName,
		ProductImageUrl: productOrder.ProductImageUrl,
		SkuDescribe:     productOrder.SkuDescribe,
		Quantity:        productOrder.Quantity,
		ProductPrice:    productOrder.ProductPrice,
		RealPrice:       productOrder.RealPrice,
		RealAmount:      productOrder.RealAmount,
		CreateUser:      productOrder.CreateUser,
		UpdateUser:      productOrder.UpdateUser,
		CreateTime:      productOrder.CreateTime.AsTime(),
		//UpdateTime:      timestamppb.New(*productOrder.UpdateTime),
		ActivityId:    productOrder.ActivityId,
		DetailStatus:  productOrder.DetailStatus,
		ActivityType:  productOrder.ActivityType,
		CommentStatus: productOrder.CommentStatus,
	}
}
