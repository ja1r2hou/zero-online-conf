package handler

import (
	"net/http"
	"zero-online-conf/api/internal/logic"
	"zero-online-conf/api/internal/svc"
	"zero-online-conf/common/respx"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()
		if err != nil {
			respx.ErrorCtx(r.Context(), w, err)
		} else {
			respx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
