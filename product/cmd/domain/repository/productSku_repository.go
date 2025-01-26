package repository

import (
	"QMall/product/cmd/domain/model"
	"QMall/product/cmd/rpc/product/pb"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type IProductSkuRepository interface {
	ProductSkuRepo
	FindProductSku(id int64) (*model.ProductSku, error)
	GetProductSkuList() []model.ProductSku
	DeleteProductSku(id int64) error
	CreateProductSku(product_id int32, attribute_symbol_list, name string, sell_price, cost_price float64, stock int64, stock_warn int64) error
	UpdateProductSku(in *pb.UpdateProductSkuReq) error
	//UpdateProduct(product *model.Product)
	UpdateProductSkuBySkuId(in *pb.UpdateProductSkuBySkuIdReq) (*model.ProductSku, error)

	GetProductSkuStock(skuId int64) (int64, error)
	UpdateProductSkuStock(skuId, quantity int64) error
}

func (p *ProductRepository) UpdateProductSkuStock(skuId, quantity int64) error {
	tx := p.mysqlClient.Begin()
	sku := &model.ProductSku{}
	find := tx.Model(&model.ProductSku{}).Find(&sku, skuId)
	if find.Error != nil {
		return find.Error
	}
	sku.Stock -= quantity
	sku.UpdateTime = time.Now()
	updates := tx.Model(&model.ProductSku{}).Where("id=?", skuId).Updates(&sku)

	if updates.Error != nil {
		tx.Rollback()
		return updates.Error
	}
	tx.Commit()
	return nil
}

func (p *ProductRepository) GetProductSkuStock(skuId int64) (int64, error) {
	sku := &model.ProductSku{}

	tx := p.mysqlClient.Model(&model.ProductSku{}).Find(&sku, skuId)

	if tx.Error != nil {
		return 0, tx.Error
	}
	return sku.Stock, nil
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
func (p *ProductRepository) UpdateProductSkuBySkuId(in *pb.UpdateProductSkuBySkuIdReq) (*model.ProductSku, error) {

	sku := &model.ProductSku{}

	tx := p.mysqlClient.Begin()
	if tx.Error != nil {
		return rollbackAndReturnError(tx, tx.Error)
	}

	//err := tx.Model(&model.ProductSku{}).Find(&sku, in.SkuId).Error
	err := tx.Model(&model.ProductSku{}).Set("gorm:query_option", "FOR UPDATE").Where("id=?", in.SkuId).First(&sku).Error
	if err != nil {
		return rollbackAndReturnError(tx, err)
	}
	if sku.Id <= 0 {
		return rollbackAndReturnError(tx, gorm.ErrRecordNotFound)
	}

	if sku.Stock < in.Stock {
		return rollbackAndReturnError(tx, errors.New("库存不足！"))
	}
	sku.Stock -= in.Stock
	sku.UpdateTime = time.Now()

	err = tx.Save(&sku).Error
	if err != nil {
		return rollbackAndReturnError(tx, err)
	}
	if err := tx.Commit().Error; err != nil {
		return rollbackAndReturnError(tx, err)
	}
	fmt.Printf("成功更新商品SKU库存，SKU ID: %d，新库存: %d\n", sku.Id, sku.Stock)
	return sku, nil
}

func rollbackAndReturnError(tx *gorm.DB, err error) (*model.ProductSku, error) {
	if tx.Error != nil {
		err = tx.Error
	}
	if rErr := tx.Rollback().Error; rErr != nil {
		fmt.Printf("回滚事务失败！原因：%v\n", rErr)
	}
	return nil, err
}
