package logic

import (
	"context"
	"time"
	"zero-online-conf/rpc/internal/svc"
	"zero-online-conf/rpc/onlineConf"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *onlineConf.Request) (*onlineConf.Response, error) {
	v, exist := l.svcCtx.Cache.Get("rpc")
	if !exist {
		nowTime := time.Now().String()
		l.svcCtx.Cache.SetWithExpire("rpc", nowTime, 10*time.Second)
		return &onlineConf.Response{Pong: nowTime}, nil
	}
	value, _ := v.(string)

	return &onlineConf.Response{Pong: value}, nil
}
