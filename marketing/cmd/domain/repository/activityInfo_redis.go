package repository

import (
	"QMall/marketing/cmd/domain/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type IActivityInfoRepo interface {
	GetActivityInfoWithCache(ctx context.Context, cacheKey string) (*model.ActivityInfo, error)
	SetActivityInfoWithCache(ctx context.Context, cacheKey string, activityInfo *model.ActivityInfo) error
}

func (a *ActivityRepository) GetActivityInfoWithCache(ctx context.Context, cacheKey string) (*model.ActivityInfo, error) {
	fmt.Printf("[Redis] 正在查询缓存中...key：%v\n", cacheKey)
	activityInfo := &model.ActivityInfo{}

	result, err := a.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("Cache miss for key: %s\n", cacheKey)
			return nil, errors.New("cache miss")
		}
		fmt.Printf("Error getting data from Redis for key %s: %v\n", cacheKey, err)
		return nil, err
	}

	err = json.Unmarshal([]byte(result), activityInfo)
	if err != nil {
		// 反序列化出错，返回错误信息
		fmt.Printf("Error unmarshaling data from Redis for key %s: %v\n", cacheKey, err)
		return nil, err
	}
	return activityInfo, nil
}

func (a *ActivityRepository) SetActivityInfoWithCache(ctx context.Context, cacheKey string, activityInfo *model.ActivityInfo) error {
	fmt.Printf("[Redis] 正在设置缓存中...key：%v\n", cacheKey)
	data, err := json.Marshal(activityInfo)
	if err != nil {
		fmt.Printf("Error marshaling activity info for key %s: %v\n", cacheKey, err)
		return err
	}
	setNX := a.redisClient.SetNX(ctx, cacheKey, string(data), 5*time.Minute)
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
