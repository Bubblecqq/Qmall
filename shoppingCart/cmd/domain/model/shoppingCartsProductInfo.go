package model

type ShoppingCartsProductInfo struct {
	ProductID      int64   `gorm:"column:product_id"`
	ProductName    string  `gorm:"column:product_name"`
	ProductSkuID   int64   `gorm:"column:product_sku_id"`
	SellPrice      float64 `gorm:"column:sell_price"`
	Quantity       int64   `gorm:"column:quantity"`
	ProductMainPic string  `gorm:"column:product_main_picture"`
	SkuDescribe    string  `gorm:"column:name"`
}
