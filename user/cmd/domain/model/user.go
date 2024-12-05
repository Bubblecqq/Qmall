package model

import (
	"gorm.io/gorm"
	"time"
)

// User 结构体对应MySQL中的user表
type User struct {
	BaseModel
	Account         string    `gorm:"column:account;default:''"`
	Nickname        string    `gorm:"column:nickname;default:'momo'"`
	Avatar          string    `gorm:"column:avatar;default:''"`
	Gender          int32     `gorm:"column:gender;default:3"`
	Phone           string    `gorm:"column:phone"`
	Password        string    `gorm:"column:password;default:'123456'"`
	IDCard          string    `gorm:"column:id_card;default:''"`
	Source          int32     `gorm:"column:source"`
	SystemId        int32     `gorm:"column:system_id;default:1"`
	ClientId        int32     `gorm:"column:client_id;default:1"`
	UnionId         string    `gorm:"column:union_id;default:1"`
	IsEnable        int32     `gorm:"column:is_enable;default:1"`
	LastLoginTime   time.Time `gorm:"column:last_login_time"`
	CreateUser      int64     `gorm:"column:create_user"`
	UpdateUser      int64     `gorm:"column:update_user"`
	Token           string    `gorm:"-"`
	SessionId       string    `gorm:"-"`
	TokenExpireTime string    `gorm:"-"`
}

func (u *User) TableName() string {
	return "user"
}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.LastLoginTime.IsZero() {
		u.LastLoginTime = time.Now()
	}
	if u.CreateTime.IsZero() {
		u.CreateTime = time.Now()
	}
	if u.UpdateTime.IsZero() {
		u.UpdateTime = time.Now()
	}
	return nil
}
