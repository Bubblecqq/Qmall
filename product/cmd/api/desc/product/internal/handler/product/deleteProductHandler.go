package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goZeroDemo4/product/cmd/api/desc/product/internal/logic/product"
	"goZeroDemo4/product/cmd/api/desc/product/internal/svc"
	"goZeroDemo4/product/cmd/api/desc/product/internal/types"
)

// 根据id删除商品
func DeleteProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewDeleteProductLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
