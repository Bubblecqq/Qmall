package model

import "time"

type Activity struct {
	ID                int64      `gorm:"primaryKey;column:id" json:"id"`
	Name              string     `gorm:"column:name" json:"name"`
	ActivityStartTime time.Time  `gorm:"column:activity_start_time" json:"activity_start_time"`
	ActivityEndTime   time.Time  `gorm:"column:activity_end_time" json:"activity_end_time"`
	IsOnline          int32      `gorm:"column:is_online" json:"is_online"`
	CreateUser        int64      `gorm:"column:create_user" json:"create_user"`
	CreateTime        time.Time  `gorm:"column:create_time" json:"create_time"`
	UpdateUser        int64      `gorm:"column:update_user" json:"update_user"`
	UpdateTime        *time.Time `gorm:"column:update_time" json:"update_time"`
	IsDeleted         int32      `gorm:"column:is_deleted" json:"is_deleted"`
}

func (t *Activity) TableName() string {
	return "activity"
}
