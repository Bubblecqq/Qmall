package svc

import (
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/config"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	ShoppingCartRpcConf shoppingcart.ShoppingCartZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		ShoppingCartRpcConf: shoppingcart.NewShoppingCartZrpcClient(zrpc.MustNewClient(c.ShoppingCartRpcConf)),
	}
}
