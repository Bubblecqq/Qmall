package activity

import (
	"QMall/common"
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
			//httpx.ErrorCtx(r.Context(), w, err)
			common.RespOk(r.Context(), w, "", "请求失败！", common.TOKEN_FAIL)
			return
		}

		l := activity.NewAddActivityLogic(r.Context(), svcCtx)
		resp, err := l.AddActivity(&req)
		if err != nil {
			common.RespOk(r.Context(), w, "", err.Error(), common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			common.RespOk(r.Context(), w, resp, "获取成功！", common.SUCCESS)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
