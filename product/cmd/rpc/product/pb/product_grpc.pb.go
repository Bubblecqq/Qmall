// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: product.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Product_CreateProduct_FullMethodName           = "/pb.product/CreateProduct"
	Product_GetProductList_FullMethodName          = "/pb.product/GetProductList"
	Product_DeleteProduct_FullMethodName           = "/pb.product/DeleteProduct"
	Product_GetProduct_FullMethodName              = "/pb.product/GetProduct"
	Product_PageIndex_FullMethodName               = "/pb.product/PageIndex"
	Product_ShowProductDetail_FullMethodName       = "/pb.product/ShowProductDetail"
	Product_UpdateProductSku_FullMethodName        = "/pb.product/UpdateProductSku"
	Product_CreateProductSku_FullMethodName        = "/pb.product/CreateProductSku"
	Product_GetProductListSku_FullMethodName       = "/pb.product/GetProductListSku"
	Product_DeleteProductSku_FullMethodName        = "/pb.product/DeleteProductSku"
	Product_GetProductSku_FullMethodName           = "/pb.product/GetProductSku"
	Product_UpdateProductSkuBySkuId_FullMethodName = "/pb.product/UpdateProductSkuBySkuId"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error)
	GetProductList(ctx context.Context, in *GetProductListReq, opts ...grpc.CallOption) (*GetProductListResp, error)
	DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error)
	GetProduct(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdResp, error)
	PageIndex(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*PageResp, error)
	ShowProductDetail(ctx context.Context, in *ShowProductDetailReq, opts ...grpc.CallOption) (*ShowProductDetailResp, error)
	UpdateProductSku(ctx context.Context, in *UpdateProductSkuReq, opts ...grpc.CallOption) (*UpdateProductSkuResp, error)
	CreateProductSku(ctx context.Context, in *CreateProductSkuReq, opts ...grpc.CallOption) (*CreateProductSkuResp, error)
	GetProductListSku(ctx context.Context, in *GetProductSkuListReq, opts ...grpc.CallOption) (*GetProductSkuListResp, error)
	DeleteProductSku(ctx context.Context, in *DeleteProductSkuReq, opts ...grpc.CallOption) (*DeleteProductSkuResp, error)
	GetProductSku(ctx context.Context, in *GetProductSkuByIdReq, opts ...grpc.CallOption) (*GetProductSkuByIdResp, error)
	UpdateProductSkuBySkuId(ctx context.Context, in *UpdateProductSkuBySkuIdReq, opts ...grpc.CallOption) (*UpdateProductSkuBySkuIdResp, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProductResp)
	err := c.cc.Invoke(ctx, Product_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetProductList(ctx context.Context, in *GetProductListReq, opts ...grpc.CallOption) (*GetProductListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductListResp)
	err := c.cc.Invoke(ctx, Product_GetProductList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProductResp)
	err := c.cc.Invoke(ctx, Product_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetProduct(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductByIdResp)
	err := c.cc.Invoke(ctx, Product_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) PageIndex(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*PageResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PageResp)
	err := c.cc.Invoke(ctx, Product_PageIndex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ShowProductDetail(ctx context.Context, in *ShowProductDetailReq, opts ...grpc.CallOption) (*ShowProductDetailResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShowProductDetailResp)
	err := c.cc.Invoke(ctx, Product_ShowProductDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) UpdateProductSku(ctx context.Context, in *UpdateProductSkuReq, opts ...grpc.CallOption) (*UpdateProductSkuResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductSkuResp)
	err := c.cc.Invoke(ctx, Product_UpdateProductSku_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CreateProductSku(ctx context.Context, in *CreateProductSkuReq, opts ...grpc.CallOption) (*CreateProductSkuResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProductSkuResp)
	err := c.cc.Invoke(ctx, Product_CreateProductSku_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetProductListSku(ctx context.Context, in *GetProductSkuListReq, opts ...grpc.CallOption) (*GetProductSkuListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductSkuListResp)
	err := c.cc.Invoke(ctx, Product_GetProductListSku_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteProductSku(ctx context.Context, in *DeleteProductSkuReq, opts ...grpc.CallOption) (*DeleteProductSkuResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProductSkuResp)
	err := c.cc.Invoke(ctx, Product_DeleteProductSku_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetProductSku(ctx context.Context, in *GetProductSkuByIdReq, opts ...grpc.CallOption) (*GetProductSkuByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductSkuByIdResp)
	err := c.cc.Invoke(ctx, Product_GetProductSku_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) UpdateProductSkuBySkuId(ctx context.Context, in *UpdateProductSkuBySkuIdReq, opts ...grpc.CallOption) (*UpdateProductSkuBySkuIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductSkuBySkuIdResp)
	err := c.cc.Invoke(ctx, Product_UpdateProductSkuBySkuId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility.
type ProductServer interface {
	CreateProduct(context.Context, *CreateProductReq) (*CreateProductResp, error)
	GetProductList(context.Context, *GetProductListReq) (*GetProductListResp, error)
	DeleteProduct(context.Context, *DeleteProductReq) (*DeleteProductResp, error)
	GetProduct(context.Context, *GetProductByIdReq) (*GetProductByIdResp, error)
	PageIndex(context.Context, *PageReq) (*PageResp, error)
	ShowProductDetail(context.Context, *ShowProductDetailReq) (*ShowProductDetailResp, error)
	UpdateProductSku(context.Context, *UpdateProductSkuReq) (*UpdateProductSkuResp, error)
	CreateProductSku(context.Context, *CreateProductSkuReq) (*CreateProductSkuResp, error)
	GetProductListSku(context.Context, *GetProductSkuListReq) (*GetProductSkuListResp, error)
	DeleteProductSku(context.Context, *DeleteProductSkuReq) (*DeleteProductSkuResp, error)
	GetProductSku(context.Context, *GetProductSkuByIdReq) (*GetProductSkuByIdResp, error)
	UpdateProductSkuBySkuId(context.Context, *UpdateProductSkuBySkuIdReq) (*UpdateProductSkuBySkuIdResp, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServer struct{}

func (UnimplementedProductServer) CreateProduct(context.Context, *CreateProductReq) (*CreateProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServer) GetProductList(context.Context, *GetProductListReq) (*GetProductListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductList not implemented")
}
func (UnimplementedProductServer) DeleteProduct(context.Context, *DeleteProductReq) (*DeleteProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServer) GetProduct(context.Context, *GetProductByIdReq) (*GetProductByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServer) PageIndex(context.Context, *PageReq) (*PageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageIndex not implemented")
}
func (UnimplementedProductServer) ShowProductDetail(context.Context, *ShowProductDetailReq) (*ShowProductDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowProductDetail not implemented")
}
func (UnimplementedProductServer) UpdateProductSku(context.Context, *UpdateProductSkuReq) (*UpdateProductSkuResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductSku not implemented")
}
func (UnimplementedProductServer) CreateProductSku(context.Context, *CreateProductSkuReq) (*CreateProductSkuResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductSku not implemented")
}
func (UnimplementedProductServer) GetProductListSku(context.Context, *GetProductSkuListReq) (*GetProductSkuListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductListSku not implemented")
}
func (UnimplementedProductServer) DeleteProductSku(context.Context, *DeleteProductSkuReq) (*DeleteProductSkuResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductSku not implemented")
}
func (UnimplementedProductServer) GetProductSku(context.Context, *GetProductSkuByIdReq) (*GetProductSkuByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductSku not implemented")
}
func (UnimplementedProductServer) UpdateProductSkuBySkuId(context.Context, *UpdateProductSkuBySkuIdReq) (*UpdateProductSkuBySkuIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductSkuBySkuId not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}
func (UnimplementedProductServer) testEmbeddedByValue()                 {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	// If the following call pancis, it indicates UnimplementedProductServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateProduct(ctx, req.(*CreateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetProductList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProductList(ctx, req.(*GetProductListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteProduct(ctx, req.(*DeleteProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProduct(ctx, req.(*GetProductByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_PageIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).PageIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_PageIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).PageIndex(ctx, req.(*PageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ShowProductDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowProductDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ShowProductDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ShowProductDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ShowProductDetail(ctx, req.(*ShowProductDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_UpdateProductSku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductSkuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).UpdateProductSku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_UpdateProductSku_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).UpdateProductSku(ctx, req.(*UpdateProductSkuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CreateProductSku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductSkuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateProductSku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CreateProductSku_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateProductSku(ctx, req.(*CreateProductSkuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetProductListSku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductSkuListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProductListSku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetProductListSku_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProductListSku(ctx, req.(*GetProductSkuListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteProductSku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductSkuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteProductSku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DeleteProductSku_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteProductSku(ctx, req.(*DeleteProductSkuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetProductSku_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductSkuByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProductSku(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetProductSku_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProductSku(ctx, req.(*GetProductSkuByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_UpdateProductSkuBySkuId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductSkuBySkuIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).UpdateProductSkuBySkuId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_UpdateProductSkuBySkuId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).UpdateProductSkuBySkuId(ctx, req.(*UpdateProductSkuBySkuIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _Product_CreateProduct_Handler,
		},
		{
			MethodName: "GetProductList",
			Handler:    _Product_GetProductList_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Product_DeleteProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Product_GetProduct_Handler,
		},
		{
			MethodName: "PageIndex",
			Handler:    _Product_PageIndex_Handler,
		},
		{
			MethodName: "ShowProductDetail",
			Handler:    _Product_ShowProductDetail_Handler,
		},
		{
			MethodName: "UpdateProductSku",
			Handler:    _Product_UpdateProductSku_Handler,
		},
		{
			MethodName: "CreateProductSku",
			Handler:    _Product_CreateProductSku_Handler,
		},
		{
			MethodName: "GetProductListSku",
			Handler:    _Product_GetProductListSku_Handler,
		},
		{
			MethodName: "DeleteProductSku",
			Handler:    _Product_DeleteProductSku_Handler,
		},
		{
			MethodName: "GetProductSku",
			Handler:    _Product_GetProductSku_Handler,
		},
		{
			MethodName: "UpdateProductSkuBySkuId",
			Handler:    _Product_UpdateProductSkuBySkuId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
