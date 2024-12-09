package model

import "time"

type ProductPicture struct {
	Id         int64     `gorm:"column:id"`
	ProductId  int64     `gorm:"column:product_id"`
	Picture    string    `gorm:"column:picture"`
	Sort       int       `gorm:"column:sort"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	CreateUser string    `gorm:"column:create_user"`
	UpdateUser string    `gorm:"column:update_user"`
}

func (p *ProductPicture) TableName() string {
	return "product_picture"
}
