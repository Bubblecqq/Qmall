package logic

import (
	"QMall/user/cmd/domain/model"
	"context"
	"fmt"
	"strconv"
	"time"

	"QMall/user/cmd/rpc/internal/svc"
	"QMall/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginUserLogic) LoginUser(in *pb.LoginUserReq) (*pb.LoginUserResp, error) {

	login, err := l.svcCtx.UserRepository.Login(in.ClientId, in.Phone, in.SystemId, in.VerificationCode)
	if err != nil {
		return nil, err
	}
	l.Info(fmt.Sprintf("[LoginRPC] Current Login User id:%v\n", login.Id))
	// UserModelConvertPbUser 有生成token
	loginUserPb := model.UserModelConvertPbUser(login)
	l.svcCtx.UserRepository.SetUserToken(strconv.Itoa(int(login.Id)), []byte(login.Token), time.Duration(1)*time.Hour)
	return &pb.LoginUserResp{
		User: loginUserPb,
	}, nil
}
