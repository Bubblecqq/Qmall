package model

type ActivityInfo struct {
	ActivityName         string  `gorm:"column:activity_name" json:"activity_name"`
	ProductName          string  `gorm:"column:product_name" json:"product_name"`
	ProductStartingPrice float64 `gorm:"column:product_starting_price" json:"product_starting_price"`
	CategoryId           int64   `gorm:"column:category_id" json:"category_id"`
	ProductId            int64   `gorm:"column:product_id" json:"product_id"`
	ProductsNum          string  `gorm:"column:products_num" json:"products_num"`
	ActivityTimeId       int64   `gorm:"column:activity_time_id" json:"activity_time_id"`
	StartTime            string  `gorm:"column:start_time" json:"start_time"`
	EndTime              string  `gorm:"column:end_time" json:"end_time"`
	IsOnline             int32   `gorm:"column:is_online" json:"is_online"`
}
