package svc

import (
	"QMall/common"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"QMall/tradeOrder/cmd/domain/repository"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/config"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config               config.Config
	Redis                *redis.Client
	TradeOrderRepository repository.ITradeOrderRepository
	ShoppingCartRPC      shoppingcart.ShoppingCartZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := fmt.Sprintf(common.MysqlConnect, c.Mysql.User, c.Mysql.Pass, c.Mysql.Host, c.Mysql.Database, c.Mysql.Charset)

	dialector := mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	})
	db, err := gorm.Open(dialector, &gorm.Config{
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
		PoolSize:     30,
		MinIdleConns: 30,
	})
	shoppingCartTarget, _ := zrpc.NewClientWithTarget(c.RemoteCall.ShoppingCartRPC)
	shoppingCartClient := shoppingcart.NewShoppingCartZrpcClient(shoppingCartTarget)

	return &ServiceContext{
		Config:               c,
		Redis:                redisClient,
		TradeOrderRepository: repository.NewTradeOrderRepository(db, redisClient),
		ShoppingCartRPC:      shoppingCartClient,
	}
}
