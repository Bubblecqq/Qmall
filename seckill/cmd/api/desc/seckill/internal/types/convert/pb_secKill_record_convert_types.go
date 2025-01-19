package convert

import (
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"QMall/seckill/cmd/rpc/pb"
)

func PbSecKillRecordConvertTypes(secKillRecord *pb.SecKillRecord) types.SecKillRecord {
	return types.SecKillRecord{
		Id:         secKillRecord.Id,
		UserId:     secKillRecord.UserId,
		ProductsId: secKillRecord.ProductsId,
		SecNum:     secKillRecord.SecNum,
		OrderNum:   secKillRecord.OrderNum,
		Price:      secKillRecord.Price,
		Status:     secKillRecord.Status,
		CreateTime: secKillRecord.CreateTime.AsTime().String(),
	}
}
