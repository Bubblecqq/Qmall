package model

import "time"

type ActivityProduct struct {
	ID                   int64      `gorm:"primaryKey;column:id" json:"id"`
	ActivityTimeID       int64      `gorm:"column:activity_time_id" json:"activity_time_id"`
	ProductID            int64      `gorm:"column:product_id" json:"product_id"`
	ProductName          string     `gorm:"column:product_name" json:"product_name"`
	ProductMainPic       string     `gorm:"column:product_main_picture" json:"product_main_picture"`
	ProductsNum          string     `gorm:"column:products_num" json:"products_num"`
	ProductStartingPrice float64    `gorm:"column:product_starting_price" json:"product_starting_price"`
	CategoryID           int64      `gorm:"column:category_id" json:"category_id"`
	CreateUser           int64      `gorm:"column:create_user" json:"create_user"`
	CreateTime           time.Time  `gorm:"column:create_time" json:"create_time"`
	UpdateUser           int64      `gorm:"column:update_user" json:"update_user"`
	UpdateTime           *time.Time `gorm:"column:update_time" json:"update_time"`
	IsDeleted            int32      `gorm:"column:is_deleted" json:"is_deleted"`
}

func (t *ActivityProduct) TableName() string {
	return "activity_product"
}
