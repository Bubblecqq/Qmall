package repository

import (
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/rpc/pb"
	"errors"
	"fmt"
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
	IncreaseSecKillProducts(in *pb.IncreaseSecKillProductsReq) (*model.SecKillProducts, error)

	// IncreaseSecKillOrder 添加秒杀订单
	IncreaseSecKillOrder(in *pb.IncreaseSecKillOrderReq) (*model.SecKillOrder, error)

	GetSecKillQuotaByProductsId(products_id int64) (*model.SecKillQuota, error)

	GetSecKillProductsByProductsId(products_id int64) (*model.SecKillProducts, error)

	GetSecKillUserQuota(user_id, products_id int64) (*model.SecKillUserQuota, error)

	GetSecKillQuotaById(id int64) (*model.SecKillQuota, error)

	UpdateSecKillUserQuotaById(user_id, products_id, killedNum int64) (*model.SecKillUserQuota, error)
}

type SecKillRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (s *SecKillRepository) GetSecKillUserQuota(user_id, products_id int64) (*model.SecKillUserQuota, error) {
	userQuota := &model.SecKillUserQuota{}
	tx := s.mysqlClient.Model(&model.SecKillUserQuota{}).Where("products_id=? and user_id=?", products_id, user_id).Find(&userQuota)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if userQuota.Id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return userQuota, nil
}

// UpdateSecKillUserQuotaById TODO 需要增加锁机制，这里暂不加
func (s *SecKillRepository) UpdateSecKillUserQuotaById(user_id, products_id, killedNum int64) (*model.SecKillUserQuota, error) {

	quotaByProductsId, err := s.GetSecKillQuotaByProductsId(products_id)

	if err != nil {
		return nil, err
	}
	// 是否超过限额
	if killedNum > quotaByProductsId.Num {
		return nil, errors.New(fmt.Sprintf("当前用户秒杀的额度超过商品额度！具体用户秒杀数：%v，商品当前限量：%v", killedNum, quotaByProductsId.Num))
	}
	userQuota, err := s.GetSecKillUserQuota(user_id, products_id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	if userQuota.KilledNum+killedNum > quotaByProductsId.Num {
		return nil, errors.New(fmt.Sprintf("当前用户秒杀的额度超过商品额度！具体用户秒杀数：%v，商品当前限量：%v", killedNum, quotaByProductsId.Num))
	}
	userQuota.KilledNum += killedNum
	updateTime := time.Now()
	userQuota.UpdateTime = &updateTime
	tx := s.mysqlClient.Model(&model.SecKillUserQuota{}).Updates(&userQuota)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return userQuota, nil
}

func (s *SecKillRepository) GetSecKillQuotaById(id int64) (*model.SecKillQuota, error) {
	quota := &model.SecKillQuota{}

	tx := s.mysqlClient.Model(&model.SecKillQuota{}).Find(&quota, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if quota.Id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return quota, nil
}

func (s *SecKillRepository) GetSecKillProductsByProductsId(products_id int64) (*model.SecKillProducts, error) {
	products := &model.SecKillProducts{}
	tx := s.mysqlClient.Model(&model.SecKillProducts{}).Where("products_id=?", products_id).Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if products.Id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return products, nil
}

func (s *SecKillRepository) GetSecKillQuotaByProductsId(products_id int64) (*model.SecKillQuota, error) {
	quota := &model.SecKillQuota{}
	tx := s.mysqlClient.Model(&model.SecKillQuota{}).Where("products_id=?", products_id).Find(&quota)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if quota.Id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return quota, nil
}

func (s *SecKillRepository) IncreaseSecKillOrder(in *pb.IncreaseSecKillOrderReq) (*model.SecKillOrder, error) {
	now := time.Now()
	secKillOrder := &model.SecKillOrder{
		CreateTime:  now,
		Seller:      in.Seller,
		Buyer:       in.Buyer,
		Price:       in.Price,
		ProductsId:  in.ProductsId,
		ProductsNum: in.ProductsNum,
		OrderNum:    in.OrderNum,
		Status:      1,
	}

	tx := s.mysqlClient.Model(&model.SecKillOrder{}).Create(secKillOrder)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return secKillOrder, nil
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

	productsId, err := s.GetSecKillQuotaByProductsId(in.ProductsId)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	if in.Num+in.KilledNum > productsId.Num {
		return nil, errors.New(fmt.Sprintf("当前用户超出限额：%v", in.Num+in.KilledNum))
	}
	secKillUserQuota := &model.SecKillUserQuota{
		UserId:     in.UserId,
		ProductsId: in.ProductsId,
		Num:        in.Num,
		CreateTime: time.Now(),
		KilledNum:  in.KilledNum,
	}

	tx := s.mysqlClient.Model(&model.SecKillUserQuota{}).Create(&secKillUserQuota)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return secKillUserQuota, nil
}

func (s *SecKillRepository) IncreaseSecKillRecord(in *pb.IncreaseSecKillRecordReq) (*model.SecKillRecord, error) {
	now := time.Now()
	secKillRecord := &model.SecKillRecord{
		UserId:     in.UserId,
		ProductsId: in.ProductsId,
		Price:      in.Price,
		CreateTime: now,
		OrderNum:   in.OrderNo,
		Status:     1,
		SecNum:     in.SecNo,
	}
	tx := s.mysqlClient.Model(&model.SecKillRecord{}).Create(&secKillRecord)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//库存扣

	return secKillRecord, nil
}

func (s *SecKillRepository) IncreaseSecKillProducts(in *pb.IncreaseSecKillProductsReq) (*model.SecKillProducts, error) {
	secKillProducts := &model.SecKillProducts{
		CreateTime:   time.Now(),
		Price:        in.Price,
		Seller:       in.Seller,
		PictureUrl:   in.PictureUrl,
		ProductsName: in.ProductName,
	}

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
