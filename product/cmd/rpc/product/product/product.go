// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: product.proto

package product

import (
	"context"

	"QMall/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateProductReq      = pb.CreateProductReq
	CreateProductResp     = pb.CreateProductResp
	CreateProductSkuReq   = pb.CreateProductSkuReq
	CreateProductSkuResp  = pb.CreateProductSkuResp
	DeleteProductReq      = pb.DeleteProductReq
	DeleteProductResp     = pb.DeleteProductResp
	DeleteProductSkuReq   = pb.DeleteProductSkuReq
	DeleteProductSkuResp  = pb.DeleteProductSkuResp
	GetProductByIdReq     = pb.GetProductByIdReq
	GetProductByIdResp    = pb.GetProductByIdResp
	GetProductListReq     = pb.GetProductListReq
	GetProductListResp    = pb.GetProductListResp
	GetProductSkuByIdReq  = pb.GetProductSkuByIdReq
	GetProductSkuByIdResp = pb.GetProductSkuByIdResp
	GetProductSkuListReq  = pb.GetProductSkuListReq
	GetProductSkuListResp = pb.GetProductSkuListResp
	Product               = pb.Product
	ProductSku            = pb.ProductSku

	ProductZrpcClient interface {
		CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error)
		GetProductList(ctx context.Context, in *GetProductListReq, opts ...grpc.CallOption) (*GetProductListResp, error)
		DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error)
		GetProduct(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdResp, error)
		CreateProductSku(ctx context.Context, in *CreateProductSkuReq, opts ...grpc.CallOption) (*CreateProductSkuResp, error)
		GetProductListSku(ctx context.Context, in *GetProductSkuListReq, opts ...grpc.CallOption) (*GetProductSkuListResp, error)
		DeleteProductSku(ctx context.Context, in *DeleteProductSkuReq, opts ...grpc.CallOption) (*DeleteProductSkuResp, error)
		GetProductSku(ctx context.Context, in *GetProductSkuByIdReq, opts ...grpc.CallOption) (*GetProductSkuByIdResp, error)
	}

	defaultProductZrpcClient struct {
		cli zrpc.Client
	}
)

func NewProductZrpcClient(cli zrpc.Client) ProductZrpcClient {
	return &defaultProductZrpcClient{
		cli: cli,
	}
}

func (m *defaultProductZrpcClient) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CreateProduct(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) GetProductList(ctx context.Context, in *GetProductListReq, opts ...grpc.CallOption) (*GetProductListResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.GetProductList(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.DeleteProduct(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) GetProduct(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.GetProduct(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) CreateProductSku(ctx context.Context, in *CreateProductSkuReq, opts ...grpc.CallOption) (*CreateProductSkuResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CreateProductSku(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) GetProductListSku(ctx context.Context, in *GetProductSkuListReq, opts ...grpc.CallOption) (*GetProductSkuListResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.GetProductListSku(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) DeleteProductSku(ctx context.Context, in *DeleteProductSkuReq, opts ...grpc.CallOption) (*DeleteProductSkuResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.DeleteProductSku(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) GetProductSku(ctx context.Context, in *GetProductSkuByIdReq, opts ...grpc.CallOption) (*GetProductSkuByIdResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.GetProductSku(ctx, in, opts...)
}
