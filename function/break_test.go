package function

import (
	"fmt"
	"testing"
)

func TestBreak(t *testing.T) {
	for i := 0; i < 3; i++ {
		if i == 1 {
			fmt.Printf("目前i=%d,跳出整个循环\n", i)
			break
		}
		fmt.Printf("循环继续i=%d\n", i)
	}
	fmt.Println("继续执行循环以后的操作")
}

func TestContinue(t *testing.T) {
	for i := 0; i < 3; i++ {
		if i == 1 {
			fmt.Printf("目前i=%d,跳出本次循环\n", i)
			continue
		}
		fmt.Printf("循环继续i=%d\n", i)
	}
	fmt.Println("继续执行循环以后的操作")
}

func TestReturn(t *testing.T) {
	for i := 0; i < 3; i++ {
		if i == 1 {
			fmt.Printf("目前i=%d,结束循环\n", i)
			return
		}
		fmt.Printf("循环继续i=%d\n", i)
	}
	fmt.Println("继续执行循环以后的操作")
}
