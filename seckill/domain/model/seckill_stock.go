package model

// SeckillStockTInfo
type SeckillStock struct {
	Id         int64  `gorm:"id" json:"id"`                   // ID
	ProductsId int64  `gorm:"products_id" json:"products_id"` // 商品ID
	Stock      int64  `gorm:"stock" json:"stock"`             // 库存大小
	CreateTime string `gorm:"create_time" json:"create_time"` // 创建时间
	UpdateTime string `gorm:"update_time" json:"update_time"` // 修改时间
}

func (t *SeckillStock) TableName() string {
	return "mall_seckill_stock_t"
}
