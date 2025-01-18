package convert

import (
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"QMall/seckill/cmd/rpc/pb"
)

func PbSecKillQuotaConvertTypes(secKillQuota *pb.SecKillQuota) *types.SecKillQuota {
	return &types.SecKillQuota{
		Id:         secKillQuota.Id,
		ProductsId: secKillQuota.ProductsId,
		Num:        secKillQuota.Num,
		CreateTime: secKillQuota.CreateTime.AsTime().String(),
	}
}
