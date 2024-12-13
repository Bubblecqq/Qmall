package main

import (
	"flag"
	"fmt"

	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/config"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/handler"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "D:\\development\\MicroService\\QMall\\shoppingCart\\cmd\\api\\desc\\shoppingCart\\etc\\shoppingCart.yaml", "the config file")

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
