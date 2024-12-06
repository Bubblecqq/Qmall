package repository

import (
	"QMall/common"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

var mysqlConnect = "root:000000@tcp(192.168.23.233:3306)/user_center?charset=utf8mb4&parseTime=True&loc=Local"

func TestProduct(t *testing.T) {
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
	productRepository := NewProductRepository(db, nil)
	list := productRepository.GetProductSkuList()

	fmt.Println("list>", list)
	fmt.Println("len?", len(list))
}
