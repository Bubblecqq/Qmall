package logic

import (
	"QMall/user/cmd/domain/model"
	"QMall/user/cmd/rpc/internal/svc"
	"QMall/user/cmd/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	userModel, err := l.svcCtx.UserRepository.GetUserById(in.Id)
	if err != nil {
		return nil, err
	}
	l.Info("Current Query User id:", userModel.Id)

	userPb := model.UserModelConvertPbUser(userModel)
	l.Info("Current Query User id:", userPb)
	return &pb.GetUserByIdResp{
		User: userPb,
	}, nil
}
