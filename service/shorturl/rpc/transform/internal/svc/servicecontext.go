package svc

import "github.com/kiritokun07/go-zero-study/service/shorturl/rpc/transform/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
