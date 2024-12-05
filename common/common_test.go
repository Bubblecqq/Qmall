package common

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	intput := "101010101110110"
	c := GetInput(intput)
	out := S1Q(c)
	sum := 0
	for o := range out {
		sum += o
	}
	fmt.Println(sum)
}
