package seckill

import (
	"net/http"

	"QMall/seckill/cmd/api/desc/seckill/internal/logic/seckill"
	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加秒杀限额
func IncreaseSecKillQuotaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IncreaseSecKillQuotaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := seckill.NewIncreaseSecKillQuotaLogic(r.Context(), svcCtx)
		resp, err := l.IncreaseSecKillQuota(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
