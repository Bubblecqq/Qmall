package svc

import (
	"QMall/marketing/cmd/api/desc/marketing/internal/config"
	"QMall/marketing/cmd/rpc/activity"
	"QMall/product/cmd/rpc/product/product"
	"QMall/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	ActivityRpcConf activity.ActivityZrpcClient
	ProductRPC      product.ProductZrpcClient
	UserRPC         user.UserZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	productTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.ProductRPC)
	productClient := product.NewProductZrpcClient(productTarget)

	userTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.UserRPC)
	userClient := user.NewUserZrpcClient(userTarget)

	return &ServiceContext{
		Config:          c,
		ProductRPC:      productClient,
		UserRPC:         userClient,
		ActivityRpcConf: activity.NewActivityZrpcClient(zrpc.MustNewClient(c.ActivityRpcConf)),
	}
}
