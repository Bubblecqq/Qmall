package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type IPreSecKillStockRepo interface {
	DecreaseStockWithLock(ctx context.Context, products_id, quantity int64) (bool, error)
}

type PreSecKillStockRepo struct {
	RedisClient *redis.Client
}

func (p *PreSecKillStockRepo) DecreaseStockWithLock(ctx context.Context, products_id, quantity int64) (bool, error) {

	stockKey := fmt.Sprintf("product_stock:%d", products_id)
	lockKey := fmt.Sprintf("product_lock:%d", products_id)
	locked, err := p.RedisClient.SetNX(ctx, lockKey, "locked", 10*time.Second).Result()
	if err != nil {
		return false, fmt.Errorf("获取分布式锁时出错：%v", err)
	}
	if !locked {
		return false, fmt.Errorf("获取分布式锁失败，请稍后尝试！")
	}
	defer p.RedisClient.Del(ctx, lockKey)
	//检查库存并扣除
	stock, err := p.RedisClient.Get(ctx, stockKey).Int64()

	if err != nil {
		if err == redis.Nil {
			return false, fmt.Errorf("秒杀商品不存在")
		}
		return false, fmt.Errorf("获取秒杀商品时出错：%v", err)
	}
	if stock < quantity {
		return false, nil
	}
	_, err = p.RedisClient.DecrBy(ctx, stockKey, quantity).Result()

	if err != nil {
		return false, fmt.Errorf("扣除秒杀商品时出错：%v", err)
	}
	return true, nil
}

func NewPreSecKillStockRepo() IPreSecKillStockRepo {
	return &PreSecKillStockRepo{}
}
