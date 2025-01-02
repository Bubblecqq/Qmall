package tradeOrder

import (
	"context"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTradeOrdersByPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取订单列表
func NewGetTradeOrdersByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTradeOrdersByPageLogic {
	return &GetTradeOrdersByPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTradeOrdersByPageLogic) GetTradeOrdersByPage(req *types.PageTradeOrderReq) (resp *types.PageTradeOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
