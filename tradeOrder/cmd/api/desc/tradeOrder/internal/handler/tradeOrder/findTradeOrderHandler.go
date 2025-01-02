package tradeOrder

import (
	"net/http"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/logic/tradeOrder"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询订单
func FindTradeOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindTradeOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tradeOrder.NewFindTradeOrderLogic(r.Context(), svcCtx)
		resp, err := l.FindTradeOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
