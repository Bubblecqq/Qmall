package repository

import (
	"QMall/shoppingCart/cmd/domain/model"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type IShoppingCartRepository interface {
	InCreaseShoppingCart(number, product_id, product_sku_id, user_id int64, product_name, product_main_picture string) (*model.ShoppingCart, error)
	UpdateShoppingCart(in *pb.UpdateCartReq) (*model.ShoppingCart, error)
	FindShoppingCart(user_id int64) (*model.ShoppingCart, error)
}

type ShoppingCartRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (s *ShoppingCartRepository) InCreaseShoppingCart(number, product_id, product_sku_id, user_id int64, product_name, product_main_picture string) (*model.ShoppingCart, error) {
	var updateTime *time.Time
	//updateTime2 := time.Now()
	shoppingCart := &model.ShoppingCart{
		UserId:             user_id,
		ProductId:          product_id,
		ProductSkuId:       product_sku_id,
		ProductName:        product_name,
		ProductMainPicture: product_main_picture,
		Number:             number,
		CreateTime:         time.Now(),
		UpdateTime:         updateTime,
		CreateUser:         user_id,
		//UpdateUser:         user_id,
	}
	return shoppingCart, s.mysqlClient.Model(&model.ShoppingCart{}).Create(&shoppingCart).Error
}

func (s *ShoppingCartRepository) UpdateShoppingCart(in *pb.UpdateCartReq) (*model.ShoppingCart, error) {
	updateTime := time.Now()
	shoppingCart := &model.ShoppingCart{
		Id:                 in.Id,
		UserId:             in.UserId,
		ProductId:          in.ProductId,
		ProductSkuId:       in.ProductSkuId,
		ProductName:        in.ProductName,
		ProductMainPicture: in.ProductMainPicture,
		UpdateTime:         &updateTime,
		UpdateUser:         in.UserId,
		IsDeleted:          in.IsDelete,
	}

	//err := s.mysqlClient.Save(&shoppingCart).Error

	return shoppingCart, s.mysqlClient.Save(&shoppingCart).Error
}

func (s *ShoppingCartRepository) FindShoppingCart(user_id int64) (*model.ShoppingCart, error) {
	shoppingCart := &model.ShoppingCart{}
	err := s.mysqlClient.Model(&model.ShoppingCart{}).Where("user_id = ?", user_id).Find(&shoppingCart).Error
	return shoppingCart, err
}

func NewShoppingCartRepository(mysqlClient *gorm.DB, redisClient *redis.Client) IShoppingCartRepository {
	return &ShoppingCartRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
