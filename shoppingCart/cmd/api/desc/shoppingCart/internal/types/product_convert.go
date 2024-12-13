package types

import "QMall/product/cmd/domain/model"

func ProductModelConvertTypes(product *model.Product) Product {
	return Product{
		Id:                product.Id,
		CreateTime:        product.CreateTime.String(),
		UpdateTime:        product.UpdateTime.String(),
		IsEnabled:         product.IsEnabled == 1,
		CreateUser:        product.CreateUser,
		UpdateUser:        product.UpdateUser,
		Name:              product.Name,
		ProductType:       product.ProductType,
		CategoryId:        product.CategoryId,
		StartingPrice:     product.StartingPrice,
		TotalStock:        product.TotalStock,
		MainPicture:       product.MainPicture,
		RemoteAreaPostage: product.RemoteAreaPostage,
		SingleBuyLimit:    product.SingleBuyLimit,
		Remark:            product.Remark,
	}
}

func ProductSkuModelConvertTypes(productSku *model.ProductSku) ProductSku {
	return ProductSku{
		Id:                  productSku.Id,
		CreateTime:          productSku.CreateTime.String(),
		UpdateTime:          productSku.UpdateTime.String(),
		IsEnabled:           productSku.IsEnabled,
		CreateUser:          productSku.CreateUser,
		UpdateUser:          productSku.UpdateUser,
		Name:                productSku.Name,
		ProductId:           productSku.ProductId,
		AttributeSymbolList: productSku.AttributeSymbolList,
		SellPrice:           productSku.SellPrice,
		CostPrice:           productSku.CostPrice,
		Stock:               productSku.Stock,
		StockWarn:           productSku.StockWarn,
	}
}
