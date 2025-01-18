package convert

import (
	"QMall/marketing/cmd/domain/model"
	"QMall/marketing/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ModelActivityConvertPb(activity *model.Activity) *pb.Activity {
	return &pb.Activity{
		Id:                activity.ID,
		Name:              activity.Name,
		ActivityStartTime: timestamppb.New(activity.ActivityStartTime),
		ActivityEndTime:   timestamppb.New(activity.ActivityEndTime),
		IsOnline:          activity.IsOnline,
		CreateUser:        activity.CreateUser,
		CreateTime:        timestamppb.New(activity.CreateTime),
		UpdateUser:        activity.UpdateUser,
		//UpdateTime: func(update *time.Time) *timestamppb.Timestamp {
		//	return timestamppb.New(*activity.UpdateTime)
		//}(activity.UpdateTime),
		//UpdateTime: timestamppb.New(*activity.UpdateTime),
		IsDeleted: activity.IsDeleted,
	}
}

func PbActivityConvertModel(activity *pb.Activity) *model.Activity {
	return &model.Activity{
		ID:                activity.Id,
		Name:              activity.Name,
		ActivityStartTime: activity.ActivityStartTime.AsTime(),
		ActivityEndTime:   activity.ActivityEndTime.AsTime(),
		IsOnline:          activity.IsOnline,
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
