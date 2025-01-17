package convert

import (
	"QMall/marketing/cmd/api/desc/marketing/internal/types"
	"QMall/marketing/cmd/rpc/pb"
)

func PbActivityConvertTypes(activity *pb.Activity) types.Activity {
	return types.Activity{
		ID:                activity.Id,
		Name:              activity.Name,
		ActivityStartTime: activity.ActivityStartTime.AsTime().String(),
		ActivityEndTime:   activity.ActivityEndTime.AsTime().String(),
		IsOnline:          activity.IsOnline,
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
