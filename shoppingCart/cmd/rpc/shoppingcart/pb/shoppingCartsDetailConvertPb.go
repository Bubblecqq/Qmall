package pb

import "QMall/shoppingCart/cmd/domain/model"

func ShoppingCartsDetailConvertPb(shoppingCartsProduct model.ShoppingCartsProductInfo) *ShoppingCartsProductInfo {
	return &ShoppingCartsProductInfo{
		ProductId:          shoppingCartsProduct.ProductID,
		ProductName:        shoppingCartsProduct.ProductName,
		ProductSkuId:       shoppingCartsProduct.ProductSkuID,
		ProductMainPicture: shoppingCartsProduct.ProductMainPic,
		SellPrice:          shoppingCartsProduct.SellPrice,
		Quantity:           shoppingCartsProduct.Quantity,
		SkuDescribe:        shoppingCartsProduct.SkuDescribe,
	}
}
