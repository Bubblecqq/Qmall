package common

import (
	"math"
	"strconv"
)

func StringToArray(input string) []int {
	var output []int
	for _, v := range input {
		output = append(output, int(v))
	}
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	return output
}

func GetInput(input string) <-chan int {
	out := make(chan int)
	go func() {
		for _, b := range StringToArray(input) {
			out <- b
		}
		close(out)
	}()
	return out
}

func S1Q(in <-chan int) <-chan int {
	out := make(chan int)
	var base, i float64 = 2, 0
	go func() {
		for n := range in {
			out <- (n - 48) * int(math.Pow(base, i))
			i++
		}
		close(out)
	}()
	return out
}

func ToInput(intput string) int {
	c := GetInput(intput)
	out := S1Q(c)
	sum := 0
	for o := range out {
		sum += o
	}
	return sum
}

func ConverToBinary(n int) string {
	res := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}
