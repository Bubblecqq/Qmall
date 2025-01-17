package model

import "time"

type ActivityTime struct {
	ID         int64      `gorm:"primaryKey;column:id" json:"id"`
	ActivityID int64      `gorm:"column:activity_id" json:"activity_id"`
	TimeName   string     `gorm:"column:time_name" json:"time_name"`
	StartTime  time.Time  `gorm:"column:start_time" json:"start_time"`
	EndTime    time.Time  `gorm:"column:end_time" json:"end_time"`
	IsEnable   int32      `gorm:"column:is_enable" json:"is_enable"`
	CreateUser int64      `gorm:"column:create_user" json:"create_user"`
	CreateTime time.Time  `gorm:"column:create_time" json:"create_time"`
	UpdateUser int64      `gorm:"column:update_user" json:"update_user"`
	UpdateTime *time.Time `gorm:"column:update_time" json:"update_time"`
	IsDeleted  int32      `gorm:"column:is_deleted" json:"is_deleted"`
}

func (t *ActivityTime) TableName() string {
	return "activity_time"
}
