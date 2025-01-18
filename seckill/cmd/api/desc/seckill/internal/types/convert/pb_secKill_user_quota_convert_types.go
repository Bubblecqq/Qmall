package convert

import (
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"QMall/seckill/cmd/rpc/pb"
)

func PbSecKillUserQuotaConvertTypes(secKillUserQuota *pb.SecKillUserQuota) *types.SecKillUserQuota {
	return &types.SecKillUserQuota{
		Id:         secKillUserQuota.Id,
		ProductsId: secKillUserQuota.ProductsId,
		UserId:     secKillUserQuota.UserId,
		Num:        secKillUserQuota.Num,
		KilledNum:  secKillUserQuota.KilledNum,
		CreateTime: secKillUserQuota.CreateTime.AsTime().String(),
	}
}
