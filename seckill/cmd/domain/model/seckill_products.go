package model

import "time"

// SecKillProducts 秒杀商品表
type SecKillProducts struct {
	Id           int64      `gorm:"id" json:"id"`
	ProductsNum  string     `gorm:"products_num" json:"products_num"` // 商品编号
	ProductsId   int64      `gorm:"products_id" json:"products_id"`
	ProductsName string     `gorm:"products_name" json:"products_name"` // 商品名字
	Price        float64    `gorm:"price" json:"price"`                 // 价格
	PictureUrl   string     `gorm:"picture_url" json:"picture_url"`     // 商品图片
	Seller       int64      `gorm:"seller" json:"seller"`               // 卖家ID
	CreateTime   time.Time  `gorm:"create_time" json:"create_time"`     // 创建时间
	UpdateTime   *time.Time `gorm:"update_time" json:"update_time"`     // 修改时间
}

func (t *SecKillProducts) TableName() string {
	return "mall_products_t"
}
