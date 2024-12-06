package repository

import (
	"github.com/redis/go-redis/v9"
	"goZeroDemo4/product/cmd/domain/model"
	"gorm.io/gorm"
)

type IProductSkuRepository interface {
	FindOne(id int64) (*model.ProductSku, error)
	GetProductSkuList() []model.ProductSku
	DeleteOne(id int64) error
	CreateProductSku(product_id int32, attribute_symbol_list, name string, sell_price, cost_price float64, stock int32, stock_warn int) error
	//UpdateProduct(product *model.Product)
}

type ProductSkuRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (p *ProductSkuRepository) FindOne(id int64) (*model.ProductSku, error) {
	productSku := &model.ProductSku{}
	err := p.mysqlClient.Model(&model.ProductSku{}).Find(&productSku).Error
	return productSku, err
}

func (p *ProductSkuRepository) GetProductSkuList() []model.ProductSku {
	var productSkuList []model.ProductSku
	p.mysqlClient.Model(&model.ProductSku{}).Find(&productSkuList)
	return productSkuList
}

func (p *ProductSkuRepository) DeleteOne(id int64) error {
	return p.mysqlClient.Model(&model.ProductSku{}).Delete(&model.ProductSku{}, id).Error
}

func (p *ProductSkuRepository) CreateProductSku(product_id int32, attribute_symbol_list, name string, sell_price, cost_price float64, stock int32, stock_warn int) error {
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

func NewProductSkuRepository(mysqlClient *gorm.DB, redisClient *redis.Client) IProductSkuRepository {
	return &ProductSkuRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
