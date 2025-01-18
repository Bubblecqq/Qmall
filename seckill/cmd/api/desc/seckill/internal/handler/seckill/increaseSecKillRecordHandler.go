package seckill

import (
	"net/http"

	"QMall/seckill/cmd/api/desc/seckill/internal/logic/seckill"
	"QMall/seckill/cmd/api/desc/seckill/internal/svc"
	"QMall/seckill/cmd/api/desc/seckill/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加秒杀记录
func IncreaseSecKillRecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IncreaseSecKillRecordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := seckill.NewIncreaseSecKillRecordLogic(r.Context(), svcCtx)
		resp, err := l.IncreaseSecKillRecord(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
