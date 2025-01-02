package tradeOrder

import (
	"net/http"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/logic/tradeOrder"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 生成订单
func AddTradeOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTradeOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tradeOrder.NewAddTradeOrderLogic(r.Context(), svcCtx)
		resp, err := l.AddTradeOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
