// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type CreateProductReq struct {
	Name              string  `json:"name"`
	ProductType       int32   `json:"product_type"`
	CategoryId        int32   `json:"category_id"`
	StartingPrice     float64 `json:"starting_price"`
	TotalStock        int64   `json:"total_stock"`
	MainPicture       string  `json:"main_picture"`
	RemoteAreaPostage float64 `json:"remote_area_postage"`
	SingleBuyLimit    int32   `json:"single_buy_limit"`
	Remark            string  `json:"remark"`
}

type CreateProductResp struct {
}

type CreateProductSkuReq struct {
	Name                string  `json:"name"`
	ProductId           int32   `json:"product_id"`
	AttributeSymbolList string  `json:"attribute_symbol_list"`
	SellPrice           float64 `json:"sell_price"`
	CostPrice           float64 `json:"cost_price"`
	Stock               int64   `json:"stock"`
	StockWarn           int64   `json:"stock_warn"`
}

type CreateProductSkuResp struct {
}

type DeleteProductReq struct {
	Id int64 `json:"id"`
}

type DeleteProductResp struct {
}

type DeleteProductSkuReq struct {
	Id int64 `json:"id"`
}

type DeleteProductSkuResp struct {
}

type DetailProduct struct {
	Id                int64    `json:"id"`
	Name              string   `json:"name"`
	ProductType       int32    `json:"product_type"`
	CategoryId        int32    `json:"category_id"`
	StartingPrice     float64  `json:"starting_price"`
	TotalStock        int64    `json:"total_stock"`
	MainPicture       string   `json:"main_picture"`
	RemoteAreaPostage float64  `json:"remote_area_postage"`
	SingleBuyLimit    int32    `json:"single_buy_limit"`
	Remark            string   `json:"remark"`
	CreateTime        string   `json:"create_time"`
	UpdateTime        string   `json:"update_time"`
	IsEnabled         bool     `json:"is_enable"`
	CreateUser        string   `json:"create_user"`
	UpdateUser        string   `json:"update_user"`
	Detail            string   `json:"detail"`
	PictureList       []string `json:"pictureList"`
}

type GetProductByIdReq struct {
	Id int64 `json:"id"`
}

type GetProductByIdResp struct {
	Product Product `json:"product"`
}

type GetProductListReq struct {
}

type GetProductListResp struct {
	ProductList []Product `json:"products"`
}

type GetProductSkuByIdReq struct {
	Id int64 `json:"id"`
}

type GetProductSkuByIdResp struct {
	ProductSku ProductSku `json:"productSku"`
}

type GetProductSkuListReq struct {
}

type GetProductSkuListResp struct {
	ProductSkuList []ProductSku `json:"productSkus"`
}

type PageReq struct {
	Length    int32 `json:"length"`
	PageIndex int32 `json:"pageIndex"`
}

type PageResp struct {
	ResponsePageIndex []Product `json:"product"`
}

type Product struct {
	Id                int64   `json:"id"`
	Name              string  `json:"name"`
	ProductType       int32   `json:"product_type"`
	CategoryId        int32   `json:"category_id"`
	StartingPrice     float64 `json:"starting_price"`
	TotalStock        int64   `json:"total_stock"`
	MainPicture       string  `json:"main_picture"`
	RemoteAreaPostage float64 `json:"remote_area_postage"`
	SingleBuyLimit    int32   `json:"single_buy_limit"`
	Remark            string  `json:"remark"`
	CreateTime        string  `json:"create_time"`
	UpdateTime        string  `json:"update_time"`
	IsEnabled         bool    `json:"is_enable"`
	CreateUser        int64   `json:"create_user"`
	UpdateUser        int64   `json:"update_user"`
}

type ProductSku struct {
	Id                  int64   `json:"id"`
	Name                string  `json:"name"`
	ProductId           int32   `json:"product_id"`
	AttributeSymbolList string  `json:"attribute_symbol_list"`
	SellPrice           float64 `json:"sell_price"`
	CostPrice           float64 `json:"cost_price"`
	Stock               int64   `json:"stock"`
	StockWarn           int64   `json:"stock_warn"`
	CreateTime          string  `json:"create_time"`
	UpdateTime          string  `json:"update_time"`
	IsEnabled           int32   `json:"is_enable"`
	CreateUser          int64   `json:"create_user"`
	UpdateUser          int64   `json:"update_user"`
}

type ShowProductDetailReq struct {
	Id int64 `json:"id"`
}

type ShowProductDetailResp struct {
	DetailProduct DetailProduct `json:"detail_product"`
}
