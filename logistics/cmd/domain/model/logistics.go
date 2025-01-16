package model

import "time"

// TradeOrderLogistics 订单物流信息结构体
type TradeOrderLogistics struct {
	Id               int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrderId          int64      `gorm:"column:order_id;uniqueIndex" json:"order_id"`
	RecipientName    string     `gorm:"column:recipient_name" json:"recipient_name"`
	RecipientPhone   string     `gorm:"column:recipient_phone" json:"recipient_phone"`
	RecipientAddress string     `gorm:"column:recipient_address" json:"recipient_address"`
	Province         string     `gorm:"column:province" json:"province"`
	ProvinceCode     string     `gorm:"column:province_code" json:"province_code"`
	City             string     `gorm:"column:city" json:"city"`
	CityCode         string     `gorm:"column:city_code" json:"city_code"`
	Area             string     `gorm:"column:area" json:"area"`
	AreaCode         string     `gorm:"column:area_code" json:"area_code"`
	DetailAddress    string     `gorm:"column:detail_address" json:"detail_address"`
	CompanyCode      string     `gorm:"column:company_code" json:"company_code"`
	CompanyName      string     `gorm:"column:company_name" json:"company_name"`
	TrackingNo       string     `gorm:"column:tracking_no" json:"tracking_no"`
	LogisticsApi     string     `gorm:"column:logistics_api" json:"logistics_api"`
	LogisticsData    string     `gorm:"column:logistics_data" json:"logistics_data"`
	IsSubscribe      bool       `gorm:"column:is_subscribe" json:"is_subscribe"`
	LogisticsStatus  int        `gorm:"column:logistics_status" json:"logistics_status"`
	CreateTime       time.Time  `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"create_time"`
	CreateUser       int64      `gorm:"column:create_user" json:"create_user"`
	UpdateTime       *time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"update_time"`
	UpdateUser       int64      `gorm:"column:update_user" json:"update_user"`
}

func (t *TradeOrderLogistics) TableName() string {
	return "trade_order_logistics"
}
