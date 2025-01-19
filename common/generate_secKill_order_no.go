package common

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GenerateSecKillOrderNo 生产 订单号   Y2022 06 27 11 00 53 948 97 103564
//
//	年    月 日 时  分 秒 毫秒 ID  随机数
func GenerateSecKillOrderNo(time2 time.Time, userId int64) string {
	var orderNo string
	tempNum := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	orderNo = "SK" + time2.Format("20060102150405.000") + strconv.Itoa(int(userId)) + tempNum
	orderNo = strings.Replace(orderNo, ".", "", -1)
	return orderNo
}
