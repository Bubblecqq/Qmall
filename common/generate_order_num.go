package common

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GenerateProductsNum 生产 订单号   Y2022 06 27 11 00 53 948 97 103564
//
//	年    月 日 时  分 秒 毫秒 ID  随机数
func GenerateOrderNum(time2 time.Time, productId int64) string {
	var orderNum string
	tempNum := strconv.Itoa(rand.Intn(999999-100000+1) + 100000)
	orderNum = "T" + time2.Format("20060102150405.000") + strconv.Itoa(int(productId)) + tempNum
	orderNum = strings.Replace(orderNum, ".", "", -1)
	return orderNum
}
