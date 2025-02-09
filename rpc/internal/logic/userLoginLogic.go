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
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserLogin 登录
func (l *UserLoginLogic) UserLogin(in *onlineConf.UserLoginReq) (*onlineConf.UserLoginResp, error) {
	// todo: add your logic here and delete this line

	return &onlineConf.UserLoginResp{}, nil
}
