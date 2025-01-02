package repository

import (
	"QMall/tradeOrder/cmd/domain/model"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type ITradeOrderRepository interface {
	FindTradeOrder(id int64, orderNo string) (*model.TradeOrder, error)
	GetTradeOrderById(id int64) (*model.TradeOrder, error)
	GetTradeOrderByOrderNo(orderNo string) (*model.TradeOrder, error)
	AddTradeOrder(req *pb.AddTradeOrderReq) (*model.TradeOrder, error)
	UpdateTradeOrder(req *pb.AddTradeOrderReq) (*model.TradeOrder, error)
	DeleteTradeOrder(id int64) error
	Page(length int32, pageIndex int32) ([]model.TradeOrder, error)
}

type TradeOrderRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (t *TradeOrderRepository) GetTradeOrderById(id int64) (*model.TradeOrder, error) {
	tradeOrder := &model.TradeOrder{}
	err := t.mysqlClient.Model(&model.TradeOrder{}).Find(&tradeOrder, id).Error
	return tradeOrder, err
}

func (t *TradeOrderRepository) GetTradeOrderByOrderNo(orderNo string) (*model.TradeOrder, error) {
	tradeOrder := &model.TradeOrder{}
	err := t.mysqlClient.Model(&model.TradeOrder{}).Where("order_no = ?", orderNo).Find(&tradeOrder).Error
	return tradeOrder, err
}

func (t *TradeOrderRepository) AddTradeOrder(req *pb.AddTradeOrderReq) (*model.TradeOrder, error) {
	return nil, nil
}

func (t *TradeOrderRepository) UpdateTradeOrder(req *pb.AddTradeOrderReq) (*model.TradeOrder, error) {
	return nil, nil
}

func (t *TradeOrderRepository) DeleteTradeOrder(id int64) error {
	tradeOrderById, err := t.GetTradeOrderById(id)
	if err != nil {
		return err
	}
	err = t.mysqlClient.Model(&model.TradeOrder{}).Delete(&tradeOrderById).Error
	return err
}

func (t *TradeOrderRepository) Page(length int32, pageIndex int32) ([]model.TradeOrder, error) {

	tradeOrders := make([]model.TradeOrder, length)
	var err error
	if length > 0 && pageIndex > 0 {
		err = t.mysqlClient.Model(&model.TradeOrder{}).Limit(int(length)).Offset(int(pageIndex-1) * int(length)).Find(&tradeOrders).Error
	}

	return tradeOrders, err
}

func (t *TradeOrderRepository) FindTradeOrder(id int64, orderNo string) (*model.TradeOrder, error) {
	tradeOrder := &model.TradeOrder{}
	err := t.mysqlClient.Model(&model.TradeOrder{}).Where("id= ? or order_no= ?", id, orderNo).Find(&tradeOrder).Error
	return tradeOrder, err
}

// GenerateOrderNo 生产 订单号   Y2022 06 27 11 00 53 948 97 103564
//
//	年    月 日 时  分 秒 毫秒 ID  随机数
func GenerateOrderNo(time2 time.Time, userId int64) string {
	var orderNo string
	tempNum := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	orderNo = "Y" + time2.Format("20060102150405.000") + strconv.Itoa(int(userId)) + tempNum
	orderNo = strings.Replace(orderNo, ".", "", -1)
	return orderNo
}

func NewTradeOrderRepository(mysqlClient *gorm.DB, redisClient *redis.Client) ITradeOrderRepository {
	return &TradeOrderRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
