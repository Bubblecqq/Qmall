package repository

import (
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/rpc/pb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type ISecKillRepository interface {

	// IncreaseSecKillStock 添加秒杀库存
	IncreaseSecKillStock(in *pb.IncreaseSecKillStockReq) (*model.SecKillStock, error)

	// IncreaseSecKillQuota 添加秒杀限额
	IncreaseSecKillQuota(in *pb.IncreaseSecKillQuotaReq) (*model.SecKillQuota, error)

	// IncreaseSecKillUserQuota 添加秒杀用户限额
	IncreaseSecKillUserQuota(in *pb.IncreaseSecKillUserQuotaReq) (*model.SecKillUserQuota, error)

	// IncreaseSecKillRecord 添加秒杀记录
	IncreaseSecKillRecord(in *pb.IncreaseSecKillRecordReq) (*model.SecKillRecord, error)

	// IncreaseSecKillProducts 添加秒杀商品
	IncreaseSecKillProducts(secKillProducts *model.SecKillProducts) (*model.SecKillProducts, error)
}

type SecKillRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (s *SecKillRepository) IncreaseSecKillStock(in *pb.IncreaseSecKillStockReq) (*model.SecKillStock, error) {
	secKillStock := &model.SecKillStock{}
	secKillStock.ProductsId = in.ProductsId
	secKillStock.Stock = in.Stock
	secKillStock.CreateTime = time.Now()
	tx := s.mysqlClient.Model(&model.SecKillStock{}).Create(&secKillStock)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return secKillStock, nil
}

func (s *SecKillRepository) IncreaseSecKillQuota(in *pb.IncreaseSecKillQuotaReq) (*model.SecKillQuota, error) {
	secKillQuota := &model.SecKillQuota{}
	secKillQuota.CreateTime = time.Now()
	secKillQuota.Num = in.LimitNumber
	secKillQuota.ProductsId = in.ProductId
	tx := s.mysqlClient.Model(&model.SecKillQuota{}).Create(&secKillQuota)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return secKillQuota, nil
}

func (s *SecKillRepository) IncreaseSecKillUserQuota(in *pb.IncreaseSecKillUserQuotaReq) (*model.SecKillUserQuota, error) {
	secKillUserQuota := &model.SecKillUserQuota{
		UserId:     in.UserId,
		ProductsId: in.ProductsId,
		Num:        in.Num,
		CreateTime: time.Now(),
	}
	tx := s.mysqlClient.Model(&model.SecKillUserQuota{}).Create(&secKillUserQuota)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return secKillUserQuota, nil
}

func (s *SecKillRepository) IncreaseSecKillRecord(in *pb.IncreaseSecKillRecordReq) (*model.SecKillRecord, error) {
	secKillRecord := &model.SecKillRecord{
		UserId:     in.UserId,
		ProductsId: in.ProductsId,
		Price:      in.Price,
		CreateTime: time.Now(),
	}
	tx := s.mysqlClient.Model(&model.SecKillRecord{}).Create(&secKillRecord)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return secKillRecord, nil
}

func (s *SecKillRepository) IncreaseSecKillProducts(secKillProducts *model.SecKillProducts) (*model.SecKillProducts, error) {
	//&model.SecKillProducts{
	//	CreateTime: time.Now(),
	//	Price: in.Price,
	//	Seller: in.Seller,
	//}
	tx := s.mysqlClient.Model(&model.SecKillProducts{}).Create(secKillProducts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return secKillProducts, nil
}

func NewSecKillRepository(mysqlClient *gorm.DB, redisClient *redis.Client) ISecKillRepository {
	return &SecKillRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
