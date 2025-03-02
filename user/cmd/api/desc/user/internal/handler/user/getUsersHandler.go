package user

import (
	"QMall/user/cmd/api/desc/user/internal/logic/user"
	"QMall/user/cmd/api/desc/user/internal/svc"
	"QMall/user/cmd/api/desc/user/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

// GetUsersHandler 获取用户列表
func GetUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserListReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			//httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUsersLogic(r.Context(), svcCtx)
		resp, err := l.GetUsers(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
