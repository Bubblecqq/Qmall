package model

import "time"

// SecKillQuota 秒杀限额表对应的结构体
type SecKillQuota struct {
	Id                int64      `gorm:"id" json:"id"`
	ProductsId        int64      `gorm:"products_id" json:"products_id"`
	SecKillProductsId int64      `gorm:"column:seckill_products_id" json:"seckill_products_id"`
	Num               int64      `gorm:"num" json:"num"`
	CreateTime        time.Time  `gorm:"create_time" json:"create_time"`
	UpdateTime        *time.Time `gorm:"update_time" json:"update_time"`
}

func (t *SecKillQuota) TableName() string {
	return "mall_quota_t"
}
