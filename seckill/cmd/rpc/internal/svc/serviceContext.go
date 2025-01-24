package svc

import (
	"QMall/common"
	"QMall/product/cmd/rpc/product/product"
	kafka2 "QMall/seckill/cmd/domain/kafka"
	"QMall/seckill/cmd/domain/repository"
	"QMall/seckill/cmd/rpc/internal/config"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config            config.Config
	Redis             *redis.Client
	SecKillRepository repository.ISecKillRepository
	ProductRPC        product.ProductZrpcClient
	//KqPusherClient    *kafka.Writer
	KqPusherClient2 kafka2.ISecKillOrderPusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := fmt.Sprintf(common.MysqlConnect, c.Mysql.User, c.Mysql.Pass, c.Mysql.Host, c.Mysql.Database, c.Mysql.Charset)

	directory := mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	})
	db, err := gorm.Open(directory, &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Host,
		Password:     c.Redis.Pass,
		PoolSize:     100,
		MinIdleConns: 50,
	})
	productTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.ProductRPC)
	productClient := product.NewProductZrpcClient(productTarget)

	return &ServiceContext{
		Config:            c,
		Redis:             redisClient,
		SecKillRepository: repository.NewSecKillRepository(db, redisClient),
		ProductRPC:        productClient,
		KqPusherClient2:   kafka2.NewSecKillOrderPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		//KqPusherClient: &kafka.Writer{
		//	Addr:         kafka.TCP(c.KqPusherConf.Brokers...),
		//	Topic:        c.KqPusherConf.Topic,
		//	RequiredAcks: kafka.RequireAll,
		//	Async:        true,
		//	Balancer:     &kafka.LeastBytes{},
		//},
	}
}
