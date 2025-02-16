package svc

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-online-conf/rpc/onlineConf"
)

type UserSvc struct {
	logx.Logger
	ctx    context.Context
	svcCtx *ServiceContext
}

func NewUserSvc(ctx context.Context, svcCtx *ServiceContext) *UserSvc {
	return &UserSvc{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (u *UserSvc) UserLogin(req *onlineConf.UserLoginReq) (string, error) {
	//pwd := UserMap[req.UserName]

	return "", nil
}

func (u *UserSvc) VerifyUserToken() {

}
