package logic

import (
	"context"

	"github.com/kiritokun07/go-zero-study/service/shorturl/model"
	"github.com/kiritokun07/go-zero-study/service/shorturl/rpc/transform/internal/svc"
	"github.com/kiritokun07/go-zero-study/service/shorturl/rpc/transform/transform"
	"github.com/kiritokun07/go-zero-study/shared/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandReq) (*transform.ExpandResp, error) {
	shorturl, err := l.svcCtx.ShorturlModel.FindOne(l.ctx, in.Shorten)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.ShorturlNotExist
		}
		return nil, err
	}
	return &transform.ExpandResp{
		Url: shorturl.Url,
	}, nil
}
