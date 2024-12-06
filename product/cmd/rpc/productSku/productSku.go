package main

import (
	"flag"
	"fmt"

	"goZeroDemo4/product/cmd/rpc/productSku/internal/config"
	"goZeroDemo4/product/cmd/rpc/productSku/internal/server"
	"goZeroDemo4/product/cmd/rpc/productSku/internal/svc"
	"goZeroDemo4/product/cmd/rpc/productSku/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/productSku.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterProductSkuServer(grpcServer, server.NewProductSkuServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
