package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql  Mysql
	Cache  cache.CacheConf
	Consul consul.Conf
}
type Mysql struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}
