package logic

import (
	"context"

	"zero-online-conf/rpc/internal/svc"
	"zero-online-conf/rpc/onlineConf"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userSvc *svc.UserSvc
}

func NewUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAuthLogic {
	return &UserAuthLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userSvc: svc.NewUserSvc(ctx, svcCtx),
	}
}

// UserAuth 验证用户token
func (l *UserAuthLogic) UserAuth(in *onlineConf.UserAuthReq) (*onlineConf.UserAuthResp, error) {

	return &onlineConf.UserAuthResp{}, nil
}
