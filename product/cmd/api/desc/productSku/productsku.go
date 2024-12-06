package main

import (
	"flag"
	"fmt"

	"goZeroDemo4/product/cmd/api/desc/productSku/internal/config"
	"goZeroDemo4/product/cmd/api/desc/productSku/internal/handler"
	"goZeroDemo4/product/cmd/api/desc/productSku/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/productsku.yaml", "the config file")

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
