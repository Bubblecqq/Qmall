package model

import "time"

type ShoppingCart struct {
	Id                 int64      `gorm:"column:id;primaryKey;autoIncrement"`
	UserId             int64      `gorm:"column:user_id"`
	ProductId          int64      `gorm:"column:product_id"`
	ProductSkuId       int64      `gorm:"column:product_sku_id"`
	ProductName        string     `gorm:"column:product_name"`
	ProductMainPicture string     `gorm:"column:product_main_picture"`
	Number             int64      `gorm:"column:number"`
	CreateTime         time.Time  `gorm:"column:create_time"`
	UpdateTime         *time.Time `gorm:"column:update_time"`
	CreateUser         int64      `gorm:"column:create_user"`
	UpdateUser         int64      `gorm:"column:update_user"`
	IsDeleted          int32      `gorm:"column:is_deleted;default:0"`
}

func (p *ShoppingCart) TableName() string {
	return "shopping_cart"
}
