package svc

import (
	"github.com/kiritokun07/go-zero-study/service/shorturl/model"
	"github.com/kiritokun07/go-zero-study/service/shorturl/rpc/transform/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ShorturlModel model.ShorturlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	testUrl := sqlx.NewMysql(c.Mysql.Test)
	return &ServiceContext{
		Config:        c,
		ShorturlModel: model.NewShorturlModel(testUrl, c.CacheConf),
	}
}
