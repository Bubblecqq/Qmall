package user

import (
	"context"
	"goZeroDemo4/user/cmd/domain/model"
	"goZeroDemo4/user/cmd/rpc/user"

	"goZeroDemo4/user/cmd/api/desc/internal/svc"
	"goZeroDemo4/user/cmd/api/desc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 单点用户登录
func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginUserLogic) LoginUser(req *types.LoginUserReq) (resp *types.LoginUserResp, err error) {

	loginUser, err := l.svcCtx.UserRpc.LoginUser(l.ctx, &user.LoginUserReq{
		ClientId:         req.ClientId,
		SystemId:         req.SystemId,
		Phone:            req.Phone,
		VerificationCode: req.VerificationCode,
	})
	loginUserModel := model.PbUserModelConvert(loginUser.User)
	resp = new(types.LoginUserResp)
	resp.User = types.ConvertResponseUser(loginUserModel)
	if err != nil {
		return nil, err
	}
	return
}
