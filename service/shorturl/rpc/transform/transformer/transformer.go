// Code generated by goctl. DO NOT EDIT.
// Source: transform.proto

package transformer

import (
	"context"

	"github.com/kiritokun07/go-zero-study/service/shorturl/rpc/transform/transform"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ExpandReq   = transform.ExpandReq
	ExpandResp  = transform.ExpandResp
	ShortenReq  = transform.ShortenReq
	ShortenResp = transform.ShortenResp

	Transformer interface {
		// 通过原链设置短链
		Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error)
		// 通过短链获取原链
		Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error)
	}

	defaultTransformer struct {
		cli zrpc.Client
	}
)

func NewTransformer(cli zrpc.Client) Transformer {
	return &defaultTransformer{
		cli: cli,
	}
}

// 通过原链设置短链
func (m *defaultTransformer) Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error) {
	client := transform.NewTransformerClient(m.cli.Conn())
	return client.Shorten(ctx, in, opts...)
}

// 通过短链获取原链
func (m *defaultTransformer) Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error) {
	client := transform.NewTransformerClient(m.cli.Conn())
	return client.Expand(ctx, in, opts...)
}
