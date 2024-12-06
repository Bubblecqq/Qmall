package productSku

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goZeroDemo4/product/cmd/api/desc/productSku/internal/logic/productSku"
	"goZeroDemo4/product/cmd/api/desc/productSku/internal/svc"
	"goZeroDemo4/product/cmd/api/desc/productSku/internal/types"
)

// 获取商品关联列表
func GetProductSkusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductSkuListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := productSku.NewGetProductSkusLogic(r.Context(), svcCtx)
		resp, err := l.GetProductSkus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
