// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package service

import (
	"context"

	"app/core/user/admin/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AdminShowReq = pb.AdminShowReq
	LoginData    = pb.LoginData
	LoginResp    = pb.LoginResp

	Service interface {
		// 获取单个管理员
		AdminShow(ctx context.Context, in *AdminShowReq, opts ...grpc.CallOption) (*LoginResp, error)
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

// 获取单个管理员
func (m *defaultService) AdminShow(ctx context.Context, in *AdminShowReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.AdminShow(ctx, in, opts...)
}
