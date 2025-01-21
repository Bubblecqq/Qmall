package convert

import (
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelSecKillQuotaConvertPb(data *model.SecKillQuota) *pb.SecKillQuota {
	return &pb.SecKillQuota{
		Id:                data.Id,
		ProductsId:        data.ProductsId,
		Num:               data.Num,
		CreateTime:        timestamppb.New(data.CreateTime),
		SeckillProductsId: data.SecKillProductsId,
	}
}

func PbSecKillQuotaConvertModel(data *pb.SecKillQuota) *model.SecKillQuota {
	return &model.SecKillQuota{
		Id:                data.Id,
		ProductsId:        data.ProductsId,
		Num:               data.Num,
		SecKillProductsId: data.SeckillProductsId,
		CreateTime:        data.CreateTime.AsTime(),
	}
}
