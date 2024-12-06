package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"goZeroDemo4/user/cmd/api/desc/user/internal/svc"
	"goZeroDemo4/user/cmd/api/desc/user/internal/types"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据id删除用户
func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
