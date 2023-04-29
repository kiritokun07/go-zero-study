package logic

import (
	"context"

	"github.com/kiritokun07/go-zero-study/service/shorturl/api/internal/svc"
	"github.com/kiritokun07/go-zero-study/service/shorturl/api/internal/types"
	"github.com/kiritokun07/go-zero-study/service/shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	rpcResp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transform.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return &types.ShortenResp{}, err
	}
	return &types.ShortenResp{
		Shorten: rpcResp.Shorten,
	}, nil
}
