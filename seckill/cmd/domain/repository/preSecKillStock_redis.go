package repository

import (
	"QMall/seckill/cmd/domain/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type IPreSecKillStockRepo interface {
	DecreaseStockWithLock(ctx context.Context, cacheKey string, products_id, quantity int64) (bool, error)
	DecreaseStockWithLockV3(ctx context.Context, cacheKey string, products_id, quantity int64) (bool, error)
	SetSecKillStockWithCacheV3(ctx context.Context, cacheKey string, quantity int64) error
	GetSecKillStockWithCache2(ctx context.Context, cacheKey string) (int64, error)
	SetSecKillStockWithCache2(ctx context.Context, cacheKey string, quantity int64) error

	GetSecKillStockWithCache(ctx context.Context, cacheKey string) (*model.SecKillStock, error)
	SetSecKillStockWithCache(ctx context.Context, cacheKey string, stock *model.SecKillStock) error

	// 库存预扣
	DecreaseStockWithLock2(ctx context.Context, cacheKey string, quantity int64) (bool, error)

	//
	SecKillStockRollbackStockInCache(ctx context.Context, cacheKey string, quantity int64) error
}

func (s *SecKillRepository) DecreaseStockWithLockV3(ctx context.Context, cacheKey string, products_id, quantity int64) (bool, error) {
	lockKey := fmt.Sprintf("product_lock:%d", products_id)
	locked, err := s.redisClient.SetNX(ctx, lockKey, "locked", 10*time.Second).Result()
	if err != nil {
		return false, fmt.Errorf("获取分布式锁时出错：%v", err)
	}
	if !locked {
		// 重试机制
		maxRetries := 3
		for i := 0; i < maxRetries; i++ {
			time.Sleep(100 * time.Millisecond)
			locked, err = s.redisClient.SetNX(ctx, lockKey, "locked", 10*time.Second).Result()
			if err == nil && locked {
				break
			}
			if i == maxRetries-1 {
				return false, fmt.Errorf("获取分布式锁失败，请稍后尝试！")
			}
		}
	}
	// 锁续期协程
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-ticker.C:
				_, err := s.redisClient.Expire(ctx, lockKey, 10*time.Second).Result()
				if err != nil {
					fmt.Printf("续期分布式锁失败：%v\n", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	defer s.redisClient.Del(ctx, lockKey)
	result, err := s.redisClient.Eval(ctx, DecreaseStockLuaScript, []string{cacheKey}, quantity).Int64()
	if err != nil {
		return false, fmt.Errorf("执行 Lua 脚本出错：%v", err)
	}

	if result == -1 {
		return false, nil
	}
	return true, nil
}

func (s *SecKillRepository) SecKillStockRollbackStockInCache(ctx context.Context, cacheKey string, quantity int64) error {
	incrBy := s.redisClient.IncrBy(ctx, cacheKey, quantity)
	return incrBy.Err()
}

func (s *SecKillRepository) GetSecKillStockWithCache2(ctx context.Context, cacheKey string) (int64, error) {
	fmt.Printf("[Redis] 正在查询缓存中...key：%v\n", cacheKey)

	quantity, err := s.redisClient.Get(ctx, cacheKey).Int64()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("Cache miss for key: %s\n", cacheKey)
			return 0, errors.New("cache miss")
		}
		fmt.Printf("Error getting data from Redis for key %s: %v\n", cacheKey, err)
		return 0, err
	}
	return quantity, nil
}
func (s *SecKillRepository) SetSecKillStockWithCacheV3(ctx context.Context, cacheKey string, quantity int64) error {
	fmt.Printf("[Redis] 正在设置缓存中...key：%v\n", cacheKey)

	err := s.redisClient.Set(ctx, cacheKey, quantity, 1*time.Minute).Err()
	if err != nil {
		fmt.Printf("Error setting data in Redis for key %s: %v\n", cacheKey, err)
		return err
	}
	return nil
}

