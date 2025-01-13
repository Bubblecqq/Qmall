package repository

import (
	"QMall/common"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestCRUD(t *testing.T) {
	dsn := fmt.Sprintf(common.MysqlConnect, "root", "000000", "192.168.23.233:3306", "mall_trade", "utf8mb4")
	var dialector = mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	})
	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	orderRepository := NewTradeOrderRepository(db, nil)
	order, _ := orderRepository.GetTradeOrderById(1)
	fmt.Println("order>", order)
	//totalPrice := orderRepository.GetTotalPrice(19)
	//fmt.Println("total_price>", totalPrice)
}
