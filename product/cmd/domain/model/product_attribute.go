package model

import "time"

type ProductAttribute struct {
	Id                      int64     `gorm:"column:id"`
	Name                    string    `gorm:"column:name"`
	ProductAttributeGroupId int64     `gorm:"column:product_attribute_group_id"`
	Symbol                  int32     `gorm:"column:symbol"`
	ProductId               int64     `gorm:"column:product_id"`
	Sort                    int       `gorm:"column:sort"`
	CreateTime              time.Time `gorm:"column:create_time"`
	UpdateTime              time.Time `gorm:"column:update_time"`
	CreateUser              string    `gorm:"column:create_user"`
	UpdateUser              string    `gorm:"column:update_user"`
}

func (p *ProductAttribute) TableName() string {
	return "product_attribute"
}
