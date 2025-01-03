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
	FindShoppingCart(in *pb.FindCartReq) (*model.ShoppingCart, error)
	GetShoppingCartById(id int64) (*model.ShoppingCart, error)
	GetShoppingCarts() ([]model.ShoppingCart, error)
}

type ShoppingCartRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (s *ShoppingCartRepository) GetShoppingCarts() ([]model.ShoppingCart, error) {
	var carts []model.ShoppingCart
	tx := s.mysqlClient.Model(&model.ShoppingCart{}).Find(&carts)
	return carts, tx.Error
}

func (s *ShoppingCartRepository) GetShoppingCartById(id int64) (*model.ShoppingCart, error) {
	shoppingCart := &model.ShoppingCart{}
	return shoppingCart, s.mysqlClient.Model(&model.ShoppingCart{}).Find(&shoppingCart, id).Error
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
	shoppingCartById, _ := s.GetShoppingCartById(in.GetId())
	updateTime := time.Now()
	shoppingCartById.UpdateTime = &updateTime
	shoppingCartById.UserId = in.UserId
	shoppingCartById.UpdateUser = in.UserId
	shoppingCartById.IsDeleted = in.IsDelete
	shoppingCartById.ProductId = in.ProductId
	shoppingCartById.ProductSkuId = in.ProductSkuId
	shoppingCartById.ProductName = in.ProductName
	shoppingCartById.ProductMainPicture = in.ProductMainPicture
	//shoppingCart := &model.ShoppingCart{
	//	Id:                 in.Id,
	//	UserId:             in.UserId,
	//	ProductId:          in.ProductId,
	//	ProductSkuId:       in.ProductSkuId,
	//	ProductName:        in.ProductName,
	//	ProductMainPicture: in.ProductMainPicture,
	//	UpdateTime:         &updateTime,
	//	UpdateUser:         in.UserId,
	//	IsDeleted:          in.IsDelete,
	//}
	//err := s.mysqlClient.Save(&shoppingCart).Error

	return shoppingCartById, s.mysqlClient.Save(&shoppingCartById).Error
}

func (s *ShoppingCartRepository) FindShoppingCart(in *pb.FindCartReq) (*model.ShoppingCart, error) {
	shoppingCart := &model.ShoppingCart{}
	err := s.mysqlClient.Model(&model.ShoppingCart{}).Where("user_id = ?", in.UserId).Find(&shoppingCart, in.Id).Error
	return shoppingCart, err
}

func NewShoppingCartRepository(mysqlClient *gorm.DB, redisClient *redis.Client) IShoppingCartRepository {
	return &ShoppingCartRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
