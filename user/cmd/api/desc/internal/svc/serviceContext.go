package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goZeroDemo4/user/cmd/api/desc/internal/config"
	"goZeroDemo4/user/cmd/rpc/user"
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
