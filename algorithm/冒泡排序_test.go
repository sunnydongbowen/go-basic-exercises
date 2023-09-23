package algorithm

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      冒泡排序_test.go
// @author:    bowen
// @time:      2023-09-23 20:08
// @description:

func BulleSort(a []int) []int {
	//外层循环控制轮数
	for i := 0; i < len(a)-1; i++ {
		//内层循环控制比较的次数，每一轮都会少比较一次
		for j := 0; j < len(a)-1-i; j++ {
			//两数交换
			//var tmp int
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
	return a
}

func TestBubbleSort(t *testing.T) {
	a := []int{31, 11, 33, 10, 11, 8, 100}
	fmt.Println(BulleSort(a))

}
