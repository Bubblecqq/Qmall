package main

import (
	"QMall/product/cmd/api/desc/product/internal/config"
	"QMall/product/cmd/api/desc/product/internal/handler"
	"QMall/product/cmd/api/desc/product/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "D:\\development\\MicroService\\QMall\\product\\cmd\\api\\desc\\product\\etc\\product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//consul.RegisterService()
	server.Start()
}
