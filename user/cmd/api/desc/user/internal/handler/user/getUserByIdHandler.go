package user

import (
	"fmt"
	"goZeroDemo4/common"
	"goZeroDemo4/user/cmd/api/desc/user/internal/logic/user"
	"goZeroDemo4/user/cmd/api/desc/user/internal/svc"
	"goZeroDemo4/user/cmd/api/desc/user/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 根据id获取用户
func GetUserByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserByIdReq
		if err := httpx.Parse(r, &req); err != nil {

			//xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			//httpx.ErrorCtx(r.Context(), w, err)
			common.RespOk(r.Context(), w, "", "失败！", common.TOKEN_FAIL)
			return
		}

		l := user.NewGetUserByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUserById(&req)
		fmt.Println("resp>", resp)
		if err != nil {
			common.RespOk(r.Context(), w, resp, "失败！", common.TOKEN_FAIL)

			//xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			common.RespOk(r.Context(), w, resp, "获取成功！", common.SUCCESS)
			//xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
