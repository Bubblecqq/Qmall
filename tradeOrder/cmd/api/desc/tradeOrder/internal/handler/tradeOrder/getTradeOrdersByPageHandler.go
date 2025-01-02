package tradeOrder

import (
	"net/http"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/logic/tradeOrder"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 分页获取订单列表
func GetTradeOrdersByPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageTradeOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tradeOrder.NewGetTradeOrdersByPageLogic(r.Context(), svcCtx)
		resp, err := l.GetTradeOrdersByPage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
