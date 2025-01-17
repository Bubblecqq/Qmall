package repository

import (
	"QMall/marketing/cmd/domain/model"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type IActivityRepository interface {
	GetActivityById(id int64) (*model.Activity, error)
	DeleteActivityById(id, userId int64) error
}

type ActivityRepository struct {
	mysqlClient *gorm.DB
	redisClient *redis.Client
}

func (a *ActivityRepository) GetActivityById(id int64) (*model.Activity, error) {
	activity := &model.Activity{}
	fmt.Printf("正在获取活动表信息！\n")
	tx := a.mysqlClient.Model(&model.Activity{}).Find(&activity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Printf("活动信息获取成功！，活动内容如下：%v\n", activity)
	return activity, nil
}

func (a *ActivityRepository) DeleteActivityById(id, userId int64) error {
	activityById, err := a.GetActivityById(id)

	if err != nil {
		fmt.Printf("当前id：%v的活动不存在！\n", id)
		return err
	}
	activityById.IsDeleted = 1
	updateTime := time.Now()
	activityById.UpdateTime = &updateTime
	activityById.UpdateUser = userId
	tx := a.mysqlClient.Model(&model.Activity{}).Updates(activityById)
	if tx.Error != nil {
		fmt.Printf("当前id：%v的活动更新失败！原因见：%v\n", id, tx.Error)
		return tx.Error
	}
	return nil
}

func NewActivityRepository(mysqlClient *gorm.DB, redisClient *redis.Client) IActivityRepository {
	return &ActivityRepository{
		mysqlClient: mysqlClient,
		redisClient: redisClient,
	}
}
