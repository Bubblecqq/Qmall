package model

import "time"

// SecKillUserQuota 秒杀用户限额表
type SecKillUserQuota struct {
	Id         int64      `gorm:"id" json:"id"`
	UserId     int64      `gorm:"user_id" json:"user_id"`
	ProductsId int64      `gorm:"products_id" json:"products_id"`
	Num        int        `gorm:"num" json:"num"`
	KilledNum  int        `gorm:"killed_num" json:"killed_num"`
	CreateTime time.Time  `gorm:"create_time" json:"create_time"`
	UpdateTime *time.Time `gorm:"update_time" json:"update_time"`
}

func (t *SecKillUserQuota) TableName() string {
	return "mall_user_quota_t"
}
