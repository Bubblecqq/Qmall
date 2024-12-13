package types

import (
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"
)

func ShoppingCartPbConvertTypes(shoppingCart *pb.ShoppingCart) ShoppingCart {
	return ShoppingCart{
		Id:                 shoppingCart.Id,
		UserId:             shoppingCart.UserId,
		ProductId:          shoppingCart.ProductId,
		ProductSkuId:       shoppingCart.ProductSkuId,
		ProductMainPicture: shoppingCart.ProductMainPicture,
		ProductName:        shoppingCart.ProductName,
		Number:             shoppingCart.Number,
		CreateUser:         shoppingCart.CreateUser,
		UpdateUser:         shoppingCart.UpdateUser,
		CreateTime:         shoppingCart.CreateTime.AsTime().String(),
		UpdateTime:         shoppingCart.UpdateTime.AsTime().String(),
	}
}
