package svc

import (
	"github.com/zeromicro/go-zero/core/collection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"zero-online-conf/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Cache  *collection.Cache
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	cache, err := collection.NewCache(24*time.Hour, collection.WithLimit(1000))
	if err != nil {
		panic("系统加载错误:缓存初始化失败！")
	}

	db, err := gorm.Open(mysql.Open(c.MySql.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(c.MySql.MaxIdle)
	sqlDb.SetMaxOpenConns(c.MySql.MaxCon)
	sqlDb.SetConnMaxLifetime(time.Duration(c.MySql.MaxLifeTime))

	return &ServiceContext{
		Config: c,
		Cache:  cache,
		DB:     db,
	}
}
