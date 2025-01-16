package Repository

import (
	"QMall/logistics/cmd/domain/model"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ILogisticsRepository interface {
	GetLogisticsByOrderId(orderId int64) (*model.TradeOrderLogistics, error)
	GetLogisticsById(Id int64) (*model.TradeOrderLogistics, error)
	AddLogistics() error
	DeleteLogistics(id int64) error
}

type LogisticsRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (l *LogisticsRepository) GetLogisticsByOrderId(orderId int64) (*model.TradeOrderLogistics, error) {
	logistics := &model.TradeOrderLogistics{}

	tx := l.mysqlClient.Model(&model.TradeOrderLogistics{}).Where("order_id=?", orderId).Find(&logistics)
	if tx.Error != nil {
		fmt.Printf("订单Id：%v的物流不存在！原因见：%v\n", orderId, tx.Error)
	}
	return logistics, tx.Error
}

func (l *LogisticsRepository) GetLogisticsById(Id int64) (*model.TradeOrderLogistics, error) {
	logistics := &model.TradeOrderLogistics{}

	tx := l.mysqlClient.Model(&model.TradeOrderLogistics{}).Find(&logistics, Id)
	if tx.Error != nil {
		fmt.Printf("物流订单Id：%v不存在！原因见：%v\n", Id, tx.Error)
	}
	return logistics, tx.Error
}

func (l *LogisticsRepository) AddLogistics() error {
	//TODO implement me
	panic("implement me")
}

func (l *LogisticsRepository) DeleteLogistics(id int64) error {
	tx := l.mysqlClient.Model(&model.TradeOrderLogistics{}).Delete(id)
	if tx.Error != nil {
		fmt.Printf("物流订单Id：%v删除失败！原因见：%v\n", id, tx.Error)
	}
	return tx.Error
}

func NewLogisticsRepository(mysqlClient *gorm.DB, redisClient *redis.Client) ILogisticsRepository {
	return &LogisticsRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
