package repository

import (
	"QMall/common"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestFind(t *testing.T) {
	dsn := fmt.Sprintf(common.MysqlConnect, "root", "000000", "192.168.23.233:3306", "mall_product", "utf8mb4")
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
	cartRepository := NewShoppingCartRepository(db, nil)
	fmt.Println(cartRepository.FindShoppingCart(10))
	s, _ := cartRepository.InCreaseShoppingCart(12, 3, 1, 233, "云南小黄姜", "https://msb-edu-dev.oss-cn-beijing.aliyuncs.com/mall-product/product4291c68c8352edee6a08255fc8ccbe6fwww755-755.jpg")
	fmt.Println("s>>>>>>>>>", s)
}
