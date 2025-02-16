package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	LogConf   logx.LogConf
	AuthToken struct {
		AccessSecret string
		AccessExpire int64
	}
}
