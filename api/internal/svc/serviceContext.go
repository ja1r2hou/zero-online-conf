package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-online-conf/api/internal/config"
	"zero-online-conf/rpc/onlineconfrpc"
)

type ServiceContext struct {
	Config     config.Config
	Onlineconf onlineconfrpc.OnlineConfRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Onlineconf: onlineconfrpc.NewOnlineConfRpc(zrpc.MustNewClient(c.OnlineConf)),
	}
}
