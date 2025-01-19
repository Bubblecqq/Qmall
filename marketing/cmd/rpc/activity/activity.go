// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: marketing.proto

package activity

import (
	"context"

	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Activity                   = pb.Activity
	ActivityProduct            = pb.ActivityProduct
	ActivityProductSku         = pb.ActivityProductSku
	ActivityTime               = pb.ActivityTime
	AddActivityProductReq      = pb.AddActivityProductReq
	AddActivityProductResp     = pb.AddActivityProductResp
	AddActivityProductSkuReq   = pb.AddActivityProductSkuReq
	AddActivityProductSkuResp  = pb.AddActivityProductSkuResp
	AddActivityReq             = pb.AddActivityReq
	AddActivityResp            = pb.AddActivityResp
	AddActivityTimeReq         = pb.AddActivityTimeReq
	AddActivityTimeResp        = pb.AddActivityTimeResp
	GetActivityProductByIdReq  = pb.GetActivityProductByIdReq
	GetActivityProductByIdResp = pb.GetActivityProductByIdResp

	ActivityZrpcClient interface {
		AddActivity(ctx context.Context, in *AddActivityReq, opts ...grpc.CallOption) (*AddActivityResp, error)
		AddActivityTime(ctx context.Context, in *AddActivityTimeReq, opts ...grpc.CallOption) (*AddActivityTimeResp, error)
		AddActivityProduct(ctx context.Context, in *AddActivityProductReq, opts ...grpc.CallOption) (*AddActivityProductResp, error)
		AddActivityProductSku(ctx context.Context, in *AddActivityProductSkuReq, opts ...grpc.CallOption) (*AddActivityProductSkuResp, error)
		// 查询接口
		GetActivityProductById(ctx context.Context, in *GetActivityProductByIdReq, opts ...grpc.CallOption) (*GetActivityProductByIdResp, error)
	}

	defaultActivityZrpcClient struct {
		cli zrpc.Client
	}
)

func NewActivityZrpcClient(cli zrpc.Client) ActivityZrpcClient {
	return &defaultActivityZrpcClient{
		cli: cli,
	}
}

func (m *defaultActivityZrpcClient) AddActivity(ctx context.Context, in *AddActivityReq, opts ...grpc.CallOption) (*AddActivityResp, error) {
	client := pb.NewActivityClient(m.cli.Conn())
	return client.AddActivity(ctx, in, opts...)
}

func (m *defaultActivityZrpcClient) AddActivityTime(ctx context.Context, in *AddActivityTimeReq, opts ...grpc.CallOption) (*AddActivityTimeResp, error) {
	client := pb.NewActivityClient(m.cli.Conn())
	return client.AddActivityTime(ctx, in, opts...)
}

func (m *defaultActivityZrpcClient) AddActivityProduct(ctx context.Context, in *AddActivityProductReq, opts ...grpc.CallOption) (*AddActivityProductResp, error) {
	client := pb.NewActivityClient(m.cli.Conn())
	return client.AddActivityProduct(ctx, in, opts...)
}

func (m *defaultActivityZrpcClient) AddActivityProductSku(ctx context.Context, in *AddActivityProductSkuReq, opts ...grpc.CallOption) (*AddActivityProductSkuResp, error) {
	client := pb.NewActivityClient(m.cli.Conn())
	return client.AddActivityProductSku(ctx, in, opts...)
}

// 查询接口
func (m *defaultActivityZrpcClient) GetActivityProductById(ctx context.Context, in *GetActivityProductByIdReq, opts ...grpc.CallOption) (*GetActivityProductByIdResp, error) {
	client := pb.NewActivityClient(m.cli.Conn())
	return client.GetActivityProductById(ctx, in, opts...)
}
