package main

import (
	"QMall/user/cmd/api/desc/user/internal/config"
	"QMall/user/cmd/api/desc/user/internal/handler"
	"QMall/user/cmd/api/desc/user/internal/svc"
	"flag"
	"fmt"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "D:\\development\\MicroService\\QMall\\user\\cmd\\api\\desc\\user\\etc\\user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
