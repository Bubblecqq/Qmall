package model

import "time"

// SecKillRecord 秒杀记录表
type SecKillRecord struct {
	Id         int64      `gorm:"id" json:"id"`                   // ID
	UserId     int64      `gorm:"user_id" json:"user_id"`         // 用户ID
	ProductsId int64      `gorm:"products_id" json:"products_id"` // 商品ID
	SecNum     string     `gorm:"sec_num" json:"sec_num"`         // 秒杀号
	OrderNum   string     `gorm:"order_num" json:"order_num"`     // 订单号
	Price      string     `gorm:"price" json:"price"`             // 金额
	Status     int32      `gorm:"status" json:"status"`           // 状态
	CreateTime time.Time  `gorm:"create_time" json:"create_time"` // 创建时间
	UpdateTime *time.Time `gorm:"update_time" json:"update_time"` // 修改时间
}

func (t *SecKillRecord) TableName() string {
	return "mall_seckill_record_t"
}
