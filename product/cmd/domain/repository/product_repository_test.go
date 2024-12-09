package repository

import (
	"QMall/common"
	"QMall/product/cmd/domain/model"
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
	_, list, _ := productRepository.Page(11, 3)

	fmt.Println("list>", list)
	fmt.Println("len?", len(list))
	fmt.Println("count>", productRepository.CountNum())

	var p model.Product
	db.Model(&model.Product{}).Preload("Detail").Preload("PictureList").Find(&p, 15)
	//fmt.Println("product Detail>", p.Detail)
	//fmt.Println("product pictureList>", p.PictureList)

	detail, err := productRepository.ShowProductDetail(15)
	fmt.Println("detail>", *detail)
}
