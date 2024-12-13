package shoppingCart

import (
	"QMall/common"
	"net/http"

	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/logic/shoppingCart"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/svc"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查看购物车
func FindCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindShoppingCartReq
		if err := httpx.Parse(r, &req); err != nil {
			common.RespOk(r.Context(), w, "", "失败！", common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := shoppingCart.NewFindCartLogic(r.Context(), svcCtx)
		resp, err := l.FindCart(&req)
		if err != nil {
			common.RespOk(r.Context(), w, resp, "失败！", common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			common.RespOk(r.Context(), w, resp, "获取成功！", common.SUCCESS)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
