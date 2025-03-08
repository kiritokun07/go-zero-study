package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/kiritokun07/go-zero-study/service/xxljob/api/internal/config"
	"github.com/kiritokun07/go-zero-study/service/xxljob/api/internal/handler"
	"github.com/kiritokun07/go-zero-study/service/xxljob/api/internal/svc"
	"github.com/kiritokun07/go-zero-study/shared/xxljob"
	"github.com/xxl-job/xxl-job-executor-go/example/task"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/xxljob-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	xxljob.InitXxlJob(context.TODO(), ctx.Config.XxlJob,
		xxljob.WithTask("task.test", task.Test),
		xxljob.WithTask("task.test2", task.Test2),
	)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
