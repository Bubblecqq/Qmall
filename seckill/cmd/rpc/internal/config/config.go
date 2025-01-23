package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql        Mysql
	Cache        cache.CacheConf
	Consul       consul.Conf
	RemoteCall   RemoteCall
	KqPusherConf KqPusherConf
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

type KqPusherConf struct {
	Brokers []string `json:"brokers"`
	Topic   string   `json:"topic"`
}
