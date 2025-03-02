package product

import (
	"QMall/common"
	"net/http"

	"QMall/product/cmd/api/desc/product/internal/logic/product"
	"QMall/product/cmd/api/desc/product/internal/svc"
	"QMall/product/cmd/api/desc/product/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 根据id获取商品
func GetProductByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			common.RespOk(r.Context(), w, "", "失败！", common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewGetProductByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetProductById(&req)
		if err != nil {
			common.RespOk(r.Context(), w, resp, "失败！", common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			//httpx.OkJsonCtx(r.Context(), w, resp)
			common.RespOk(r.Context(), w, resp, "获取成功！", common.SUCCESS)
		}
	}
}
