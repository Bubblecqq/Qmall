package svc

import (
	"QMall/common"
	"QMall/seckill/cmd/domain/repository"
	"QMall/seckill/cmd/rpc/internal/config"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config            config.Config
	Redis             *redis.Client
	SecKillRepository repository.ISecKillRepository
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
		PoolSize:     30,
		MinIdleConns: 30,
	})
	return &ServiceContext{
		Config:            c,
		Redis:             redisClient,
		SecKillRepository: repository.NewSecKillRepository(db, redisClient),
	}
}
