package convert

import (
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelSecKillStockConvertPb(data *model.SecKillStock) *pb.SecKillStock {
	return &pb.SecKillStock{
		Id:         data.Id,
		ProductsId: data.ProductsId,
		Stock:      data.Stock,
		CreateTime: timestamppb.New(data.CreateTime),
	}
}

func PbSecKillStockConvertModel(data *pb.SecKillStock) *model.SecKillStock {
	return &model.SecKillStock{
		Id:         data.Id,
		ProductsId: data.ProductsId,
		Stock:      data.Stock,
		CreateTime: data.CreateTime.AsTime(),
	}
}
