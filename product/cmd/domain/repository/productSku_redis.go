package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type ProductSkuRepo interface {
	GetProductSkuWithCache(ctx context.Context, cacheKey string) (int64, error)
	SetProductSkuWithCache(ctx context.Context, cacheKey string, quantity int64) error
}

func (p *ProductRepository) GetProductSkuWithCache(ctx context.Context, cacheKey string) (int64, error) {
	fmt.Printf("[Redis] 正在查询缓存中...key：%v\n", cacheKey)
	quantity := int64(0)

	result, err := p.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("Cache miss for key: %s\n", cacheKey)
			return 0, errors.New("cache miss")
		}
		fmt.Printf("Error getting data from Redis for key %s: %v\n", cacheKey, err)
		return 0, err
	}

	err = json.Unmarshal([]byte(result), quantity)
	if err != nil {
		// 反序列化出错，返回错误信息
		fmt.Printf("Error unmarshaling data from Redis for key %s: %v\n", cacheKey, err)
		return 0, err
	}
	return quantity, nil
}

func (p *ProductRepository) SetProductSkuWithCache(ctx context.Context, cacheKey string, quantity int64) error {
	fmt.Printf("[Redis] 正在设置缓存中...key：%v\n", cacheKey)
	data, err := json.Marshal(quantity)
	if err != nil {
		fmt.Printf("Error marshaling activity info for key %s: %v\n", cacheKey, err)
		return err
	}
	setNX := p.redisClient.SetNX(ctx, cacheKey, string(data), 5*time.Minute)
	if setNX.Err() != nil {
		fmt.Printf("Error setting data in Redis for key %s: %v\n", cacheKey, setNX.Err())
		return setNX.Err()
	}

	if !setNX.Val() {
		fmt.Printf("Failed to set cache for key %s: key already exists\n", cacheKey)
		return errors.New("failed to set cache: key already exists")
	}
	return nil
}
