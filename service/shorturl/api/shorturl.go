package main

import (
	"flag"
	"fmt"

	"github.com/kiritokun07/go-zero-study/service/shorturl/api/internal/config"
	"github.com/kiritokun07/go-zero-study/service/shorturl/api/internal/handler"
	"github.com/kiritokun07/go-zero-study/service/shorturl/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shorturl-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	logx.DisableStat()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
