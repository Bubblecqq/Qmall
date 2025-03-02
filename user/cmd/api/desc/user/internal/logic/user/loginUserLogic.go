package user

import (
	"QMall/user/cmd/api/desc/user/internal/svc"
	types2 "QMall/user/cmd/api/desc/user/internal/types"
	"QMall/user/cmd/domain/model"
	"QMall/user/cmd/rpc/user"
	"context"

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

func (l *LoginUserLogic) LoginUser(req *types2.LoginUserReq) (resp *types2.LoginUserResp, err error) {

	loginUser, err := l.svcCtx.UserRpc.LoginUser(l.ctx, &user.LoginUserReq{
		ClientId:         req.ClientId,
		SystemId:         req.SystemId,
		Phone:            req.Phone,
		VerificationCode: req.VerificationCode,
	})
	loginUserModel := model.PbUserModelConvert(loginUser.User)
	resp = new(types2.LoginUserResp)
	resp.User = types2.ConvertResponseUser(loginUserModel)
	if err != nil {
		return nil, err
	}
	return
}
