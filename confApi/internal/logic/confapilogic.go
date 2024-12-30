package logic

import (
	"context"

	"zero-online-conf/confApi/internal/svc"
	"zero-online-conf/confApi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfApiLogic {
	return &ConfApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfApiLogic) ConfApi(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
