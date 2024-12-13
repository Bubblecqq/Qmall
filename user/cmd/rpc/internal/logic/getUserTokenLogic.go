package logic

import (
	"context"

	"QMall/user/cmd/rpc/internal/svc"
	"QMall/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTokenLogic {
	return &GetUserTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserTokenLogic) GetUserToken(in *pb.TokenReq) (*pb.TokenResp, error) {
	// todo: add your logic here and delete this line

	return &pb.TokenResp{}, nil
}
