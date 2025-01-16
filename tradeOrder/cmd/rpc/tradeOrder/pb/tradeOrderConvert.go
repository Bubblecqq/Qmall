package pb

import (
	"QMall/tradeOrder/cmd/domain/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelTradeOrderConvertPb(order *model.TradeOrder) *TradeOrder {
	return &TradeOrder{
		Id:              order.Id,
		OrderNo:         order.OrderNo,
		UserId:          order.UserId,
		TotalAmount:     order.TotalAmount,
		ShippingAmount:  order.ShippingAmount,
		DiscountAmount:  order.DiscountAmount,
		PayAmount:       order.PayAmount,
		RefundAmount:    order.RefundAmount,
		SubmitTime:      timestamppb.New(order.SubmitTime),
		ExpireTime:      timestamppb.New(order.ExpireTime),
		OrderStatus:     order.OrderStatus,
		CancelReason:    order.CancelReason,
		CreateUser:      order.CreateUser,
		UpdateUser:      order.UpdateUser,
		CreateTime:      timestamppb.New(order.CreateTime),
		IsDeleted:       order.IsDeleted == 1,
		OrderType:       order.OrderType,
		OrderSource:     order.OrderSource,
		UserMessage:     order.UserMessage,
		IsPackageFree:   order.IsPackageFree,
		IsAfterSale:     order.IsAfterSale,
		IsDisabled:      order.IsDisabled,
		AutoReceiveTime: timestamppb.New(order.AutoReceiveTime),
		PayType:         order.PayType,
		//ReceiveTime:     timestamppb.New(order.ReceiveTime),
	}
}

func PbTradeOrderConvertModel(order *TradeOrder) *model.TradeOrder {
	return &model.TradeOrder{
		Id:             order.Id,
		OrderNo:        order.OrderNo,
		UserId:         order.UserId,
		TotalAmount:    order.TotalAmount,
		ShippingAmount: order.ShippingAmount,
		DiscountAmount: order.DiscountAmount,
		PayAmount:      order.PayAmount,
		RefundAmount:   order.RefundAmount,
		SubmitTime:     order.SubmitTime.AsTime(),
		ExpireTime:     order.ExpireTime.AsTime(),
		OrderStatus:    order.OrderStatus,
		CancelReason:   order.CancelReason,
		CreateUser:     order.CreateUser,
		UpdateUser:     order.UpdateUser,
		CreateTime:     order.CreateTime.AsTime(),
		IsDeleted: func(isDeleted bool) int32 {
			if isDeleted {
				return 1
			}
			return 0
		}(order.IsDeleted),
		OrderSource:   order.OrderSource,
		OrderType:     order.OrderType,
		IsAfterSale:   order.IsAfterSale,
		IsDisabled:    order.IsDisabled,
		IsPackageFree: order.IsPackageFree,
		//UpdateTime: &(order.UpdateTime.AsTime()),
		UserMessage: order.UserMessage,
		PayType:     order.PayType,
	}
}
