package svc

import (
	"github.com/kiritokun07/go-zero-study/service/shorturl/api/internal/config"
	"github.com/kiritokun07/go-zero-study/service/shorturl/rpc/transform/transformer"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer //rpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.TransformerRpcConf)),
	}
}
