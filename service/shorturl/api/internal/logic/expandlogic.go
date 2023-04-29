package logic

import (
	"context"

	"github.com/kiritokun07/go-zero-study/service/shorturl/api/internal/svc"
	"github.com/kiritokun07/go-zero-study/service/shorturl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (resp *types.ExpandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
