package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goZeroDemo4/product/cmd/api/desc/product/internal/config"
	"goZeroDemo4/product/cmd/rpc/product/product"
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
