package algorithm

import (
	"fmt"
	"math"
	"testing"
)

// @program:   go-basic-exercises
// @file:      查找素数_test.go
// @author:    bowen
// @time:      2023-09-23 20:38
// @description: 3

func FindPrime(start, end int) {
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

// 核心逻辑在这
func isPrime(a int) bool {
	if a <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}

	return true
}

func TestIsPrime(t *testing.T) {
	FindPrime(1, 100)
}
