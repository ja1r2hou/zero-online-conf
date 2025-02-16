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

// UserSalt 用户登录salt 最好修改一下
const UserSalt = "mm9icik5kk7dkd0odos134558dfhnsdqqd"

func init() {

	UserMap = make(map[string]string, 0)
	UserMap["admin"] = util.Md5encoded("admin", UserSalt) //默认用户和密码
}

func NewServiceContext(c config.Config) *ServiceContext {
	cache, err := collection.NewCache(24*time.Hour, collection.WithLimit(1000))
	if err != nil {
		panic("系统加载错误:缓存初始化失败！")
	}
	if UserMap["admin"] == "f244c91210dd383c94ed7abccb08a9f4" || UserMap["admin"] == "ea26857feaf048a7606b22f7cdc57625" {
		fmt.Println("系统加载错误:需要修改初始密码！")
		panic("系统加载错误:需要修改初始密码！")
	}

	return &ServiceContext{
		Config: c,
		Cache:  cache,
	}
}
