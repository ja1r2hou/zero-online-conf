package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero-online-conf/api/internal/logic"
	"zero-online-conf/api/internal/svc"
)

// GenerateRSAKeyHandler 获取新的公钥和私钥
func GenerateRSAKeyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGenerateRSAKeyLogic(r.Context(), svcCtx)
		resp, err := l.GenerateRSAKey()

		//因为返回值会打印日志 所以旧暂时用go-zero自带的httpx
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
