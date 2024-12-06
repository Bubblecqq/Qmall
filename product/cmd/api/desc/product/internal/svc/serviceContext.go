package svc

import (
	"QMall/product/cmd/api/desc/product/internal/config"
	"QMall/product/cmd/rpc/product/product"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	ProductRpcConf product.ProductZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		ProductRpcConf: product.NewProductZrpcClient(zrpc.MustNewClient(c.ProductRpcConf)),
	}
}
