package main

import (
	"flag"
	"fmt"

	"QMall/marketing/cmd/api/desc/marketing/internal/config"
	"QMall/marketing/cmd/api/desc/marketing/internal/handler"
	"QMall/marketing/cmd/api/desc/marketing/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "D:\\development\\MicroService\\Qmall\\marketing\\cmd\\api\\desc\\marketing\\etc\\activity.yaml", "the config file")

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
