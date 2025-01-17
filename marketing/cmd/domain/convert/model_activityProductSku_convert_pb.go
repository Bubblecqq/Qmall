package convert

import (
	"QMall/marketing/cmd/domain/model"
	"QMall/marketing/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ModelActivityProductSkuConvertPb(activity *model.ActivityProductSku) *pb.ActivityProductSku {
	return &pb.ActivityProductSku{
		Id:                activity.ID,
		ActivityProductId: activity.ActivityProductID,
		ProductId:         activity.ProductID,
		ProductSkuId:      activity.ProductSkuID,
		Price:             activity.Price,
		OriginalPrice:     activity.OriginalPrice,
		Number:            activity.Number,
		Stock:             activity.Stock,
		CreateUser:        activity.CreateUser,
		CreateTime:        timestamppb.New(activity.CreateTime),
		UpdateUser:        activity.UpdateUser,
		//UpdateTime: func(update *time.Time) *timestamppb.Timestamp {
		//	return timestamppb.New(*activity.UpdateTime)
		//}(activity.UpdateTime),
		UpdateTime: timestamppb.New(*activity.UpdateTime),
		IsDeleted:  activity.IsDeleted,
	}
}

func PbActivityProductSkuConvertModel(activity *pb.ActivityProductSku) *model.ActivityProductSku {
	return &model.ActivityProductSku{
		ID:                activity.Id,
		ActivityProductID: activity.ActivityProductId,
		ProductID:         activity.ProductId,
		ProductSkuID:      activity.ProductSkuId,
		Price:             activity.Price,
		OriginalPrice:     activity.OriginalPrice,
		Number:            activity.Number,
		Stock:             activity.Stock,
		CreateUser:        activity.CreateUser,
		CreateTime:        activity.CreateTime.AsTime(),
		UpdateUser:        activity.UpdateUser,
		//UpdateTime: func(update *time.Time) *timestamppb.Timestamp {
		//	return timestamppb.New(*activity.UpdateTime)
		//}(activity.UpdateTime),
		UpdateTime: func(updateTime *timestamppb.Timestamp) *time.Time {
			curUpdateTime := activity.UpdateTime.AsTime()
			return &curUpdateTime
		}(activity.UpdateTime),
		IsDeleted: activity.IsDeleted,
	}
}
