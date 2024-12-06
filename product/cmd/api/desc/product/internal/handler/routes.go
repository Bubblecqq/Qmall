// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	product "goZeroDemo4/product/cmd/api/desc/product/internal/handler/product"
	"goZeroDemo4/product/cmd/api/desc/product/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 创建商品
				Method:  http.MethodPost,
				Path:    "/createProduct",
				Handler: product.CreateProductHandler(serverCtx),
			},
			{
				// 根据id删除商品
				Method:  http.MethodPost,
				Path:    "/deleteProduct",
				Handler: product.DeleteProductHandler(serverCtx),
			},
			{
				// 根据id获取商品
				Method:  http.MethodPost,
				Path:    "/getProductById",
				Handler: product.GetProductByIdHandler(serverCtx),
			},
			{
				// 获取商品列表
				Method:  http.MethodPost,
				Path:    "/getProducts",
				Handler: product.GetProductsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/product/v1"),
	)
}
