package repository

import (
	"QMall/common"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	dsn := fmt.Sprintf(common.MysqlConnect, "root", "000000", "192.168.23.233:3306", "mall_marketing", "utf8mb4")
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

	activityRepository := NewActivityRepository(db, nil)
	info, err := activityRepository.GetActivityInfo(2)
	// 定义时间格式，这里假设时间格式为 "15:04:05"
	const timeFormat = "15:04:05"
	const currentTime = "10:23:01"
	currentTime2, err := time.Parse(timeFormat, currentTime)

	fmt.Println("info>", info)
	for _, info := range info {
		startTime, err := time.Parse(timeFormat, info.StartTime)
		if err != nil {
			// 处理解析错误
			fmt.Printf("解析开始时间错误: %v", err)
			continue
		}
		endTime, err := time.Parse(timeFormat, info.EndTime)
		if err != nil {
			// 处理解析错误
			fmt.Printf("解析开始时间错误: %v", err)
			continue
		}
		fmt.Println("hour>", startTime.Hour(), "minute>", startTime.Minute(), "second>", startTime.Second())
		if currentTime2.Hour() > startTime.Hour() && currentTime2.Hour() < endTime.Hour() {
			fmt.Printf("活动正在进行中////")
		}
	}
}
