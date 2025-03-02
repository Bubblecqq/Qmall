// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package user

import (
	"context"

	"QMall/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateUserReq   = pb.CreateUserReq
	CreateUserResp  = pb.CreateUserResp
	DeleteUserReq   = pb.DeleteUserReq
	DeleteUserResp  = pb.DeleteUserResp
	GetUserByIdReq  = pb.GetUserByIdReq
	GetUserByIdResp = pb.GetUserByIdResp
	LoginUserReq    = pb.LoginUserReq
	LoginUserResp   = pb.LoginUserResp
	TokenReq        = pb.TokenReq
	TokenResp       = pb.TokenResp
	User            = pb.User

	UserZrpcClient interface {
		CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error)
		DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error)
		GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
		LoginUser(ctx context.Context, in *LoginUserReq, opts ...grpc.CallOption) (*LoginUserResp, error)
		GetUserToken(ctx context.Context, in *TokenReq, opts ...grpc.CallOption) (*TokenResp, error)
	}

	defaultUserZrpcClient struct {
		cli zrpc.Client
	}
)

func NewUserZrpcClient(cli zrpc.Client) UserZrpcClient {
	return &defaultUserZrpcClient{
		cli: cli,
	}
}

func (m *defaultUserZrpcClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultUserZrpcClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}

func (m *defaultUserZrpcClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultUserZrpcClient) LoginUser(ctx context.Context, in *LoginUserReq, opts ...grpc.CallOption) (*LoginUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.LoginUser(ctx, in, opts...)
}

func (m *defaultUserZrpcClient) GetUserToken(ctx context.Context, in *TokenReq, opts ...grpc.CallOption) (*TokenResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUserToken(ctx, in, opts...)
}
