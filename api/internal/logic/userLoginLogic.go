package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"zero-online-conf/api/internal/svc"
	"zero-online-conf/api/internal/types"
	"zero-online-conf/rpc/onlineconfrpc"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserLoginLogic 用户登录
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq, ip string) (*types.LoginResp, error) {
	loginReq := &onlineconfrpc.UserLoginReq{
		UserName: req.UserName,
		Password: req.Password,
		Ip:       ip,
	}
	loginResp, err := l.svcCtx.OnlineConf.UserLogin(l.ctx, loginReq)
	if err != nil {
		return nil, err
	}
	resp := &types.LoginResp{
		UserName: loginResp.UserName,
		//UserId:   loginResp,
		Token: loginResp.Token,
	}

	return resp, nil
}
