package convert

import (
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelSecKillProductsConvertPb(data *model.SecKillProducts) *pb.SecKillProducts {
	return &pb.SecKillProducts{
		Id:           data.Id,
		Seller:       data.Seller,
		Price:        data.Price,
		ProductsName: data.ProductsName,
		ProductsNum:  data.ProductsNum,
		PictureUrl:   data.PictureUrl,
		CreateTime:   timestamppb.New(data.CreateTime),
	}
}

func PbSecKillProductsConvertModel(data *pb.SecKillProducts) *model.SecKillProducts {
	return &model.SecKillProducts{
		Id:           data.Id,
		Seller:       data.Seller,
		Price:        data.Price,
		ProductsName: data.ProductsName,
		ProductsNum:  data.ProductsNum,
		PictureUrl:   data.PictureUrl,
		CreateTime:   data.CreateTime.AsTime(),
	}
}
