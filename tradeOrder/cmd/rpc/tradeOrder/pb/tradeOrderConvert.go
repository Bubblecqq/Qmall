package pb

import (
	"QMall/tradeOrder/cmd/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelTradeOrderConvertPb(order *model.TradeOrder) *TradeOrder {
	return &TradeOrder{
		Id:             order.Id,
		OrderNo:        order.OrderNo,
		UserId:         order.UserId,
		TotalAmount:    order.TotalAmount,
		ShippingAmount: order.ShippingAmount,
		DiscountAmount: order.DiscountAmount,
		PayAmount:      order.PayAmount,
		RefundAmount:   order.RefundAmount,
		SubmitTime:     timestamppb.New(order.SubmitTime),
		ExpireTime:     timestamppb.New(order.ExpireTime),
		OrderStatus:    order.OrderStatus,
		CancelReason:   order.CancelReason,
		CreateUser:     order.CreateUser,
		UpdateUser:     order.UpdateUser,
		CreateTime:     timestamppb.New(order.CreateTime),
		IsDeleted:      order.IsDeleted == 1,
	}
}
