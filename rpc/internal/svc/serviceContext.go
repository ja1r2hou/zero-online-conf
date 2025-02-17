package svc

import (
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
)

func init() {
	UserMap = make(map[string]string, 0)
	UserMap["admin"] = util.Md5encoded("admin", UserSalt) //默认用户和密码 可手动新增用户 ，一定要修改后自行删除 util.Md5encoded("admin", UserSalt)  直接设置密文

	UserMap["lookPwd"] = util.Md5encoded("admin", UserSalt) //手动获取后密文后写死设置进去
}

func NewServiceContext(c config.Config) *ServiceContext {
	cache, err := collection.NewCache(24*time.Hour, collection.WithLimit(1000))
	if err != nil {
		panic("系统加载错误:缓存初始化失败！")
	}

	if true { //修改密码后 直接删掉即可
		panic("系统加载错误:需要修改初始查看密码/用户密码修改！")
	}

	return &ServiceContext{
		Config: c,
		Cache:  cache,
	}
}
