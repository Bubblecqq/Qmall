package repository

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"goZeroDemo4/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

var mysqlConnect = "root:000000@tcp(192.168.23.233:3306)/user_center?charset=utf8mb4&parseTime=True&loc=Local"

//var dialector = mysql.New(mysql.Config{
//	DSN:               mysqlConnect,
//	DefaultStringSize: 256,
//})

func TestName(t *testing.T) {
	//mysqlConnect := "root:000000@tcp(192.168.23.233:3306)/user_center?charset=utf8mb4&parseTime=True&loc=Local"
	//dialector := mysql.New(mysql.Config{
	//	DSN:               mysqlConnect,
	//	DefaultStringSize: 256,
	//})
	dsn := fmt.Sprintf(common.MysqlConnect, "root", "000000", "192.168.23.233:3306", "user_center", "utf8mb4")
	var dialector = mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	})
	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	userRepository := NewUserRepository(db, nil)
	userRepository.GetUserById(1)
	users, _ := userRepository.GetUsers()
	fmt.Println(users)
	//userRepository.AddUser(1, "120", "000000", 1)

}

func TestRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.23.233:6379",
		Password: "123456",
		PoolSize: 20,
	})
	redis := NewUserRepository(nil, client)
	redis.SetUserToken("uuid1111", []byte("tokenxxxx"), time.Duration(1)*time.Hour)
	token := redis.GetUserToken("uuid1111")
	fmt.Println("token>", token)
}

func TestLoginUser(t *testing.T) {
	dsn := fmt.Sprintf(common.MysqlConnect, "root", "000000", "192.168.23.233:3306", "user_center", "utf8mb4")
	var dialector = mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	})
	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	userRepository := NewUserRepository(db, nil)
	userRepository.Login(0, "12306", 0, "6666")
}
