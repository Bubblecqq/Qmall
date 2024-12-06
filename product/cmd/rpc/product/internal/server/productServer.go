// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: product.proto

package server

import (
	"context"

	"QMall/product/cmd/rpc/product/internal/logic"
	"QMall/product/cmd/rpc/product/internal/svc"
	"QMall/product/cmd/rpc/product/pb"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) CreateProduct(ctx context.Context, in *pb.CreateProductReq) (*pb.CreateProductResp, error) {
	l := logic.NewCreateProductLogic(ctx, s.svcCtx)
	return l.CreateProduct(in)
}

func (s *ProductServer) GetProductList(ctx context.Context, in *pb.GetProductListReq) (*pb.GetProductListResp, error) {
	l := logic.NewGetProductListLogic(ctx, s.svcCtx)
	return l.GetProductList(in)
}

func (s *ProductServer) DeleteProduct(ctx context.Context, in *pb.DeleteProductReq) (*pb.DeleteProductResp, error) {
	l := logic.NewDeleteProductLogic(ctx, s.svcCtx)
	return l.DeleteProduct(in)
}

func (s *ProductServer) GetProduct(ctx context.Context, in *pb.GetProductByIdReq) (*pb.GetProductByIdResp, error) {
	l := logic.NewGetProductLogic(ctx, s.svcCtx)
	return l.GetProduct(in)
}

func (s *ProductServer) CreateProductSku(ctx context.Context, in *pb.CreateProductSkuReq) (*pb.CreateProductSkuResp, error) {
	l := logic.NewCreateProductSkuLogic(ctx, s.svcCtx)
	return l.CreateProductSku(in)
}

func (s *ProductServer) GetProductListSku(ctx context.Context, in *pb.GetProductSkuListReq) (*pb.GetProductSkuListResp, error) {
	l := logic.NewGetProductListSkuLogic(ctx, s.svcCtx)
	return l.GetProductListSku(in)
}

func (s *ProductServer) DeleteProductSku(ctx context.Context, in *pb.DeleteProductSkuReq) (*pb.DeleteProductSkuResp, error) {
	l := logic.NewDeleteProductSkuLogic(ctx, s.svcCtx)
	return l.DeleteProductSku(in)
}

func (s *ProductServer) GetProductSku(ctx context.Context, in *pb.GetProductSkuByIdReq) (*pb.GetProductSkuByIdResp, error) {
	l := logic.NewGetProductSkuLogic(ctx, s.svcCtx)
	return l.GetProductSku(in)
}
