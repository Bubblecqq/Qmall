package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ITradeOrderProductRepository interface {
}

type TradeOrderProductRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func NewTradeOrderProductRepository(mysqlClient *gorm.DB, redisClient *redis.Client) ITradeOrderProductRepository {
	return &TradeOrderProductRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
