package model

import (
	"fmt"
	"testing"
)

func TestConvertTime(t *testing.T) {
	time := "2022-07-14 16:46:10"
	fmt.Println(StringTimeConvertNew(time))
}
