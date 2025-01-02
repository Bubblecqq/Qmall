package model

import "time"

type TradeOrder struct {
	Id                    int64      `json:"id" gorm:"primaryKey;autoIncrement"`                                                     // 订单ID
	OrderNo               string     `json:"order_no" gorm:"type:varchar(32);not null;unique"`                                       // 订单编号
	UserId                int64      `json:"user_id" gorm:"not null"`                                                                // 用户ID
	TotalAmount           string     `json:"total_amount" gorm:"type:decimal(11,2);not null"`                                        // 总金额
	ShippingAmount        string     `json:"shipping_amount" gorm:"type:decimal(11,2);not null"`                                     // 运费
	DiscountAmount        string     `json:"discount_amount" gorm:"type:decimal(11,2);not null"`                                     // 优惠金额
	PayAmount             string     `json:"pay_amount" gorm:"type:decimal(11,2);not null"`                                          // 实付金额
	RefundAmount          string     `json:"refund_amount" gorm:"type:decimal(11,2);not null"`                                       // 已退款金额
	SubmitTime            time.Time  `json:"submit_time" gorm:"type:datetime;not null"`                                              // 下单时间
	ExpireTime            time.Time  `json:"expire_time" gorm:"type:datetime;not null"`                                              // 失效时间
	AutoReceiveTime       string     `json:"auto_receive_time" gorm:"type:datetime"`                                                 // 自动收货时间
	ReceiveTime           string     `json:"receive_time" gorm:"type:datetime"`                                                      // 收货时间
	AutoPraise            string     `json:"auto_praise" gorm:"type:datetime"`                                                       // 自动好评时间
	AfterSaleDeadlineTime string     `json:"after_sale_deadline_time" gorm:"type:datetime"`                                          // 售后截止时间
	UserMessage           string     `json:"user_message" gorm:"type:varchar(64)"`                                                   // 用户留言
	CancelReason          string     `json:"cancel_reason" gorm:"type:varchar(64)"`                                                  // 取消原因
	OrderSource           int32      `json:"order_source" gorm:"not null"`                                                           // 订单来源（1：未知来源，2：安卓端APP，3：IOS端APP）
	OrderType             int32      `json:"order_type" gorm:"not null"`                                                             // 订单类型（1：普通订单，2：免费订单，3：秒杀订单）
	OrderStatus           int32      `json:"order_status" gorm:"not null"`                                                           // 订单状态（1：待支付，2：已关闭，3：已支付，4：已发货，5：已收货，6：已完成，7：已追评）
	PayType               int32      `json:"pay_type" gorm:"not null"`                                                               // 支付方式（1：未支付，2：微信支付，3：支付宝支付）
	IsPackageFree         int32      `json:"is_package_free" gorm:"not null"`                                                        // 是否包邮（0：不包邮，1：包邮）
	IsAfterSale           int32      `json:"is_after_sale" gorm:"not null"`                                                          // 是否开启售后（0：未开启售后，1：已开启售后）
	IsDisabled            int32      `json:"is_disabled" gorm:"not null"`                                                            // 是否禁用（0：启用，1：禁用）
	IsDeleted             int32      `json:"is_deleted" gorm:"not null"`                                                             // 是否删除（0：未删除，1：已删除）
	CreateTime            time.Time  `json:"create_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`                             // 创建时间
	CreateUser            int64      `json:"create_user" gorm:"not null"`                                                            // 创建人
	UpdateTime            *time.Time `json:"update_time" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新时间
	UpdateUser            int64      `json:"update_user" gorm:""`                                                                    // 更新人
}

func (t *TradeOrder) TableName() string {
	return "trade_order"
}
