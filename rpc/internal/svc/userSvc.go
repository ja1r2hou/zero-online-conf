package svc

import (
	"context"
	"encoding/json"
	"errors"
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

// TokenPrivatKey  用来处理 用户登录的token的解密
var TokenPrivatKey = ""

// TokenPublicKey 用来处理 用户登录的token的加密
var TokenPublicKey = ""

func init() {
	TokenPrivatKey, TokenPublicKey, _ = util.GenerateRSAKey(2048)
}

func NewUserSvc(ctx context.Context, svcCtx *ServiceContext) *UserSvc {
	return &UserSvc{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserLogin 用户登录Svc
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

	verify := util.Md5Verify(req.Password, UserSalt, pwd)
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
	tokenJson, _ := json.Marshal(userToken)
	encrypt, err := util.RsaEncrypt(string(tokenJson), TokenPublicKey)
	if err != nil {
		return "", err
	}

	return encrypt, nil
}

// VerifyUserToken 验证用户token
func (u *UserSvc) VerifyUserToken(req *onlineConf.UserAuthReq) (bool, error) {
	//解密token
	decrypt, err := util.RsaDecrypt(req.Token, TokenPrivatKey)
	if err != nil {
		u.Logger.Error("userSvc:VerifyUserToken:RsaDecrypt:解密失败！")
		return false, err
	}
	userToken := model.UserToken{}
	err = json.Unmarshal([]byte(decrypt), &userToken)
	if err != nil {
		u.Logger.Error("userSvc:VerifyUserToken:RsaDecrypt:解析json失败！")
		return false, err
	}
	// 判断当前用户 是否和密文一致 如果有一项与密文不一致 就判断是非正常用户
	if userToken.UserName != req.UserName || userToken.Ip != req.Ip || time.Now().Unix() > userToken.UserTokenExpireTime.Unix() {
		u.Logger.Error("userSvc:VerifyUserToken:RsaDecrypt:登录的用户与密文不符合或是过期时间已经大于现在！")
		return false, err
	}
	return true, nil
}
