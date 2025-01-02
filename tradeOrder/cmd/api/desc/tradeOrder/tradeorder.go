package main

import (
	"flag"
	"fmt"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/config"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/handler"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "D:\\development\\MicroService\\Qmall\\tradeOrder\\cmd\\api\\desc\\tradeOrder\\etc\\tradeOrder.yaml", "the config file")

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
