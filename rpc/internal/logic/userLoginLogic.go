package logic

import (
	"context"
	"zero-online-conf/rpc/internal/svc"
	"zero-online-conf/rpc/onlineConf"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userSvc *svc.UserSvc
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userSvc: svc.NewUserSvc(ctx, svcCtx),
	}
}

// UserLogin 登录
func (l *UserLoginLogic) UserLogin(in *onlineConf.UserLoginReq) (*onlineConf.UserLoginResp, error) {
	token, err := l.userSvc.UserLogin(in)
	if err != nil {
		return nil, err
	}
	resp := &onlineConf.UserLoginResp{
		Token:    token,
		UserName: in.UserName,
	}
	return resp, nil
}
