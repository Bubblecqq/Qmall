package svc

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"goZeroDemo4/common"
	"goZeroDemo4/product/cmd/domain/repository"
	"goZeroDemo4/product/cmd/rpc/product/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config            config.Config
	Redis             *redis.Client
	ProductRepository repository.IProductRepository
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
	return &ServiceContext{
		Config:            c,
		Redis:             redisClient,
		ProductRepository: repository.NewProductRepository(db, redisClient),
	}
}
