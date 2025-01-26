package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestName(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         "192.168.23.233:6379",
		Password:     "123456",
		PoolSize:     100,
		MinIdleConns: 50,
	})
	repository := NewSecKillRepository(nil, redisClient)

	ctx := context.Background()
	cacheKey := fmt.Sprintf("secKillQuota:%d", 83)
	_ = repository.SetSecKillStockWithCache2(ctx, cacheKey, 5000)

}
