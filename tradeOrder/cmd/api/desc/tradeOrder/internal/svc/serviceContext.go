package svc

import (
	"QMall/product/cmd/rpc/product/product"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/config"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/tradeorder"
	"QMall/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	TradeOrderRpcConf tradeorder.TradeOrderZrpcClient
	ProductRPC        product.ProductZrpcClient
	UserRPC           user.UserZrpcClient
	ShoppingCartRPC   shoppingcart.ShoppingCartZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	productTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.ProductRPC)
	productClient := product.NewProductZrpcClient(productTarget)

	userTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.UserRPC)
	userClient := user.NewUserZrpcClient(userTarget)

	shoppingCartTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.ShoppingCartRPC)
	shoppingCartClient := shoppingcart.NewShoppingCartZrpcClient(shoppingCartTarget)

	return &ServiceContext{
		Config:            c,
		TradeOrderRpcConf: tradeorder.NewTradeOrderZrpcClient(zrpc.MustNewClient(c.TradeOrderRpcConf)),
		ProductRPC:        productClient,
		UserRPC:           userClient,
		ShoppingCartRPC:   shoppingCartClient,
	}
}
