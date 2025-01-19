package convert

import (
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelSecKillUserQuotaConvertPb(data *model.SecKillUserQuota) *pb.SecKillUserQuota {
	return &pb.SecKillUserQuota{
		Id:         data.Id,
		ProductsId: data.ProductsId,
		UserId:     data.UserId,
		Num:        data.Num,
		KilledNum:  data.KilledNum,
		CreateTime: timestamppb.New(data.CreateTime),
	}
}

func PbSecKillUserQuotaConvertModel(data *pb.SecKillUserQuota) *model.SecKillUserQuota {
	return &model.SecKillUserQuota{
		Id:         data.Id,
		ProductsId: data.ProductsId,
		UserId:     data.UserId,
		Num:        data.Num,
		KilledNum:  data.KilledNum,
		CreateTime: data.CreateTime.AsTime(),
	}
}
