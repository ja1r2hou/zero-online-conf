package svc

import (
	"github.com/zeromicro/go-zero/core/collection"
	"time"
	"zero-online-conf/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Cache  *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	cache, err := collection.NewCache(24*time.Hour, collection.WithLimit(1000))
	if err != nil {
		panic("系统加载错误:缓存初始化失败！")
	}
	return &ServiceContext{
		Config: c,
		Cache:  cache,
	}
}
