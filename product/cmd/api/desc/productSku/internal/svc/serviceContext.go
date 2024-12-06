package svc

import (
	"goZeroDemo4/product/cmd/api/desc/productSku/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
