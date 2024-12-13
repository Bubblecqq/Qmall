package model

import (
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ShppingCartModelConvertPb(shoppingCart *ShoppingCart) *pb.ShoppingCart {
	var updateTime *time.Time
	if shoppingCart.UpdateTime != nil {
		updateTime = shoppingCart.UpdateTime
	}
	return &pb.ShoppingCart{
		Id:                 shoppingCart.Id,
		UserId:             shoppingCart.UserId,
		ProductId:          shoppingCart.ProductId,
		ProductSkuId:       shoppingCart.ProductSkuId,
		ProductMainPicture: shoppingCart.ProductMainPicture,
		ProductName:        shoppingCart.ProductName,
		Number:             shoppingCart.Number,
		CreateUser:         shoppingCart.CreateUser,
		UpdateUser:         shoppingCart.UpdateUser,
		CreateTime:         timestamppb.New(shoppingCart.CreateTime),
		UpdateTime: func() *timestamppb.Timestamp {
			if updateTime == nil {
				return timestamppb.New(time.Time{})
			}
			return timestamppb.New(*updateTime)
		}(),
	}
}
