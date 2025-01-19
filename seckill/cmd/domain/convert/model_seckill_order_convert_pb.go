package convert

import (
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelSecKillOrderConvertPb(data *model.SecKillOrder) *pb.SecKillOrder {
	return &pb.SecKillOrder{
		Id:          data.Id,
		Seller:      data.Seller,
		Buyer:       data.Buyer,
		Price:       data.Price,
		ProductsId:  data.ProductsId,
		ProductsNum: data.ProductsNum,
		OrderNum:    data.OrderNum,
		Status:      data.Status,
		CreateTime:  timestamppb.New(data.CreateTime),
	}
}

func PbSecKillOrderConvertModel(data *pb.SecKillOrder) *model.SecKillOrder {
	return &model.SecKillOrder{
		Id:          data.Id,
		Seller:      data.Seller,
		Buyer:       data.Buyer,
		Price:       data.Price,
		ProductsId:  data.ProductsId,
		ProductsNum: data.ProductsNum,
		OrderNum:    data.OrderNum,
		Status:      data.Status,
		CreateTime:  data.CreateTime.AsTime(),
	}
}
