package logic

import (
	"QMall/user/cmd/domain/model"
	"context"

	"QMall/user/cmd/rpc/internal/svc"
	"QMall/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *pb.CreateUserReq) (*pb.CreateUserResp, error) {

	user := &model.User{
		ClientId: in.ClientId,
		SystemId: in.SystemId,
		Phone:    in.Phone,
		Password: in.Password,
	}
	err := l.svcCtx.UserRepository.AddUser(user.ClientId, user.Phone, user.Password, user.SystemId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResp{}, nil
}
