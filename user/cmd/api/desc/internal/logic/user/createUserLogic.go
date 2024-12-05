package user

import (
	"context"
	"goZeroDemo4/user/cmd/rpc/user"

	"goZeroDemo4/user/cmd/api/desc/internal/svc"
	"goZeroDemo4/user/cmd/api/desc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建用户
func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {

	_, err = l.svcCtx.UserRpc.CreateUser(l.ctx, &user.CreateUserReq{
		ClientId: req.ClientId,
		SystemId: req.SystemId,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return
}
