package model

import "QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

func ShoppingCartsDetailConvertPb(shoppingCartsProduct ShoppingCartsProductInfo) *pb.ShoppingCartsProductInfo {
	return &pb.ShoppingCartsProductInfo{
		ProductId:          shoppingCartsProduct.ProductID,
		ProductName:        shoppingCartsProduct.ProductName,
		ProductSkuId:       shoppingCartsProduct.ProductSkuID,
		ProductMainPicture: shoppingCartsProduct.ProductMainPic,
		SellPrice:          shoppingCartsProduct.SellPrice,
		Quantity:           shoppingCartsProduct.Quantity,
		SkuDescribe:        shoppingCartsProduct.SkuDescribe,
	}
}
