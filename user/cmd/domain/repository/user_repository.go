package repository

import (
	"QMall/common"
	"QMall/user/cmd/domain/model"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"time"
)

type IUserRepository interface {
	GetUserById(id int64) (*model.User, error)
	AddUser(clientId int32, phone, password string, systemId int32) error
	GetUsers() ([]model.User, error)
	DeleteUserById(id int64) error
	SetUserToken(key string, val []byte, timeTTL time.Duration)
	GetUserToken(key string) string
	Login(clientId int32, phone string, systemId int32, verificationCode string) (*model.User, error)
}

type UserRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (u *UserRepository) Login(clientId int32, phone string, systemId int32, verificationCode string) (*model.User, error) {

	user := &model.User{}
	if clientId == 0 && systemId == 0 && verificationCode == "6666" {
		u.mysqlClient.Where("phone = ?", phone).Find(&user)
		//没找到就注册
		fmt.Println("user>", user)
		if user.Id == 0 {
			user.Phone = phone
			u.mysqlClient.Create(&user)
		}
		return user, nil
	} else {
		return user, errors.New("参数不正确！")
	}
}

func (u *UserRepository) SetUserToken(key string, val []byte, timeTTL time.Duration) {
	intKey := common.ToInput(key)
	binKey := common.ConverToBinary(intKey)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>", binKey)
	u.redisClient.Set(context.Background(), key, val, timeTTL)
}

func (u *UserRepository) GetUserToken(key string) string {
	result, err := u.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		log.Println("GetUserToken err:", err)
	}
	return result
}

func (u *UserRepository) GetUserById(id int64) (*model.User, error) {
	user := &model.User{
		BaseModel: model.BaseModel{Id: id},
	}
	err := u.mysqlClient.Model(&model.User{}).Where("id = ?", id).Find(&user).Error
	fmt.Printf("[UserRepository] Successfully to Get User By id=%d,user=%v\n", id, user)
	return user, err
}

func (u *UserRepository) AddUser(clientId int32, phone, password string, systemId int32) error {

	user := &model.User{
		BaseModel: model.BaseModel{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
		ClientId:      clientId,
		Phone:         phone,
		Password:      password,
		SystemId:      systemId,
		LastLoginTime: time.Now(),
	}
	err := u.mysqlClient.Model(&model.User{}).Create(&user).Error
	if err == nil {
		fmt.Printf("[UserRepository] Successfully to Create User with %v\n", user)
	}
	return err
}

func (u *UserRepository) GetUsers() ([]model.User, error) {
	var userList []model.User
	err := u.mysqlClient.Model(&model.User{}).Find(&userList).Error
	if err == nil {
		fmt.Printf("[UserRepository] Successfully to Get User List Rows: %v\n", len(userList))
	}
	return userList, err
}

func (u *UserRepository) DeleteUserById(id int64) error {
	return u.mysqlClient.Delete(&model.User{}, id).Error
}

func NewUserRepository(mysqlClient *gorm.DB, redisClient *redis.Client) IUserRepository {
	return &UserRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
