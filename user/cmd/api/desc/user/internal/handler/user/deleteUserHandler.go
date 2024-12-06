package user

import (
	"QMall/user/cmd/api/desc/user/internal/logic/user"
	"QMall/user/cmd/api/desc/user/internal/svc"
	"QMall/user/cmd/api/desc/user/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 根据id删除用户
func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDeleteUserLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
