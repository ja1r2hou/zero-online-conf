package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	LogConf logx.LogConf
	MySql   struct {
		DataSource  string
		MaxCon      int
		MaxIdle     int
		MaxLifeTime int64
	}
	AuthToken struct {
		AccessSecret string
		AccessExpire int64
	}
}
