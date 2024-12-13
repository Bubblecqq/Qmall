package product

import (
	"QMall/common"
	"net/http"

	"QMall/product/cmd/api/desc/product/internal/logic/product"
	"QMall/product/cmd/api/desc/product/internal/svc"
	"QMall/product/cmd/api/desc/product/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取商品关联列表
func GetProductSkusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductSkuListReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			common.RespOk(r.Context(), w, "", "失败！", common.TOKEN_FAIL)
			return
		}

		l := product.NewGetProductSkusLogic(r.Context(), svcCtx)
		resp, err := l.GetProductSkus(&req)
		if err != nil {
			common.RespOk(r.Context(), w, resp, "失败！", common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			common.RespOk(r.Context(), w, resp, "获取成功！", common.SUCCESS)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
