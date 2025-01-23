package main

import (
	kafka2 "QMall/seckill/cmd/domain/kafka"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"QMall/seckill/cmd/rpc/internal/config"
	"QMall/seckill/cmd/rpc/internal/server"
	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "D:\\development\\MicroService\\Qmall\\seckill\\cmd\\rpc\\etc\\seckill.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterSecKillServer(grpcServer, server.NewSecKillServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	fmt.Println("正在启动Kafka消费者实例......")
	consumerConfig := &kafka2.ConsumerConfig{
		MysqlConfig: kafka2.Mysql{
			Host:     c.Mysql.Host,
			User:     c.Mysql.User,
			Pass:     c.Mysql.Pass,
			Database: c.Mysql.Database,
			Charset:  c.Mysql.Charset,
		},
		RedisConfig: kafka2.Redis{
			Host: c.Redis.Host,
			Pass: c.Redis.Pass,
		},
		KafkaConfig: kafka2.Kafka{
			Brokers: c.KqPusherConf.Brokers,
			Topic:   c.KqPusherConf.Topic,
			GroupId: c.KqPusherConf.GroupId,
		},
	}
	consumer := kafka2.NewSecKillOrderConsumer(consumerConfig)

	go consumer.ReaderSecKillOrder(context.Background())

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
