package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"goZeroDemo4/common"
	"goZeroDemo4/user/cmd/rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func PbUserModelConvert(user *pb.User) *User {

	return &User{
		BaseModel: BaseModel{
			Id:         user.Id,
			IsDeleted:  user.IsDeleted,
			CreateTime: user.CreateTime.AsTime(),
			UpdateTime: user.UpdateTime.AsTime(),
		},
		Account:         user.Account,
		Nickname:        user.Nickname,
		Avatar:          user.Avatar,
		Gender:          user.Gender,
		Phone:           user.Phone,
		Password:        user.Password,
		IDCard:          user.IdCard,
		Source:          user.Source,
		SystemId:        user.SystemId,
		ClientId:        user.ClientId,
		UnionId:         user.UnionId,
		IsEnable:        user.IsEnable,
		LastLoginTime:   user.LastLoginTime.AsTime(),
		CreateUser:      user.CreateUser,
		UpdateUser:      user.UpdateUser,
		Token:           user.Token,
		SessionId:       user.SessionId,
		TokenExpireTime: user.TokenExpireTime,
	}
}

func UserModelConvertPbUser(userModel *User) *pb.User {
	//生成Token
	timeStr := fmt.Sprintf("%d", time.Now().Unix())
	tp, _ := time.ParseDuration("1h")
	tokenExpireTime := time.Now().Add(tp)
	expireTimeStr := tokenExpireTime.Format("2006-01-02 15:04:05")
	userPb := &pb.User{
		Id:              userModel.Id,
		CreateTime:      timestamppb.New(userModel.CreateTime), // 假设CreateTime是time.Time类型，格式化为字符串
		UpdateTime:      timestamppb.New(userModel.UpdateTime),
		Account:         userModel.Account,
		Nickname:        userModel.Nickname,
		Avatar:          userModel.Avatar,
		Gender:          userModel.Gender,
		Phone:           userModel.Phone,
		Password:        userModel.Password,
		IdCard:          userModel.IDCard,
		Source:          userModel.Source,
		SystemId:        userModel.SystemId,
		ClientId:        userModel.ClientId,
		UnionId:         userModel.UnionId,
		IsEnable:        userModel.IsEnable,
		LastLoginTime:   timestamppb.New(userModel.LastLoginTime),
		CreateUser:      userModel.CreateUser,
		UpdateUser:      userModel.UpdateUser,
		Token:           common.Md5Encode(timeStr),
		TokenExpireTime: expireTimeStr,
	}
	return userPb
}

func StringTimeConvert(datetime string) time.Time {
	var newTime time.Time
	validate := validator.New()
	err := validate.Var(datetime, "datetime") // "datetime"表示尝试多种时间格式解析
	if err == nil {
		newTime, err = time.Parse("2006-01-01 15:04:05", datetime) // 可以尝试先按RFC3339格式解析，也可根据实际调整
		if err != nil {
			newTime = time.Time{}
		}
	}
	return newTime
}

func StringTimeConvertNew(datetime string) time.Time {
	var newTime time.Time
	layout := "2006-01-01 15:04:05"
	fmt.Println("origin time:", datetime)
	newTime, err := time.Parse(layout, datetime)
	if err != nil {
		newTime = time.Time{}
	}
	//newTime.String()
	return newTime
}
