package logic

import (
	"context"

	"github.com/kiritokun07/go-zero-study/service/shorturl/rpc/transform/internal/svc"
	"github.com/kiritokun07/go-zero-study/service/shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenReq) (*transform.ShortenResp, error) {
	// todo: add your logic here and delete this line

	return &transform.ShortenResp{}, nil
}
