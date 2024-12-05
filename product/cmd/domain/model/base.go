package model

import "time"

// BaseModel 结构体包含一些通用的字段，用于被其他模型结构体继承
type BaseModel struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	//IsDeleted  int32     `gorm:"column:is_deleted;default:0"`
	IsEnabled  int32  `gorm:"column:is_enable;default:1"`
	CreateUser string `gorm:"column:create_user"`
	UpdateUser string `gorm:"column:update_user"`
}