func (s *SecKillRepository) SetSecKillStockWithCache2(ctx context.Context, cacheKey string, quantity int64) error {
	fmt.Printf("[Redis] 正在设置缓存中...key：%v\n", cacheKey)

	setNX := s.redisClient.SetNX(ctx, cacheKey, quantity, 1*time.Minute)
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

func (s *SecKillRepository) DecreaseStockWithLock(ctx context.Context, cacheKey string, products_id, quantity int64) (bool, error) {

	//stockKey := fmt.Sprintf("product_stock:%d", products_id)
	lockKey := fmt.Sprintf("product_lock:%d", products_id)
	locked, err := s.redisClient.SetNX(ctx, lockKey, "locked", 10*time.Second).Result()
	if err != nil {
		return false, fmt.Errorf("获取分布式锁时出错：%v", err)
	}
	if !locked {
		return false, fmt.Errorf("获取分布式锁失败，请稍后尝试！")
	}
	defer s.redisClient.Del(ctx, lockKey)
	//检查库存并扣除
	stock, err := s.redisClient.Get(ctx, cacheKey).Int64()

	if err != nil {
		if err == redis.Nil {
			return false, fmt.Errorf("获取库存信息>>>秒杀商品不存在")
		}
		return false, fmt.Errorf("获取秒杀商品时出错：%v", err)
	}
	if stock < quantity {
		return false, nil
	}
	_, err = s.redisClient.DecrBy(ctx, cacheKey, quantity).Result()

	if err != nil {
		return false, fmt.Errorf("扣除秒杀商品时出错：%v", err)
	}
	return true, nil
}

func (s *SecKillRepository) GetSecKillStockWithCache(ctx context.Context, cacheKey string) (*model.SecKillStock, error) {
	fmt.Printf("[Redis] 正在查询缓存中...key：%v\n", cacheKey)
	quantity := &model.SecKillStock{}

	result, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("Cache miss for key: %s\n", cacheKey)
			return nil, errors.New("cache miss")
		}
		fmt.Printf("Error getting data from Redis for key %s: %v\n", cacheKey, err)
		return nil, err
	}

	err = json.Unmarshal([]byte(result), quantity)
	if err != nil {
		// 反序列化出错，返回错误信息
		fmt.Printf("Error unmarshaling data from Redis for key %s: %v\n", cacheKey, err)
		return nil, err
	}
	return quantity, nil
}
func (s *SecKillRepository) SetSecKillStockWithCache(ctx context.Context, cacheKey string, stock *model.SecKillStock) error {
	fmt.Printf("[Redis] 正在设置缓存中...key：%v\n", cacheKey)
	data, err := json.Marshal(stock)
	if err != nil {
		fmt.Printf("Error marshaling activity info for key %s: %v\n", cacheKey, err)
		return err
	}
	setNX := s.redisClient.SetNX(ctx, cacheKey, string(data), 1*time.Minute)
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

func (s *SecKillRepository) DecreaseStockWithLock2(ctx context.Context, cacheKey string, quantity int64) (bool, error) {
	fmt.Printf("[Redis] 正在查询缓存中...key：%v\n", cacheKey)

	stock := &model.SecKillStock{}
	stock, err := s.GetSecKillStockWithCache(ctx, cacheKey)
	if err != nil {
		if err.Error() != "cache miss" {
			return false, err
		}
	}
	lockKey := fmt.Sprintf("product_lock:%d", stock.ProductsId)
	locked, err := s.redisClient.SetNX(ctx, lockKey, "locked", 10*time.Second).Result()
	if err != nil {
		return false, fmt.Errorf("获取分布式锁时出错：%v", err)
	}
	if !locked {
		return false, fmt.Errorf("获取分布式锁失败，请稍后尝试！")
	}
	defer s.redisClient.Del(ctx, lockKey)

	if quantity > stock.Stock {
		return false, errors.New("当前秒杀数量超过库存数！")
	}
	stock.Stock -= quantity
	updateTime := time.Now()
	stock.UpdateTime = &updateTime
	err = s.SetSecKillStockWithCache(ctx, cacheKey, stock)
	if err != nil {
		return false, err
	}
	return true, nil
}

const DecreaseStockLuaScript = `
		local stock = tonumber(redis.call('GET', KEYS[1]))
        if stock and stock >= tonumber(ARGV[1]) then
            return redis.call('DECRBY', KEYS[1], ARGV[1])
        else
            return -1
        end
`
