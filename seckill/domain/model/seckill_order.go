package model

import "time"

// SecKillOrder
type SecKillOrder struct {
	Id          int64      `gorm:"id" json:"id"`                     // ID
	Seller      int64      `gorm:"seller" json:"seller"`             // 买方ID
	Buyer       int64      `gorm:"buyer" json:"buyer"`               // 卖方ID
	ProductsId  int64      `gorm:"products_id" json:"products_id"`   // 商品ID
	ProductsNum string     `gorm:"products_num" json:"products_num"` // 商品编号
	OrderNum    string     `gorm:"order_num" json:"order_num"`       // 订单号
	Price       string     `gorm:"price" json:"price"`               // 金额
	Status      int32      `gorm:"status" json:"status"`             // 状态
	CreateTime  time.Time  `gorm:"create_time" json:"create_time"`   // 创建时间
	UpdateTime  *time.Time `gorm:"update_time" json:"update_time"`   // 修改时间
}

func (t *SecKillOrder) TableName() string {
	return "mall_order_t"
}
