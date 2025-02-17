package svc

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"zero-online-conf/common/util"
	"zero-online-conf/rpc/internal/model"
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
	if req.UserName == "lookPwd" {
		u.Logger.Error("UserSvc:UserLogin:用户不存在")
		return "", errors.New("用户不存在！")
	}
	pwd := UserMap[req.UserName]
	if pwd == "" {
		u.Logger.Error("UserSvc:UserLogin:nil:用户不存在")
		return "", errors.New("用户不存在！")
	}

	verify := util.Md5Verify(pwd, UserSalt, req.Password)
	if !verify {
		u.Logger.Error("UserSvc:UserLogin:用户名或密码错误！")
		return "", errors.New("登录错误！")
	}
	//生成token 2个小时缓存时间
	tokenExpireTime := time.Now().Add(2 * time.Hour)
	userToken := model.UserToken{
		UserName:            req.UserName,
		Ip:                  req.Ip,
		UserTokenExpireTime: tokenExpireTime,
	}
	fmt.Println(userToken)
	return "", nil
}

func (u *UserSvc) VerifyUserToken() {

}
