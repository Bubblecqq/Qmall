package repository

import (
	"QMall/tradeOrder/cmd/domain/model"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type ITradeOrderProductRepository interface {
	AddOrderProduct(orderProduct *model.TradeOrderProduct) (*model.TradeOrderProduct, error)
	GetOrderProductById(id int64) (*model.TradeOrderProduct, error)
	GetOrderProductList() []*model.TradeOrderProduct
	GetOrderProductByOrderId(orderId int64) (*model.TradeOrderProduct, error)
	GetOrderProductByUserId(userId int64) (*model.TradeOrderProduct, error)
	UpdateOrderProduct(orderProduct *model.TradeOrderProduct) (*model.TradeOrderProduct, error)

	// BatchCreateOrderProduct 批量添加订单商品
	BatchCreateOrderProduct(orderProductList []*model.TradeOrderProduct) error
}

type TradeOrderProductRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (t *TradeOrderProductRepository) BatchCreateOrderProduct(orderProductList []*model.TradeOrderProduct) error {

	tx := t.mysqlClient.Model(&model.TradeOrderProduct{}).Create(&orderProductList)
	if tx.Error != nil {
		fmt.Printf("批量添加订单商品失败！原因见：%v\n", tx.Error)
	}
	return tx.Error
}

func (t *TradeOrderProductRepository) AddOrderProduct(orderProduct *model.TradeOrderProduct) (*model.TradeOrderProduct, error) {

	tx := t.mysqlClient.Model(&model.TradeOrderProduct{}).Create(orderProduct)
	return orderProduct, tx.Error
}

func (t *TradeOrderProductRepository) GetOrderProductById(id int64) (*model.TradeOrderProduct, error) {
	orderProduct := &model.TradeOrderProduct{}
	tx := t.mysqlClient.Model(&model.TradeOrderProduct{}).Find(&orderProduct, id)
	return orderProduct, tx.Error
}

func (t *TradeOrderProductRepository) GetOrderProductList() []*model.TradeOrderProduct {
	var orderProductList []*model.TradeOrderProduct
	tx := t.mysqlClient.Model(&model.TradeOrderProduct{}).Find(&orderProductList)
	if tx.Error != nil {
		return []*model.TradeOrderProduct{}
	}
	return orderProductList
}

func (t *TradeOrderProductRepository) GetOrderProductByOrderId(orderId int64) (*model.TradeOrderProduct, error) {
	orderProduct := &model.TradeOrderProduct{}
	tx := t.mysqlClient.Model(&model.TradeOrderProduct{}).Where("order_id= ?", orderId).Find(&orderProduct)
	return orderProduct, tx.Error
}

func (t *TradeOrderProductRepository) GetOrderProductByUserId(userId int64) (*model.TradeOrderProduct, error) {
	orderProduct := &model.TradeOrderProduct{}
	tx := t.mysqlClient.Model(&model.TradeOrderProduct{}).Where("user_id= ?", userId).Find(&orderProduct)
	return orderProduct, tx.Error
}

func (t *TradeOrderProductRepository) UpdateOrderProduct(orderProduct *model.TradeOrderProduct) (*model.TradeOrderProduct, error) {
	updateTime := time.Now()
	orderProduct.UpdateTime = &updateTime
	tx := t.mysqlClient.Model(&model.TradeOrderProduct{}).Updates(orderProduct)
	return orderProduct, tx.Error
}

func NewTradeOrderProductRepository(mysqlClient *gorm.DB, redisClient *redis.Client) ITradeOrderProductRepository {
	return &TradeOrderProductRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
