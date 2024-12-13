package svc

import (
	"QMall/product/cmd/rpc/product/product"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/config"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"QMall/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	ShoppingCartRpcConf shoppingcart.ShoppingCartZrpcClient
	ProductRPC          product.ProductZrpcClient
	UserRPC             user.UserZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	target, _ := zrpc.NewClientWithTarget(c.RemoteCall.ProductRPC)
	client := product.NewProductZrpcClient(target)

	userTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.UserRPC)
	userClient := user.NewUserZrpcClient(userTarget)
	return &ServiceContext{
		Config:              c,
		ShoppingCartRpcConf: shoppingcart.NewShoppingCartZrpcClient(zrpc.MustNewClient(c.ShoppingCartRpcConf)),
		ProductRPC:          client,
		UserRPC:             userClient,
	}
}
