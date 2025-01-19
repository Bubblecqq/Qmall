package repository

import (
	"QMall/tradeOrder/cmd/domain/model"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	GetTradeOrders() ([]model.TradeOrder, error)
	GetTradeOrderByUserId(id, userId int64) (*model.TradeOrder, error)
	//GetTotalPrice(userId int64) float64
	UpdateTradeOrder2(in *pb.UpdateTradeOrderReq) (*model.TradeOrder, error)

	//GetTradeOrderByOrderNo(orderNo string) (*model.TradeOrder, error)
}

type TradeOrderRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (t *TradeOrderRepository) UpdateTradeOrder2(in *pb.UpdateTradeOrderReq) (*model.TradeOrder, error) {

	var order *model.TradeOrder
	tx := t.mysqlClient.Model(&model.TradeOrder{}).Begin()
	find := tx.Model(&model.TradeOrder{}).Where("user_id=?", in.UserId).Find(&order, in.OrderId)
	if find.Error != nil {
		fmt.Printf("用户Id：%v的订单Id：%v查询不存在！原因见：%v\n", in.UserId, in.OrderId, find.Error)
		return nil, find.Error
	}
	order.PayType = in.PayType
	order.IsAfterSale = in.IsAfterSale
	if in.IsAfterSale == 1 {
		order.CancelReason = in.CancelReason
	}
	if in.IsRefund == 1 {
		//退款中
		// 商家方点确认后需要库存补偿
		order.OrderStatus = 8
	}
	tb := tx.Model(&model.TradeOrder{}).Updates(&order)
	if tb.Error != nil {
		fmt.Printf("订单更新操作失败，订单ID：%v,失败原因：%v\n", in.OrderId, tb.Error)
		tx.Rollback()
		return nil, tb.Error
	}
	if tx.Commit().Error != nil {
		fmt.Printf("事务提交失败，错误信息: %v", tx.Commit().Error)
		return nil, errors.New("事务提交失败！")
	}
	fmt.Println("TradeOrder Repository Updated Trade Order>>>>>>>")
	return order, nil
}

func (t *TradeOrderRepository) GetTradeOrderByUserId(id, userId int64) (*model.TradeOrder, error) {
	order := &model.TradeOrder{}
	tx := t.mysqlClient.Model(&model.TradeOrder{}).Where("user_id=?", userId).Find(&order, id)
	if tx.Error != nil {
		fmt.Printf("当前用户Id：%v的订单Id：%v不存在！详细原因见：%v\n", userId, id, tx.Error)
	}
	return order, tx.Error
}

//func (t *TradeOrderRepository) GetTotalPrice(userId int64) float64 {
//	var totalPrice float64
//	tx := t.mysqlClient.Model(&model2.ShoppingCart{}).
//		Select("Sum(shopping_cart.number * mall.product.product_sku.sell_price) as total_price").
//		Joins("left join mall.product.product_sku on shopping_cart.product_sku_id = mall.product.product_sku.id").
//		Where("shopping_cart.user_id = ?", userId).
//		Raw("").
//		Scan(&totalPrice)
//
//	if tx.Error != nil {
//		fmt.Printf("当前用户Id:%v，购物车中不存在任意商品！\n", userId)
//		return 0
//	}
//
//	return totalPrice
//}

func (t *TradeOrderRepository) GetTradeOrders() ([]model.TradeOrder, error) {
	var orders []model.TradeOrder
	tx := t.mysqlClient.Model(&model.TradeOrder{}).Find(&orders)
	return orders, tx.Error
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
	fmt.Printf("开始创建订单，请求参数：%v\n", req)
	trade := pb.PbTradeOrderConvertModel(req.TradeOrder)

	//err := SwapToStruct(req.TradeOrder, trade)

	//if err != nil {
	//	fmt.Printf("Swap struct err with:%v\n", err)
	//}

	fmt.Println("SwapToStruct trade:", trade)
	now := time.Now()
	trade.CreateTime = now
	trade.SubmitTime = now
	tp, _ := time.ParseDuration("30m")
	trade.ExpireTime = now.Add(tp)
	// 订单类型 1:普通 2：免费 3：秒杀
	trade.OrderType = 1
	trade.OrderStatus = 1
	trade.PayType = 1
	trade.IsAfterSale = 0
	trade.IsDisabled = 0
	trade.IsDeleted = 0
	// 15天自动收货
	trade.AutoReceiveTime = now.Add(15 * 24 * time.Hour)
	trade.ReceiveTime = &now
	trade.AutoPraise = &now
	trade.AfterSaleDeadlineTime = &now
	trade.OrderNo = GenerateOrderNo(now, trade.UserId)
	tb := t.mysqlClient.Model(&model.TradeOrder{}).Create(&trade)
	fmt.Printf("订单已生成！，订单编号：%v\n", trade.OrderNo)
	return trade, tb.Error
}

func (t *TradeOrderRepository) UpdateTradeOrder(req *pb.AddTradeOrderReq) (*model.TradeOrder, error) {
	if req.TradeOrder.OrderStatus < 1 || req.TradeOrder.OrderStatus > 8 {
		return nil, errors.New("无效订单状态修改值")
	}
	trade := &model.TradeOrder{}
	trade.Id = req.TradeOrder.Id
	trade.OrderStatus = req.TradeOrder.OrderStatus
	if req.TradeOrder.IsDeleted {
		trade.IsDeleted = 1
	} else {
		trade.IsDeleted = 0
	}
	//  1：待支付，2：已关闭，3：已支付，4：已发货，5：已收货，6：已完成，7：已追评
	if req.TradeOrder.OrderStatus == 2 {
		trade.CancelReason = req.TradeOrder.CancelReason
	}

	var updateTime *time.Time
	*updateTime = time.Now()
	trade.UpdateTime = updateTime
	fmt.Println("开始执行订单修改事务，订单ID：", req.TradeOrder.Id)
	tx := t.mysqlClient.Begin()
	err := tx.Model(&model.TradeOrder{}).
		Where("id= ?", req.TradeOrder.Id).Clauses(clause.Locking{Strength: "UPDATE"}).Error
	if err != nil {
		fmt.Printf("设置行锁失败，订单ID：%v,失败原因：%v\n", req.TradeOrder.Id, err)
		tx.Rollback()
		return nil, err
	}

	tb := tx.Model(&model.TradeOrder{}).
		Where("id= ?", req.TradeOrder.Id).
		Updates(&trade)
	if tb.Error != nil {
		fmt.Printf("订单更新操作失败，订单ID：%v,失败原因：%v\n", req.TradeOrder.Id, tb.Error)
		tx.Rollback()
		return nil, tb.Error
	}
	if tx.Commit().Error != nil {
		fmt.Printf("事务提交失败，错误信息: %v", tx.Commit().Error)
		return nil, errors.New("事务提交失败！")
	}
	fmt.Println("TradeOrder Repository Updated Trade Order>>>>>>>")
	return trade, nil
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

func SwapToStruct(req, target interface{}) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, target)
	return err
}

func NewTradeOrderRepository(mysqlClient *gorm.DB, redisClient *redis.Client) ITradeOrderRepository {
	return &TradeOrderRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
