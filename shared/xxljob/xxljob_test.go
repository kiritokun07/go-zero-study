package xxljob

import (
	"context"
	"testing"

	"github.com/xxl-job/xxl-job-executor-go/example/task"
	"github.com/zeromicro/go-zero/core/threading"
)

func TestInitXxlJob(t *testing.T) {
	cfg := Config{
		ServerAddr:   "http://127.0.0.1:8080/xxl-job-admin",
		AccessToken:  "default_token",
		ExecutorIp:   "127.0.0.1",
		ExecutorPort: "9998",
		RegistryKey:  "calculate-go-service",
	}

	InitXxlJob(context.TODO(), cfg,
		WithTask("task.test", task.Test),
		WithTask("task.test2", task.Test2),
		WithTask("task.panic", task.Panic),
	)
}

func TestInitXxlJob2(t *testing.T) {
	cfg := Config{
		ServerAddr:   "http://127.0.0.1:8080/xxl-job-admin",
		AccessToken:  "default_token",
		ExecutorIp:   "127.0.0.1",
		ExecutorPort: "9998",
		RegistryKey:  "calculate-go-service",
	}

	threading.GoSafe(func() {
		InitXxlJob(context.Background(), cfg)
	})
}
