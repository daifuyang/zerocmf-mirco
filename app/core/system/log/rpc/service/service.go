// Code generated by goctl. DO NOT EDIT.
// Source: log.proto

package service

import (
	"context"

	"app/core/log/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SystemLogReq  = pb.SystemLogReq
	SystemLogResp = pb.SystemLogResp

	Service interface {
		// 系统日志记录
		LogSystem(ctx context.Context, in *SystemLogReq, opts ...grpc.CallOption) (*SystemLogResp, error)
	}

	defaultService struct {
		cli zrpc.Client
	}
)

func NewService(cli zrpc.Client) Service {
	return &defaultService{
		cli: cli,
	}
}

// 系统日志记录
func (m *defaultService) LogSystem(ctx context.Context, in *SystemLogReq, opts ...grpc.CallOption) (*SystemLogResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.LogSystem(ctx, in, opts...)
}
