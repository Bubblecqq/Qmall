package repository

import (
	"QMall/product/cmd/domain/model"
	"QMall/product/cmd/rpc/product/pb"
	"gorm.io/gorm"
)

type IProductSkuRepository interface {
	FindProductSku(id int64) (*model.ProductSku, error)
	GetProductSkuList() []model.ProductSku
	DeleteProductSku(id int64) error
	CreateProductSku(product_id int32, attribute_symbol_list, name string, sell_price, cost_price float64, stock int64, stock_warn int64) error
	UpdateProductSku(in *pb.UpdateProductSkuReq) error
	//UpdateProduct(product *model.Product)
}

//type ProductSkuRepository struct {
//	mysqlClient *gorm.DB
//	redisClient *redis.Client
//}

func (p *ProductRepository) FindProductSku(id int64) (*model.ProductSku, error) {
	productSku := &model.ProductSku{}
	tx := p.mysqlClient.Model(&model.ProductSku{}).Find(&productSku, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if productSku.Id == 0 || productSku.Name == "" {
		return nil, gorm.ErrRecordNotFound
	}
	return productSku, nil
}

func (p *ProductRepository) GetProductSkuList() []model.ProductSku {
	var productSkuList []model.ProductSku
	p.mysqlClient.Model(&model.ProductSku{}).Find(&productSkuList)
	return productSkuList
}

func (p *ProductRepository) DeleteProductSku(id int64) error {
	return p.mysqlClient.Model(&model.ProductSku{}).Delete(&model.ProductSku{}, id).Error
}

func (p *ProductRepository) CreateProductSku(product_id int32, attribute_symbol_list, name string, sell_price, cost_price float64, stock int64, stock_warn int64) error {
	productSku := &model.ProductSku{
		ProductId:           product_id,
		AttributeSymbolList: attribute_symbol_list,
		Name:                name,
		SellPrice:           sell_price,
		CostPrice:           cost_price,
		Stock:               stock,
		StockWarn:           stock_warn,
	}
	return p.mysqlClient.Model(&model.ProductSku{}).Create(&productSku).Error
}

//func NewProductSkuRepository(mysqlClient *gorm.DB, redisClient *redis.Client) IProductSkuRepository {
//	return &ProductSkuRepository{
//		mysqlClient: mysqlClient,
//		redisClient: redisClient,
//	}
//}

func (p *ProductRepository) UpdateProductSku(in *pb.UpdateProductSkuReq) error {
	sku := in.GetProductSku()
	err := p.mysqlClient.Model(&model.ProductSku{}).Where("id= ?", sku.Id).Update("stock", sku.Stock).Error
	return err
}
