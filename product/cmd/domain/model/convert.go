package model

import (
	"goZeroDemo4/product/cmd/rpc/product/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PbProductModelConvert(product *pb.Product) *Product {
	return &Product{
		BaseModel: BaseModel{
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

func ProductModelConvertPbProduct(product *Product) *pb.Product {
	return &pb.Product{
		Id:                product.Id,
		CreateTime:        timestamppb.New(product.CreateTime),
		UpdateTime:        timestamppb.New(product.UpdateTime),
		IsEnable:          product.IsEnabled,
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

func PbProductSkuModelConvert(productSku *pb.ProductSku) *ProductSku {
	return &ProductSku{
		BaseModel: BaseModel{
			Id:         productSku.Id,
			CreateTime: productSku.CreateTime.AsTime(),
			UpdateTime: productSku.UpdateTime.AsTime(),
			IsEnabled:  productSku.IsEnable,
			CreateUser: productSku.CreateUser,
			UpdateUser: productSku.UpdateUser,
		},
		Name:                productSku.Name,
		ProductId:           productSku.ProductId,
		AttributeSymbolList: productSku.AttributeSymbolList,
		SellPrice:           productSku.SellPrice,
		CostPrice:           productSku.CostPrice,
		Stock:               productSku.Stock,
		StockWarn:           productSku.StockWarn,
	}
}

func ProductSkuModelConvertPbProductSku(productSku *ProductSku) *pb.ProductSku {
	return &pb.ProductSku{
		Id:                  productSku.Id,
		CreateTime:          timestamppb.New(productSku.CreateTime),
		UpdateTime:          timestamppb.New(productSku.UpdateTime),
		IsEnable:            productSku.IsEnabled,
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
