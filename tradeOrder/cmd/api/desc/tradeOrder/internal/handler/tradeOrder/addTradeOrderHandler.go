package tradeOrder

import (
	"QMall/common"
	"net/http"

	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/logic/tradeOrder"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// AddTradeOrderHandler 生成订单
func AddTradeOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTradeOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			common.RespOk(r.Context(), w, "", "请求失败！", common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tradeOrder.NewAddTradeOrderLogic(r.Context(), svcCtx)

		resp, err := l.AddTradeOrder(&req)
		if err != nil {
			common.RespOk(r.Context(), w, "", err.Error(), common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			common.RespOk(r.Context(), w, resp, "获取成功！", common.SUCCESS)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
