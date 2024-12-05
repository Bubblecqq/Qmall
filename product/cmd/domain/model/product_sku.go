package model

type ProductSku struct {
	BaseModel
	//SkuId               int32   `gorm:"column:id"`
	Name                string  `gorm:"column:name"`
	ProductId           int32   `gorm:"column:product_id;default:1"`
	AttributeSymbolList string  `gorm:"column:attribute_symbol_list"`
	SellPrice           float64 `gorm:"column:sell_price"`
	CostPrice           float64 `gorm:"column:cost_price;default:1"`
	Stock               int32   `gorm:"column:stock"`
	StockWarn           int     `gorm:"column:stock_warn"`
}
