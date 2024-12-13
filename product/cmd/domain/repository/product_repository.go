package repository

import (
	"QMall/product/cmd/domain/model"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type IProductRepository interface {
	IProductSkuRepository
	IProductAttributeRepository
	FindProduct(id int64) (*model.Product, error)
	GetProductList() []model.Product
	DeleteProduct(id int64) error
	CreateProduct(name string, product_type int32, category_id int32, starting_price float64, total_stock int64, main_picture string, remote_area_postage float64, single_buy_limit int32, remark string) error
	//UpdateProduct(product *model.Product)
	Page(length int32, pageIndex int32) (int64, []model.Product, error)
	CountNum() int64
	ShowProductDetail(id int64) (*model.Product, error)
}

type ProductRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (p *ProductRepository) ShowProductDetail(id int64) (*model.Product, error) {
	product := &model.Product{}
	err := p.mysqlClient.Model(&model.Product{}).Preload("Detail").Preload("PictureList").Find(&product, id).Error
	return product, err
}

func (p *ProductRepository) CountNum() int64 {

	var result int64
	p.mysqlClient.Model(&model.Product{}).Limit(-1).Offset(-1).Count(&result)
	return result
}

func (p *ProductRepository) Page(length int32, pageIndex int32) (int64, []model.Product, error) {
	pageProductList := make([]model.Product, length)
	var count int64
	if length > 0 && pageIndex > 0 {
		count = p.CountNum()
		if err := p.mysqlClient.Model(&model.Product{}).Limit(int(length)).Offset(int(pageIndex-1) * int(length)).Find(&pageProductList).Error; err != nil {
			fmt.Println("[ProductRepository] Query product error with:", err)
			return count, pageProductList, err
		}
	} else {
		return count, pageProductList, errors.New("参数有误，请重新输入！")
	}
	return count, pageProductList, nil
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

func (p *ProductRepository) CreateProduct(name string, product_type int32, category_id int32, starting_price float64, total_stock int64, main_picture string, remote_area_postage float64, single_buy_limit int32, remark string) error {
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
