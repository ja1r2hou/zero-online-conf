package logic

import (
	"context"
	"time"
	"zero-online-conf/rpc/onlineconfrpc"

	"zero-online-conf/api/internal/svc"
	"zero-online-conf/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (*types.Response, error) {

	resp, err := l.svcCtx.OnlineConf.Ping(l.ctx, &onlineconfrpc.Request{Ping: time.Now().String()})

	response := &types.Response{resp.Pong}

	return response, err
}
