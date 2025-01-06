package model

import "time"

type TradeOrderProduct struct {
	Id              int64      `grom:"id"`                // 详情ID
	OrderId         int64      `grom:"order_id"`          // 订单ID
	UserId          int64      `grom:"user_id"`           // 用户ID
	ProductId       int64      `grom:"product_id"`        // 商品ID
	ProductSkuId    int64      `grom:"product_sku_id"`    // 商品SKU-ID
	ProductName     string     `grom:"product_name"`      // 商品名称
	ProductImageUrl string     `grom:"product_image_url"` // 商品图片
	SkuDescribe     string     `grom:"sku_describe"`      // SKU规格
	Quantity        int64      `grom:"quantity"`          // 购买数量
	ProductPrice    string     `grom:"product_price"`     // 商品单价
	RealPrice       string     `grom:"real_price"`        // 实际价格
	RealAmount      string     `grom:"real_amount"`       // 实际金额
	ActivityId      int64      `grom:"activity_id"`       // 活动ID
	DetailStatus    int32      `grom:"detail_status"`     // 明细状态（1：正常状态，2：申请售后，3：退款成功，4：退款失败）
	ActivityType    int32      `grom:"activity_type"`     // 活动类型（1：正常购买，2：秒杀）
	CommentStatus   int32      `grom:"comment_status"`    // 评价状态（1：待评价，2：已评价，3：已追评）
	CreateTime      time.Time  `grom:"create_time"`       // 创建时间
	CreateUser      int64      `grom:"create_user"`       // 创建人
	UpdateTime      *time.Time `grom:"update_time"`       // 更新时间
	UpdateUser      int64      `grom:"update_user"`       // 更新人
}

func (t *TradeOrderProduct) TableName() string {
	return "trade_order_product"
}
