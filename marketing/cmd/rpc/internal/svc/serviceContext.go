package svc

import (
	"QMall/common"
	"QMall/marketing/cmd/domain/repository"
	"QMall/marketing/cmd/rpc/internal/config"
	"QMall/product/cmd/rpc/product/product"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config             config.Config
	Redis              *redis.Client
	ActivityRepository repository.IActivityRepository
	ProductRPC         product.ProductZrpcClient
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
		Config:             c,
		ActivityRepository: repository.NewActivityRepository(db, redisClient),
		Redis:              redisClient,
		ProductRPC:         productClient,
	}
}
