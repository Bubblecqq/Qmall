package tradeOrder

import (
	"context"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTradeOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成订单
func NewAddTradeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTradeOrderLogic {
	return &AddTradeOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTradeOrderLogic) AddTradeOrder(req *types.AddTradeOrderReq) (resp *types.AddTradeOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
