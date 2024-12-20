package user

import (
	"QMall/user/cmd/rpc/user"
	"context"

	"QMall/user/cmd/api/desc/user/internal/svc"
	"QMall/user/cmd/api/desc/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分布式Token提取
func NewGetUserTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTokenLogic {
	return &GetUserTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserTokenLogic) GetUserToken(req *types.GetTokenReq) (resp *types.GetTokenResp, err error) {
	resp = new(types.GetTokenResp)
	// 分布式获取token
	tokenResp, err := l.svcCtx.UserRpc.GetUserToken(l.ctx, &user.TokenReq{
		Uuid: req.Uuid,
	})

	resp.Token = tokenResp.Token
	resp.IsLogin = tokenResp.IsLogin
	return
}
