package logic

import (
	"context"
	"github.com/kiritokun07/go-zero-study/service/xxljob/api/internal/svc"
	"github.com/kiritokun07/go-zero-study/service/xxljob/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type XxljobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewXxljobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *XxljobLogic {
	return &XxljobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *XxljobLogic) Xxljob(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
