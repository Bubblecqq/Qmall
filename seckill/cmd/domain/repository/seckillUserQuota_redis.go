package repository

import (
	"QMall/seckill/cmd/domain/model"
	"context"
)

type ISecKillUserQuotaRepo interface {
	GetSecKillUserQuotaWithCache(ctx context.Context, cacheKey string) (*model.SecKillUserQuota, error)
	SetSecKillUserQuotaWithCache(ctx context.Context, cacheKey string, secKillUserQuota *model.SecKillUserQuota) error
}

func (s *SecKillRepository) GetSecKillUserQuotaWithCache(ctx context.Context, cacheKey string) (*model.SecKillUserQuota, error) {

	return nil, nil
}

func (s *SecKillRepository) SetSecKillUserQuotaWithCache(ctx context.Context, cacheKey string, secKillUserQuota *model.SecKillUserQuota) error {

	return nil
}
