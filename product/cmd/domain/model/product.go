package model

type Product struct {
	BaseModel
	Name              string  `gorm:"column:name"`
	ProductType       int32   `gorm:"column:product_type;default:1"`
	CategoryId        int32   `gorm:"column:category_id"`
	StartingPrice     float64 `gorm:"column:starting_price"`
	TotalStock        int32   `gorm:"column:total_stock;default:'1234'"`
	MainPicture       string  `gorm:"column:main_picture;default:1"`
	RemoteAreaPostage float64 `gorm:"column:remote_area_postage"`
	SingleBuyLimit    int32   `gorm:"column:single_buy_limit"`
	Remark            string  `gorm:"column:remark;default:'1'"`
}
