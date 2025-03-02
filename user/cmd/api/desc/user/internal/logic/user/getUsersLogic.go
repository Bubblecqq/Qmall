package user

import (
	"QMall/user/cmd/api/desc/user/internal/svc"
	"QMall/user/cmd/api/desc/user/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户列表
func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersLogic) GetUsers(req *types.GetUserListReq) (resp *types.GetUserListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
