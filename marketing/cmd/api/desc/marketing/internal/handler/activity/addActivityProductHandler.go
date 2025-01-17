package activity

import (
	"net/http"

	"QMall/marketing/cmd/api/desc/marketing/internal/logic/activity"
	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加活动商品信息
func AddActivityProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddActivityProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := activity.NewAddActivityProductLogic(r.Context(), svcCtx)
		resp, err := l.AddActivityProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
