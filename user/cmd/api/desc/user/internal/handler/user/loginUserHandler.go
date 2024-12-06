package user

import (
	"goZeroDemo4/common"
	"goZeroDemo4/user/cmd/api/desc/user/internal/logic/user"
	"goZeroDemo4/user/cmd/api/desc/user/internal/svc"
	"goZeroDemo4/user/cmd/api/desc/user/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// LoginUserHandler 单点用户登录
func LoginUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginUserReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			common.RespOk(r.Context(), w, "", "登录失败！", common.TOKEN_FAIL)
			return
		}

		l := user.NewLoginUserLogic(r.Context(), svcCtx)
		resp, err := l.LoginUser(&req)
		if err != nil {
			common.RespOk(r.Context(), w, resp, "登录失败！", common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			common.RespOk(r.Context(), w, resp, "登录成功！", common.SUCCESS)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
