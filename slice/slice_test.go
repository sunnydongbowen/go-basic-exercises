package slice

import (
	"fmt"
	"testing"
)

// 引用类型
func changeSlice(s []int) {
	s[0] = 100
}

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3}
	changeSlice(s)
	fmt.Println(s)
}

func changeMap(m map[string]int) {
	m["key"] = 100
}

func TestMap(t *testing.T) {
	m := map[string]int{
		"key": 1,
	}
	changeMap(m)
	fmt.Println(m)
}
