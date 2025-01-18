package convert

import (
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"QMall/seckill/cmd/rpc/pb"
	"fmt"
)

func PbSecKillOrderConvertTypes(secKillOrder *pb.SecKillOrder) *types.SecKillOrder {
	return &types.SecKillOrder{
		Id:          secKillOrder.Id,
		Seller:      secKillOrder.Seller,
		Buyer:       secKillOrder.Buyer,
		ProductsId:  secKillOrder.ProductsId,
		Price:       fmt.Sprintf("%v", secKillOrder.Price),
		ProductsNum: secKillOrder.ProductsNum,
		OrderNum:    secKillOrder.OrderNum,
		Status:      secKillOrder.Status,
		CreateTime:  secKillOrder.CreateTime.AsTime().String(),
	}
}
