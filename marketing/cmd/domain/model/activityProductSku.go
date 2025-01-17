package model

import "time"

type ActivityProductSku struct {
	ID                int64      `gorm:"primaryKey;column:id" json:"id"`
	ActivityProductID int64      `gorm:"column:activity_product_id" json:"activity_product_id"`
	ProductID         int64      `gorm:"column:product_id" json:"product_id"`
	ProductSkuID      int64      `gorm:"column:product_sku_id" json:"product_sku_id"`
	Price             float64    `gorm:"column:price" json:"price"`
	OriginalPrice     float64    `gorm:"column:original_price" json:"original_price"`
	Number            int64      `gorm:"column:number" json:"number"`
	Stock             int64      `gorm:"column:stock" json:"stock"`
	CreateUser        int64      `gorm:"column:create_user" json:"create_user"`
	CreateTime        time.Time  `gorm:"column:create_time" json:"create_time"`
	UpdateUser        int64      `gorm:"column:update_user" json:"update_user"`
	UpdateTime        *time.Time `gorm:"column:update_time" json:"update_time"`
	IsDeleted         int32      `gorm:"column:is_deleted" json:"is_deleted"`
}

func (t *ActivityProductSku) TableName() string {
	return "activity_product_sku"
}
