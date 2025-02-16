package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/collection"
	"time"
	"zero-online-conf/common/util"
	"zero-online-conf/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Cache  *collection.Cache
}

var UserMap map[string]string

const (
	UserSalt = "mm9icik5kk7dkd0odos134558dfhnsdqqd" //用户登录salt 最好修改一下
	UserPwd  = "admin123"                           //登录密码  需要修改
	LookPwd  = "admin"                              //查看配置内容的时候需要用到
)

func init() {

	UserMap = make(map[string]string, 0)
	UserMap["admin"] = util.Md5encoded(UserPwd, UserSalt) //默认用户和密码
}

func NewServiceContext(c config.Config) *ServiceContext {
	cache, err := collection.NewCache(24*time.Hour, collection.WithLimit(1000))
	if err != nil {
		panic("系统加载错误:缓存初始化失败！")
	}
	//检测是否修改初始密码
	if UserPwd == "admin" || UserPwd == "admin123" {
		fmt.Println("系统加载错误:需要修改初始密码！")
		panic("系统加载错误:需要修改初始密码！")
	}
	//检查初始查看密码是否有修改
	if LookPwd == "admin" || LookPwd == "admin123" {
		fmt.Println("系统加载错误:需要修改初始查看密码！")
		panic("系统加载错误:需要修改初始查看密码！")
	}

	return &ServiceContext{
		Config: c,
		Cache:  cache,
	}
}
