package svc

import (
	"QMall/product/cmd/rpc/product/product"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/config"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	ShoppingCartRpcConf shoppingcart.ShoppingCartZrpcClient
	Target              product.ProductZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	target, _ := zrpc.NewClientWithTarget(c.ShoppingCartRpcConf.Target)
	client := product.NewProductZrpcClient(target)
	return &ServiceContext{
		Config:              c,
		ShoppingCartRpcConf: shoppingcart.NewShoppingCartZrpcClient(zrpc.MustNewClient(c.ShoppingCartRpcConf)),
		Target:              client,
	}
}
