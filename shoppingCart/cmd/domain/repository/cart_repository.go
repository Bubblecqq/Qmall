package repository

import (
	"QMall/shoppingCart/cmd/domain/model"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"
	"fmt"
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
	GetShoppingCartsByUserId(userId int64) ([]model.ShoppingCart, error)

	DeleteShoppingCartByUserId(userId int64) error

	GetTotalPriceByUserId(userId int64) float64
}

type ShoppingCartRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (s *ShoppingCartRepository) GetTotalPriceByUserId(userId int64) float64 {
	var totalPrice float64
	tx := s.mysqlClient.Model(&model.ShoppingCart{}).
		Select("Sum(shopping_cart.number * product_sku.sell_price) as total_price").
		Joins("left join product_sku on shopping_cart.product_sku_id = product_sku.id").
		Where("shopping_cart.user_id = ?", userId).
		Raw("").
		Scan(&totalPrice)

	if tx.Error != nil {
		fmt.Printf("当前用户Id:%v，购物车中不存在任意商品！\n", userId)
		return 0
	}

	return totalPrice
}

func (s *ShoppingCartRepository) DeleteShoppingCartByUserId(userId int64) error {
	shoppingCarts, err := s.GetShoppingCartsByUserId(userId)

	if len(shoppingCarts) <= 0 {
		fmt.Printf("当前用户Id：%v，购物车为空！\n", userId)
		return err
	}
	tx := s.mysqlClient.Model(&model.ShoppingCart{}).Delete(&shoppingCarts)
	if tx.Error != nil {
		fmt.Printf("当前用户Id：%v，购物车删除失败，失败原因：%v！\n", userId, tx.Error)
		return err
	}

	return tx.Error
}

func (s *ShoppingCartRepository) GetShoppingCartsByUserId(userId int64) ([]model.ShoppingCart, error) {
	var carts []model.ShoppingCart
	tx := s.mysqlClient.Model(&model.ShoppingCart{}).Where("user_id= ?", userId).Find(&carts)
	return carts, tx.Error
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
