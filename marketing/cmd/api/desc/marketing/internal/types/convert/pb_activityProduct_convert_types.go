package convert

import (
	"QMall/marketing/cmd/api/desc/marketing/internal/types"
	"QMall/marketing/cmd/rpc/pb"
)

func PbActivityProductConvertTypes(activity *pb.ActivityProduct) types.ActivityProduct {
	return types.ActivityProduct{
		ID:                   activity.Id,
		ActivityTimeID:       activity.ActivityTimeId,
		ProductID:            activity.ProductId,
		ProductName:          activity.ProductName,
		ProductMainPic:       activity.ProductMainPicture,
		ProductStartingPrice: activity.ProductStartingPrice,
		CategoryID:           activity.CategoryId,
		CreateUser:           activity.CreateUser,
		CreateTime:           activity.CreateTime.AsTime().String(),
		UpdateUser:           activity.UpdateUser,
		//UpdateTime: func(update *time.Time) *timestamppb.Timestamp {
		//	return timestamppb.New(*activity.UpdateTime)
		//}(activity.UpdateTime),
		UpdateTime:  activity.UpdateTime.AsTime().String(),
		IsDeleted:   activity.IsDeleted,
		ProductsNum: activity.ProductsNum,
	}
}
