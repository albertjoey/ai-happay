package main

import (
	"flag"
	"fmt"

	"happy/app/content/internal/config"
	"happy/app/content/internal/handler"
	"happy/app/content/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/content/etc/content.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting content service at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
