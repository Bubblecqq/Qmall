package repository

import (
	"QMall/common"
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/rpc/pb"
	"QMall/seckill/cmd/rpc/seckill"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type ISecKillRepository interface {
	ISecKillProducts
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

	GetSecKillProductsByProductsNum(productsNum string) (*model.SecKillProducts, error)

	GetSecKillUserQuota(user_id, products_id int64) (*model.SecKillUserQuota, error)

	GetSecKillQuotaById(id int64) (*model.SecKillQuota, error)

	UpdateSecKillUserQuotaById(in *pb.UpdateSecKillUserQuotaByIdReq) (*model.SecKillUserQuota, error)

	// IncreaseSecKillOrderWithKafka 添加秒杀订单
	IncreaseSecKillOrderWithKafka(order *model.SecKillOrder) (*model.SecKillOrder, error)

	SaveSecKillUserQuota(in *pb.SaveSecKillUserQuotaReq) (*model.SecKillUserQuota, error)
}

type SecKillRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (s *SecKillRepository) SaveSecKillUserQuota(in *pb.SaveSecKillUserQuotaReq) (*model.SecKillUserQuota, error) {
	fmt.Printf("[MySQL] 正在判断当前用户Id：%v及商品Id：%v是否存在限额....\n", in.UserId, in.ProductsId)
	// 先查询是否存在
	userQuota := &model.SecKillUserQuota{}
	tx := s.mysqlClient.Model(&model.SecKillUserQuota{}).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	err := tx.Where("products_id=? and user_id=?", in.ProductsId, in.UserId).Find(&userQuota).Error
	if err != nil || userQuota.Id <= 0 {
		userQuota, err = s.IncreaseSecKillUserQuota(&seckill.IncreaseSecKillUserQuotaReq{
			UserId:     in.UserId,
			ProductsId: in.ProductsId,
			Num:        in.Num,
			KilledNum:  in.KilledNum,
		})
		if err != nil {
			return nil, err
		}

	} else {
		quotaByProductsId, err := s.GetSecKillQuotaByProductsId(in.ProductsId)

		// 是否超过限额
		if err != nil {
			return nil, err
		}
		// 是否超过限额
		if in.Num > quotaByProductsId.Num || userQuota.KilledNum+in.Num > quotaByProductsId.Num {
			return nil, fmt.Errorf("当前用户秒杀的额度超过商品额度！具体用户秒杀数：%v，商品当前限量：%v", in.Num, quotaByProductsId.Num)
		}
		userQuota.KilledNum += in.Num
		updateTime := time.Now()
		userQuota.UpdateTime = &updateTime
		err = tx.Where("products_id=? and user_id=?", in.ProductsId, in.UserId).Updates(&userQuota).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return userQuota, nil
}

func (s *SecKillRepository) GetSecKillProductsByProductsNum(productsNum string) (*model.SecKillProducts, error) {
	products := &model.SecKillProducts{}

	tx := s.mysqlClient.Model(&model.SecKillProducts{}).Where("products_num=?", productsNum).Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if products.Id <= 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return products, nil
}

func (s *SecKillRepository) IncreaseSecKillOrderWithKafka(order *model.SecKillOrder) (*model.SecKillOrder, error) {
	tx := s.mysqlClient.Model(&model.SecKillOrder{}).Create(order)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return order, nil
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
func (s *SecKillRepository) UpdateSecKillUserQuotaById(in *pb.UpdateSecKillUserQuotaByIdReq) (*model.SecKillUserQuota, error) {

	quotaByProductsId, err := s.GetSecKillQuotaByProductsId(in.ProductsId)

	if err != nil {
		return nil, err
	}
	// 是否超过限额
	if in.Num > quotaByProductsId.Num {
		return nil, errors.New(fmt.Sprintf("当前用户秒杀的额度超过商品额度！具体用户秒杀数：%v，商品当前限量：%v", in.Num, quotaByProductsId.Num))
	}
	userQuota, err := s.GetSecKillUserQuota(in.UserId, in.ProductsId)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	if userQuota.KilledNum+in.Num > quotaByProductsId.Num {
		return nil, errors.New(fmt.Sprintf("当前用户秒杀的额度超过商品额度！具体用户秒杀数：%v，商品当前限量：%v", in.Num, quotaByProductsId.Num))
	}
	userQuota.KilledNum += in.Num
	updateTime := time.Now()
	userQuota.UpdateTime = &updateTime
	tx := s.mysqlClient.Model(&model.SecKillUserQuota{}).
		Where("user_id=? and products_id=?", in.UserId, in.ProductsId).
		Updates(&userQuota)

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
	tx := s.mysqlClient.Model(&model.SecKillQuota{}).Where("products_id=? ", products_id).Find(&quota)
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
		UpdateTime:  &now,
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
	now := time.Now()
	secKillStock.CreateTime = now
	secKillStock.UpdateTime = &now
	tx := s.mysqlClient.Model(&model.SecKillStock{}).Create(&secKillStock)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return secKillStock, nil
}

func (s *SecKillRepository) IncreaseSecKillQuota(in *pb.IncreaseSecKillQuotaReq) (*model.SecKillQuota, error) {
	secKillQuota := &model.SecKillQuota{}
	now := time.Now()
	secKillQuota.CreateTime = now
	secKillQuota.UpdateTime = &now
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
	now := time.Now()
	secKillUserQuota := &model.SecKillUserQuota{
		UserId:     in.UserId,
		ProductsId: in.ProductsId,
		Num:        in.Num,
		CreateTime: now,
		UpdateTime: &now,
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
		SecNum:     common.GenerateSecKillOrderNo(now, in.UserId),
		UpdateTime: &now,
	}
	tx := s.mysqlClient.Model(&model.SecKillRecord{}).Create(&secKillRecord)
	if tx.Error != nil {
		return nil, tx.Error
	}

	//库存扣

	return secKillRecord, nil
}

func (s *SecKillRepository) IncreaseSecKillProducts(in *pb.IncreaseSecKillProductsReq) (*model.SecKillProducts, error) {
	now := time.Now()
	secKillProducts := &model.SecKillProducts{
		CreateTime:   now,
		Price:        in.Price,
		Seller:       in.Seller,
		PictureUrl:   in.PictureUrl,
		ProductsName: in.ProductName,
		UpdateTime:   &now,
		ProductsNum:  in.ProductsNum,
		ProductsId:   in.ProductId,
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
