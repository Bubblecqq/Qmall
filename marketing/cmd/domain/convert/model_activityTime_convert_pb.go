package convert

import (
	"QMall/marketing/cmd/domain/model"
	"QMall/marketing/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelActivityTimeConvertPb(activity *model.ActivityTime) *pb.ActivityTime {
	return &pb.ActivityTime{
		Id:         activity.ID,
		ActivityId: activity.ActivityID,
		TimeName:   activity.TimeName,
		StartTime:  timestamppb.New(activity.StartTime),
		EndTime:    timestamppb.New(activity.EndTime),
		IsEnable:   activity.IsEnable,
		CreateUser: activity.CreateUser,
		CreateTime: timestamppb.New(activity.CreateTime),
		UpdateUser: activity.UpdateUser,
		//UpdateTime: func(update *time.Time) *timestamppb.Timestamp {
		//	return timestamppb.New(*activity.UpdateTime)
		//}(activity.UpdateTime),
		//UpdateTime: timestamppb.New(*activity.UpdateTime),
		IsDeleted: activity.IsDeleted,
	}
}

func PbActivityTimeConvertModel(activity *pb.ActivityTime) *model.ActivityTime {
	return &model.ActivityTime{
		ID:         activity.Id,
		ActivityID: activity.ActivityId,
		TimeName:   activity.TimeName,
		StartTime:  activity.StartTime.AsTime(),
		EndTime:    activity.EndTime.AsTime(),
		IsEnable:   activity.IsEnable,
		CreateUser: activity.CreateUser,
		CreateTime: activity.CreateTime.AsTime(),
		UpdateUser: activity.UpdateUser,
		//UpdateTime: func(updateTime *timestamppb.Timestamp) *time.Time {
		//	curUpdateTime := activity.UpdateTime.AsTime()
		//	return &curUpdateTime
		//}(activity.UpdateTime),
		//UpdateTime: activity.UpdateTime.AsTime(),
		IsDeleted: activity.IsDeleted,
	}
}
