package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql          Mysql
	Cache          cache.CacheConf
	SecKillRpcConf zrpc.RpcClientConf
	RemoteCall     RemoteCall
}
type Mysql struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}

type RemoteCall struct {
	UserRPC         string `json:"UserRPC"`
	ProductRPC      string `json:"ProductRPC"`
	ActivityRPC     string `json:"ActivityRPC"`
	OrderRPC        string `json:"OrderRPC"`
	ShoppingCartRPC string `json:"ShoppingCartRPC"`
}
