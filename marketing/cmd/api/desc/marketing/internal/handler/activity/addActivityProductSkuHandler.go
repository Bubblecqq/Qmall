package activity

import (
	"QMall/common"
	"net/http"

	"QMall/marketing/cmd/api/desc/marketing/internal/logic/activity"
	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加活动商品库存信息
func AddActivityProductSkuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddActivityProductSkuReq
		if err := httpx.Parse(r, &req); err != nil {
			common.RespOk(r.Context(), w, "", "请求失败！", common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := activity.NewAddActivityProductSkuLogic(r.Context(), svcCtx)
		resp, err := l.AddActivityProductSku(&req)
		if err != nil {
			common.RespOk(r.Context(), w, "", err.Error(), common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			common.RespOk(r.Context(), w, resp, "获取成功！", common.SUCCESS)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
