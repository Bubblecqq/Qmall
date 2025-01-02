package tradeOrder

import (
	"context"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTradeOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改订单
func NewUpdateTradeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTradeOrderLogic {
	return &UpdateTradeOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTradeOrderLogic) UpdateTradeOrder(req *types.AddTradeOrderReq) (resp *types.AddTradeOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
