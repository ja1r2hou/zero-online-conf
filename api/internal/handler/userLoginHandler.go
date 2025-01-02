package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
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
		ip := util.ClientIP(r)
		logx.Infof("userLogin:ip: %v", ip)

		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			respx.Error(w, err, r.Context())
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)

		resp, err := l.UserLogin(&req)
		if err != nil {
			respx.Error(w, err, r.Context())
		} else {
			respx.OkJson(w, resp, r.Context())
		}
	}
}
