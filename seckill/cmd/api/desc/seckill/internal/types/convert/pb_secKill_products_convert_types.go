package convert

import (
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"QMall/seckill/cmd/rpc/pb"
)

func PbSecKillProductsConvertTypes(secKillProducts *pb.SecKillProducts) types.SecKillProducts {
	return types.SecKillProducts{
		Id:           secKillProducts.Id,
		Seller:       secKillProducts.Seller,
		Price:        secKillProducts.Price,
		ProductsNum:  secKillProducts.ProductsNum,
		ProductsName: secKillProducts.ProductsName,
		PictureUrl:   secKillProducts.PictureUrl,
		CreateTime:   secKillProducts.CreateTime.AsTime().String(),
	}
}
