package types

import "QMall/tradeOrder/cmd/rpc/tradeOrder/pb"

func TradeOrderPbConvertTypes(tradeOrder *pb.TradeOrder) TradeOrder {
	return TradeOrder{
		Id:             tradeOrder.Id,
		OrderNo:        tradeOrder.OrderNo,
		UserId:         tradeOrder.UserId,
		TotalAmount:    tradeOrder.TotalAmount,
		ShippingAmount: tradeOrder.ShippingAmount,
		DiscountAmount: tradeOrder.DiscountAmount,
		PayAmount:      tradeOrder.PayAmount,
		RefundAmount:   tradeOrder.RefundAmount,
		SubmitTime:     tradeOrder.SubmitTime.String(),
		ExpireTime:     tradeOrder.ExpireTime.String(),
		OrderStatus:    tradeOrder.OrderStatus,
		CancelReason:   tradeOrder.CancelReason,
		CreateUser:     tradeOrder.CreateUser,
		UpdateUser:     tradeOrder.UpdateUser,
		CreateTime:     tradeOrder.CreateTime.String(),
		IsDeleted: func(isDeleted bool) int32 {
			if isDeleted {
				return 1
			}
			return 0
		}(tradeOrder.IsDeleted),
	}
}
