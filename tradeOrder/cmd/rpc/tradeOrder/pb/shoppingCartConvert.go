package pb

import "QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

func ShoppingCartsConvert(carts *pb.ShoppingCart) *ShoppingCartSimple {
	return &ShoppingCartSimple{
		Id:                 carts.Id,
		UserId:             carts.UserId,
		ProductId:          carts.ProductId,
		ProductSkuId:       carts.ProductSkuId,
		ProductName:        carts.ProductName,
		ProductMainPicture: carts.ProductMainPicture,
		Number:             carts.Number,
		CreateTime:         carts.CreateTime,
		UpdateTime:         carts.UpdateTime,
		CreateUser:         carts.CreateUser,
		UpdateUser:         carts.UpdateUser,
		IsDelete:           carts.IsDelete,
	}
}
