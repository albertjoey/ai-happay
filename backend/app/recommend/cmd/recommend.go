package main

import (
	"flag"
	"fmt"

	"happy/app/recommend/internal/config"
	"happy/app/recommend/internal/handler"
	"happy/app/recommend/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/recommend/etc/recommend.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	defer func() {
		// 停止定时任务
		if ctx.Cron != nil {
			ctx.Cron.Stop()
		}
	}()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting recommend service at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
