package tradeOrder

import (
	"net/http"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/logic/tradeOrder"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取订单列表
func GetTradeOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTradeOrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tradeOrder.NewGetTradeOrdersLogic(r.Context(), svcCtx)
		resp, err := l.GetTradeOrders(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
