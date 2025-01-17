package convert

import (
	"QMall/marketing/cmd/api/desc/marketing/internal/types"
	"QMall/marketing/cmd/rpc/pb"
)

func PbActivityTimeConvertTypes(activity *pb.ActivityTime) types.ActivityTime {
	return types.ActivityTime{
		ID:         activity.Id,
		ActivityID: activity.ActivityId,
		TimeName:   activity.TimeName,
		StartTime:  activity.StartTime.AsTime().String(),
		EndTime:    activity.EndTime.AsTime().String(),
		IsEnable:   activity.IsEnable,
		CreateUser: activity.CreateUser,
		CreateTime: activity.CreateTime.AsTime().String(),
		UpdateUser: activity.UpdateUser,
		UpdateTime: activity.UpdateTime.AsTime().String(),
		IsDeleted:  activity.IsDeleted,
	}
}
