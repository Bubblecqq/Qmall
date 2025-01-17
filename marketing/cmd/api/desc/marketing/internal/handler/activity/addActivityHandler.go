package activity

import (
	"net/http"

	"QMall/marketing/cmd/api/desc/marketing/internal/logic/activity"
	"QMall/marketing/cmd/api/desc/marketing/internal/svc"
	"QMall/marketing/cmd/api/desc/marketing/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加活动表
func AddActivityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddActivityReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := activity.NewAddActivityLogic(r.Context(), svcCtx)
		resp, err := l.AddActivity(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
