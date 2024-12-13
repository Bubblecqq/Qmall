package shoppingCart

import (
	"QMall/common"
	"net/http"

	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/logic/shoppingCart"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/svc"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加购物车
func IncreaseCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IncreaseShoppingCartReq
		r.Header.Set("uuid", "23333")

		if err := httpx.Parse(r, &req); err != nil {
			common.RespOk(r.Context(), w, "", "请求失败！", common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := shoppingCart.NewIncreaseCartLogic(r.Context(), svcCtx)
		resp, err := l.IncreaseCart(&req)
		if err != nil {
			common.RespOk(r.Context(), w, "", err.Error(), common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			common.RespOk(r.Context(), w, resp, "获取成功！", common.SUCCESS)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
