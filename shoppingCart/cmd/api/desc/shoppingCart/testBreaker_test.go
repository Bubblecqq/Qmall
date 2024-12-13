package main

//
//import (
//	"errors"
//	"fmt"
//	"github.com/zeromicro/go-zero/core/breaker"
//	"math/rand"
//	"testing"
//	"time"
//)
//
//type mockError struct {
//	status int
//}
//
//func (e mockError) Error() string {
//	return fmt.Sprintf("HTTP STATUS: %d", e.status)
//}
//
//func TestBreaker(t *testing.T) {
//	for i := 0; i < 1000; i++ {
//		if err := breaker.Do("test", func() error {
//			return mockRequest()
//		}); err != nil {
//			println(err.Error())
//		}
//	}
//}
//
////func mockRequest() error {
////	source := rand.NewSource(time.Now().UnixNano())
////	r := rand.New(source)
////	num := r.Intn(100)
////	if num%4 == 0 {
////		return nil
////	} else if num%5 == 0 {
////		return mockError{status: 500}
////	}
////	return errors.New("dummy")
////}
