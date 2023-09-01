package day230831

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      swtich_test.go
// @author:    bowen
// @time:      2023-08-31 16:24
// @description:

func TestSwtich(t *testing.T) {
	i := 1
	switch i {
	case 2:
		fmt.Println("2")
	case 1:
		fmt.Println("1")
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("sorry")
	}
}
