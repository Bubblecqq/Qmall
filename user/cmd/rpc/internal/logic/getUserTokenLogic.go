package logic

import (
	"QMall/common"
	"context"
	"fmt"
	"strconv"
	"time"

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

	token, err := l.svcCtx.UserRepository.GetUserToken(in.Uuid)
	fmt.Println("Current uuid:", in.Uuid)
	var isLogin bool
	if err != nil {
		isLogin = false
		token = ""
		l.Info(fmt.Sprintf("Current uuid: %v is not login\n", in.Uuid))
	} else {
		isLogin = true
		uuid := common.ToInput(in.Uuid)
		l.svcCtx.UserRepository.SetUserToken(strconv.Itoa(uuid), []byte(token), time.Duration(1)*time.Hour)
		l.Info(fmt.Sprintf("GetUserToken Success with token: %v\n", token))
	}
	fmt.Println("Current user token>>>>>", token)
	return &pb.TokenResp{
		Token:   token,
		IsLogin: isLogin,
	}, nil
}
