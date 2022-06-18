package error

import (
	"fmt"
	"testing"
)

func TestCustomError(t *testing.T) {
	err := doSth()
	if err == nil {
		fmt.Println("err == nil")
	}
	fmt.Println("---------------")
	// 为什么前面检测是 nil，传入后判断就不是 nil 呢
	handleError(err)
}
