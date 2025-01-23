package svc

import (
	"QMall/marketing/cmd/rpc/activity"
	"QMall/product/cmd/rpc/product/product"
	"QMall/seckill/cmd/api/desc/seckill/internal/config"
	"QMall/seckill/cmd/rpc/seckill"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/tradeorder"
	"QMall/user/cmd/rpc/user"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	SecKillRpcConf  seckill.SecKill
	OrderRPC        tradeorder.TradeOrderZrpcClient
	ProductRPC      product.ProductZrpcClient
	UserRPC         user.UserZrpcClient
	ShoppingCartRPC shoppingcart.ShoppingCartZrpcClient
	ActivityRPC     activity.ActivityZrpcClient
	KqPusherClient  *kafka.Writer
}

func NewServiceContext(c config.Config) *ServiceContext {
	productTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.ProductRPC)
	productClient := product.NewProductZrpcClient(productTarget)

	userTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.UserRPC)
	userClient := user.NewUserZrpcClient(userTarget)

	orderTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.OrderRPC)
	orderClient := tradeorder.NewTradeOrderZrpcClient(orderTarget)

	shoppingCartTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.ShoppingCartRPC)
	shoppingCartClient := shoppingcart.NewShoppingCartZrpcClient(shoppingCartTarget)

	activityTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.ActivityRPC)
	activityClient := activity.NewActivityZrpcClient(activityTarget)

	return &ServiceContext{
		Config:          c,
		SecKillRpcConf:  seckill.NewSecKill(zrpc.MustNewClient(c.SecKillRpcConf)),
		ProductRPC:      productClient,
		UserRPC:         userClient,
		OrderRPC:        orderClient,
		ShoppingCartRPC: shoppingCartClient,
		ActivityRPC:     activityClient,
		KqPusherClient: &kafka.Writer{
			Addr:         kafka.TCP(c.KqPusherConf.Brokers...),
			Topic:        c.KqPusherConf.Topic,
			RequiredAcks: kafka.RequireAll,
			Async:        true,
			Balancer:     &kafka.LeastBytes{},
		},
	}
}
