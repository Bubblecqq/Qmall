package user

import (
	"context"
	"goZeroDemo4/user/cmd/api/desc/user/internal/svc"
	types2 "goZeroDemo4/user/cmd/api/desc/user/internal/types"
	"goZeroDemo4/user/cmd/domain/model"
	"goZeroDemo4/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据id获取用户
func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types2.GetUserByIdReq) (resp *types2.GetUserByIdResp, err error) {

	result, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	usermodel := model.PbUserModelConvert(result.User)
	//getUser := types.User{
	//	Id:              usermodel.Id,
	//	CreateTime:      usermodel.CreateTime.String(),
	//	UpdateTime:      usermodel.UpdateTime.String(),
	//	IsDeleted:       usermodel.IsDeleted,
	//	Account:         usermodel.Account,
	//	Avatar:          usermodel.Avatar,
	//	Gender:          usermodel.Gender,
	//	Phone:           usermodel.Phone,
	//	Password:        usermodel.Password,
	//	IDCard:          usermodel.IDCard,
	//	Source:          usermodel.Source,
	//	SystemId:        usermodel.SystemId,
	//	ClientId:        usermodel.ClientId,
	//	UnionId:         usermodel.UnionId,
	//	IsEnable:        usermodel.IsEnable,
	//	LastLoginTime:   usermodel.LastLoginTime.String(),
	//	CreateUser:      usermodel.CreateUser,
	//	UpdateUser:      usermodel.UpdateUser,
	//	Token:           usermodel.Token,
	//	SessionId:       usermodel.SessionId,
	//	TokenExpireTime: usermodel.TokenExpireTime,
	//}
	resp = new(types2.GetUserByIdResp)
	resp.User = types2.ConvertResponseUser(usermodel)
	//resp.User = getUser
	l.Info("current user>\n", resp.User)
	return
}
