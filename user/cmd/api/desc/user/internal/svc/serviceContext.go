package svc

import (
	"QMall/user/cmd/api/desc/user/internal/config"
	"QMall/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
