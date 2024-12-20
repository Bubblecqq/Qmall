package user

import (
	"net/http"

	"QMall/user/cmd/api/desc/user/internal/logic/user"
	"QMall/user/cmd/api/desc/user/internal/svc"
	"QMall/user/cmd/api/desc/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 分布式Token提取
func GetUserTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserTokenLogic(r.Context(), svcCtx)
		resp, err := l.GetUserToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
