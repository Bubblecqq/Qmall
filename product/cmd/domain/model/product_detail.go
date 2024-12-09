package model

type ProductDetail struct {
	ProductId int64  `gorm:"column:product_id"`
	Detail    string `gorm:"column:detail"`
}

func (p *ProductDetail) TableName() string {
	return "product_detail"
}
