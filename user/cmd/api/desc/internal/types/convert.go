package types

import "goZeroDemo4/user/cmd/domain/model"

func ConvertResponseUser(usermodel *model.User) User {
	typeUser := User{
		Id:              usermodel.Id,
		CreateTime:      usermodel.CreateTime.String(),
		UpdateTime:      usermodel.UpdateTime.String(),
		IsDeleted:       usermodel.IsDeleted,
		Account:         usermodel.Account,
		Avatar:          usermodel.Avatar,
		Gender:          usermodel.Gender,
		Phone:           usermodel.Phone,
		Password:        usermodel.Password,
		IDCard:          usermodel.IDCard,
		Source:          usermodel.Source,
		SystemId:        usermodel.SystemId,
		ClientId:        usermodel.ClientId,
		UnionId:         usermodel.UnionId,
		IsEnable:        usermodel.IsEnable,
		LastLoginTime:   usermodel.LastLoginTime.String(),
		CreateUser:      usermodel.CreateUser,
		UpdateUser:      usermodel.UpdateUser,
		Token:           usermodel.Token,
		SessionId:       usermodel.SessionId,
		TokenExpireTime: usermodel.TokenExpireTime,
	}
	return typeUser
}
