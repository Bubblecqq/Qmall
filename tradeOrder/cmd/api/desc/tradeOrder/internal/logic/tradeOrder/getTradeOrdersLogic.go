package tradeOrder

import (
	"context"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTradeOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取订单列表
func NewGetTradeOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTradeOrdersLogic {
	return &GetTradeOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTradeOrdersLogic) GetTradeOrders(req *types.GetTradeOrderListReq) (resp *types.GetTradeOrderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
