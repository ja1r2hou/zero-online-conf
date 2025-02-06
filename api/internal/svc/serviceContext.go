package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
	"zero-online-conf/api/internal/config"
	"zero-online-conf/rpc/onlineconfrpc"
)

type ServiceContext struct {
	Config     config.Config
	OnlineConf onlineconfrpc.OnlineConfRpc
	//Auth       rest.Middleware
}

func (m *ServiceContext) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OnlineConf: onlineconfrpc.NewOnlineConfRpc(zrpc.MustNewClient(c.OnlineConf)),
	}
}
