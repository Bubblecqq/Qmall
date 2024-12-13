package model

import (
	"QMall/product/cmd/domain/model"
	"QMall/product/cmd/rpc/product/pb"
)

func PbProductConvertModel(product *pb.Product) model.Product {
	return model.Product{
		BaseModel: model.BaseModel{
			Id:         product.Id,
			CreateTime: product.CreateTime.AsTime(),
			UpdateTime: product.UpdateTime.AsTime(),
			IsEnabled:  product.IsEnable,
			CreateUser: product.CreateUser,
			UpdateUser: product.UpdateUser,
		},
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
