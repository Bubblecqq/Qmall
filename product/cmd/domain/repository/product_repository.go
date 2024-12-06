package repository

import (
	"github.com/redis/go-redis/v9"
	"goZeroDemo4/product/cmd/domain/model"
	"gorm.io/gorm"
)

type IProductRepository interface {
	IProductSkuRepository
	FindProduct(id int64) (*model.Product, error)
	GetProductList() []model.Product
	DeleteProduct(id int64) error
	CreateProduct(name string, product_type int32, category_id int32, starting_price float64, total_stock int32, main_picture string, remote_area_postage float64, single_buy_limit int32, remark string) error
	//UpdateProduct(product *model.Product)
}

type ProductRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (p *ProductRepository) FindProduct(id int64) (*model.Product, error) {
	product := &model.Product{}
	err := p.mysqlClient.Model(&model.Product{}).Find(&product, id).Error
	return product, err
}

func (p *ProductRepository) GetProductList() []model.Product {
	var productList []model.Product
	p.mysqlClient.Model(&model.Product{}).Find(&productList)
	return productList
}

func (p *ProductRepository) DeleteProduct(id int64) error {
	return p.mysqlClient.Delete(&model.Product{}, id).Error
}

func (p *ProductRepository) CreateProduct(name string, product_type int32, category_id int32, starting_price float64, total_stock int32, main_picture string, remote_area_postage float64, single_buy_limit int32, remark string) error {
	product := &model.Product{
		Name:              name,
		ProductType:       product_type,
		CategoryId:        category_id,
		StartingPrice:     starting_price,
		TotalStock:        total_stock,
		MainPicture:       main_picture,
		RemoteAreaPostage: remote_area_postage,
		SingleBuyLimit:    single_buy_limit,
		Remark:            remark,
	}
	err := p.mysqlClient.Model(&model.Product{}).Create(&product).Error
	return err
}

func NewProductRepository(mysqlClient *gorm.DB, redisClient *redis.Client) IProductRepository {
	return &ProductRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
