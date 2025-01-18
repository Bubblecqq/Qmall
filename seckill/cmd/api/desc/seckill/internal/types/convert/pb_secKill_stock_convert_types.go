package convert

import (
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"QMall/seckill/cmd/rpc/pb"
)

func PbSecKillStockConvertTypes(secKillStock *pb.SecKillStock) *types.SecKillStock {
	return &types.SecKillStock{
		Id:         secKillStock.Id,
		ProductsId: secKillStock.ProductsId,
		Stock:      secKillStock.Stock,
		CreateTime: secKillStock.CreateTime.AsTime().String(),
	}
}
