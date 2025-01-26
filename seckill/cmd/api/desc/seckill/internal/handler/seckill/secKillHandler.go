package seckill

import (
	"QMall/common"
	"net/http"

	"QMall/seckill/cmd/api/desc/seckill/internal/logic/seckill"
	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加秒杀订单V2
func SecKillHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SecKillReq
		if err := httpx.Parse(r, &req); err != nil {
			common.RespOk(r.Context(), w, "", "请求失败！", common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := seckill.NewSecKillLogic(r.Context(), svcCtx)
		resp, err := l.SecKill(&req)
		if err != nil {
			common.RespOk(r.Context(), w, "", err.Error(), common.TOKEN_FAIL)
			//httpx.ErrorCtx(r.Context(), w, err)
		} else {
			common.RespOk(r.Context(), w, resp, "获取成功！", common.SUCCESS)
			//httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
