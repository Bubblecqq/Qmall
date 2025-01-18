package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ISecKillRepository interface {
}

type SecKillRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func NewSecKillRepository(mysqlClient *gorm.DB, redisClient *redis.Client) ISecKillRepository {
	return &SecKillRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
