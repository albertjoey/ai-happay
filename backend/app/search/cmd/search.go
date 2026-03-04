package main

import (
	"flag"
	"fmt"

	"happy/app/search/internal/config"
	"happy/app/search/internal/handler"
	"happy/app/search/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/search/etc/search.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting search server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
