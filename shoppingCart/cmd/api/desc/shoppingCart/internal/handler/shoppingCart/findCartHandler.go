package shoppingCart

import (
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
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := shoppingCart.NewFindCartLogic(r.Context(), svcCtx)
		resp, err := l.FindCart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
