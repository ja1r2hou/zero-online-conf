package handler

import (
	"net/http"
	"zero-online-conf/common/respx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-online-conf/api/internal/logic"
	"zero-online-conf/api/internal/svc"
	"zero-online-conf/api/internal/types"
)

// GetServiceListHandler 获取服务注册列表
func GetServiceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetServiceListReq
		if err := httpx.Parse(r, &req); err != nil {
			respx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetServiceListLogic(r.Context(), svcCtx)
		resp, err := l.GetServiceList(&req)
		if err != nil {
			respx.ErrorCtx(r.Context(), w, err)
		} else {
			respx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
