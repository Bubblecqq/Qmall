package convert

import (
	"QMall/marketing/cmd/domain/model"
	"QMall/marketing/cmd/rpc/pb"
)

func ModelActivityInfoConvertPb(activity *model.ActivityInfo) *pb.ActivityInfo {
	return &pb.ActivityInfo{
		ActivityName:         activity.ActivityName,
		ProductName:          activity.ProductName,
		ProductStartingPrice: activity.ProductStartingPrice,
		CategoryId:           activity.CategoryId,
		ProductId:            activity.ProductId,
		ActivityTimeId:       activity.ActivityTimeId,
		StartTime:            activity.StartTime,
		EndTime:              activity.EndTime,
		IsOnline:             activity.IsOnline,
		ProductsNum:          activity.ProductsNum,
	}
}

func PbActivityInfoConvertModel(activity *pb.ActivityInfo) *model.ActivityInfo {
	return &model.ActivityInfo{
		ActivityName:         activity.ActivityName,
		ProductName:          activity.ProductName,
		ProductStartingPrice: activity.ProductStartingPrice,
		CategoryId:           activity.CategoryId,
		ProductId:            activity.ProductId,
		ActivityTimeId:       activity.ActivityTimeId,
		StartTime:            activity.StartTime,
		EndTime:              activity.EndTime,
		IsOnline:             activity.IsOnline,
		ProductsNum:          activity.ProductsNum,
	}
}
