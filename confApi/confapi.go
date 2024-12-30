package main

import (
	"flag"
	"fmt"

	"zero-online-conf/confApi/internal/config"
	"zero-online-conf/confApi/internal/handler"
	"zero-online-conf/confApi/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// go run xxx.go  -f /你的配置文件路径
// win confapi.exe  -f /你的配置文件路径

var configFile = flag.String("f", "etc/confapi-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
