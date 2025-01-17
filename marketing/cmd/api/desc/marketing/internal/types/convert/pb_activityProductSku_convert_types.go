package convert

import (
	"QMall/marketing/cmd/api/desc/marketing/internal/types"
	"QMall/marketing/cmd/rpc/pb"
)

func PbActivityProductSkuConvertTypes(activity *pb.ActivityProductSku) types.ActivityProductSku {
	return types.ActivityProductSku{
		ID:                activity.Id,
		ActivityProductID: activity.ActivityProductId,
		ProductID:         activity.ProductId,
		ProductSkuID:      activity.ProductSkuId,
		Price:             activity.Price,
		OriginalPrice:     activity.OriginalPrice,
		Number:            activity.Number,
		Stock:             activity.Stock,
		CreateUser:        activity.CreateUser,
		CreateTime:        activity.CreateTime.AsTime().String(),
		UpdateUser:        activity.UpdateUser,
		//UpdateTime: func(update *time.Time) *timestamppb.Timestamp {
		//	return timestamppb.New(*activity.UpdateTime)
		//}(activity.UpdateTime),
		UpdateTime: activity.UpdateTime.AsTime().String(),
		IsDeleted:  activity.IsDeleted,
	}
}
