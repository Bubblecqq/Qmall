package repository

import (
	"QMall/common"
	"QMall/shoppingCart/cmd/domain/model"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
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
	//fmt.Println(cartRepository.FindShoppingCart(10))
	//s, _ := cartRepository.InCreaseShoppingCart(12, 3, 1, 233, "云南小黄姜", "https://msb-edu-dev.oss-cn-beijing.aliyuncs.com/mall-product/product4291c68c8352edee6a08255fc8ccbe6fwww755-755.jpg")
	//fmt.Println("s>>>>>>>>>", s)
	price := cartRepository.GetTotalPriceByUserId(19)
	fmt.Println("total_price>", price)

	var shoppingCarts []model.ShoppingCartsProductInfo

	_ = db.Table("shopping_cart sc").
		Select("sc.user_id,sc.product_id,sc.product_sku_id,sc.product_name,sc.product_main_picture,sc.number as quanitiy,psk.sell_price").
		Joins("left join product_sku psk on sc.product_sku_id = psk.id").
		Where("sc.user_id= ?", 19).Find(&shoppingCarts)
	fmt.Println("lens>", len(shoppingCarts))
	fmt.Println("shoppingCarts>", shoppingCarts)

	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var testsp []model.ShoppingCart
	tx := db.Model(&model.ShoppingCart{}).Debug().WithContext(timeout).Find(&testsp)
	if tx.Error != nil {
		fmt.Println("111")
	}
	for _, sp := range testsp {
		fmt.Println("商品名称", sp.ProductName)
		fmt.Println("UserId", sp.UserId)
	}
}
