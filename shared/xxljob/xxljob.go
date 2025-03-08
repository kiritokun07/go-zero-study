package xxljob

import (
	"context"
	"log"

	"github.com/xxl-job/xxl-job-executor-go"
)

type XxlOption func(xxl.Executor)

// 添加配置结构体
type Config struct {
	ServerAddr   string
	AccessToken  string
	ExecutorIp   string
	ExecutorPort string
	RegistryKey  string
}

func InitXxlJob(ctx context.Context, cfg Config, opts ...XxlOption) {
	exec := xxl.NewExecutor(
		xxl.ServerAddr(cfg.ServerAddr),
		xxl.AccessToken(cfg.AccessToken),
		xxl.ExecutorIp(cfg.ExecutorIp),
		xxl.ExecutorPort(cfg.ExecutorPort),
		xxl.RegistryKey(cfg.RegistryKey),
		xxl.SetLogger(&logger{}),
	)
	exec.Init()
	//设置日志查看handler
	exec.LogHandler(func(req *xxl.LogReq) *xxl.LogRes {
		return &xxl.LogRes{Code: 200, Msg: "", Content: xxl.LogResContent{
			FromLineNum: req.FromLineNum,
			ToLineNum:   2,
			LogContent:  "这个是自定义日志handler",
			IsEnd:       true,
		}}
	})
	//注册任务handler
	for _, opt := range opts {
		opt(exec)
	}
	log.Fatal(exec.Run())
}

func WithTask(pattern string, task xxl.TaskFunc) XxlOption {
	return func(exec xxl.Executor) {
		exec.RegTask(pattern, task)
	}
}
