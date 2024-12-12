package model

import (
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ShppingCartModelConvertPb(shoppingCart *ShoppingCart) *pb.ShoppingCart {
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
		UpdateTime:         timestamppb.New(*shoppingCart.UpdateTime),
	}
}
