package tradeOrder

import (
	"context"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTradeOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询订单
func NewFindTradeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTradeOrderLogic {
	return &FindTradeOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTradeOrderLogic) FindTradeOrder(req *types.FindTradeOrderReq) (resp *types.FindTradeOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
