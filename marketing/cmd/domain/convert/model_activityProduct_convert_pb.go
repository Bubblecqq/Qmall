package convert

import (
	"QMall/marketing/cmd/domain/model"
	"QMall/marketing/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ModelActivityProductConvertPb(activity *model.ActivityProduct) *pb.ActivityProduct {
	return &pb.ActivityProduct{
		Id:                   activity.ID,
		ActivityTimeId:       activity.ActivityTimeID,
		ProductId:            activity.ProductID,
		ProductName:          activity.ProductName,
		ProductMainPicture:   activity.ProductMainPic,
		ProductStartingPrice: activity.ProductStartingPrice,
		CategoryId:           activity.CategoryID,
		CreateUser:           activity.CreateUser,
		CreateTime:           timestamppb.New(activity.CreateTime),
		UpdateUser:           activity.UpdateUser,
		//UpdateTime: func(update *time.Time) *timestamppb.Timestamp {
		//	return timestamppb.New(*activity.UpdateTime)
		//}(activity.UpdateTime),
		//UpdateTime: timestamppb.New(*activity.UpdateTime),
		IsDeleted:   activity.IsDeleted,
		ProductsNum: activity.ProductsNum,
	}
}

func PbActivityProductConvertModel(activity *pb.ActivityProduct) *model.ActivityProduct {
	return &model.ActivityProduct{
		ID:                   activity.Id,
		ActivityTimeID:       activity.ActivityTimeId,
		ProductID:            activity.ProductId,
		ProductName:          activity.ProductName,
		ProductMainPic:       activity.ProductMainPicture,
		ProductStartingPrice: activity.ProductStartingPrice,
		CategoryID:           activity.CategoryId,
		CreateUser:           activity.CreateUser,
		CreateTime:           activity.CreateTime.AsTime(),
		UpdateUser:           activity.UpdateUser,
		//UpdateTime: func(update *time.Time) *timestamppb.Timestamp {
		//	return timestamppb.New(*activity.UpdateTime)
		//}(activity.UpdateTime),
		UpdateTime: func(updateTime *timestamppb.Timestamp) *time.Time {
			curUpdateTime := activity.UpdateTime.AsTime()
			return &curUpdateTime
		}(activity.UpdateTime),
		IsDeleted:   activity.IsDeleted,
		ProductsNum: activity.ProductsNum,
	}
}
