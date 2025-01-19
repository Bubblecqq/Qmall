package convert

import (
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelSecKillRecordConvertPb(data *model.SecKillRecord) *pb.SecKillRecord {
	return &pb.SecKillRecord{
		Id:         data.Id,
		ProductsId: data.ProductsId,
		SecNum:     data.SecNum,
		OrderNum:   data.OrderNum,
		Price:      data.Price,
		Status:     data.Status,
		CreateTime: timestamppb.New(data.CreateTime),
	}
}

func PbSecKillRecordConvertModel(data *pb.SecKillRecord) *model.SecKillRecord {
	return &model.SecKillRecord{
		Id:         data.Id,
		ProductsId: data.ProductsId,
		SecNum:     data.SecNum,
		OrderNum:   data.OrderNum,
		Price:      data.Price,
		Status:     data.Status,
		CreateTime: data.CreateTime.AsTime(),
	}
}
