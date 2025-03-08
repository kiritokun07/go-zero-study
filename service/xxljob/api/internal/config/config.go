package config

import (
	"github.com/kiritokun07/go-zero-study/shared/xxljob"
	"github.com/zeromicro/go-zero/rest"
)

type (
	Config struct {
		rest.RestConf
		XxlJob xxljob.Config
	}
)
