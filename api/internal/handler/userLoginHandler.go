package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero-online-conf/api/internal/logic"
	"zero-online-conf/api/internal/svc"
	"zero-online-conf/api/internal/types"
	"zero-online-conf/common/respx"
	"zero-online-conf/common/util"
)

// UserLoginHandler 用户登录
func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			respx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)

		l.Logger.Info("userLogin:req:", req)

		ip := util.ClientIP(r)
		l.Logger.Info("userLogin:ip:", ip)

		resp, err := l.UserLogin(&req)
		if err != nil {
			respx.ErrorCtx(r.Context(), w, err)
		} else {
			respx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
