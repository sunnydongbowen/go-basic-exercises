package fanxing

import (
	"fmt"
	"testing"
)

// @program:     go-basic-exercises
// @file:        typepara_test.go
// @author:      bowen
// @create:      2022-12-11 11:35
// @description: tony白专栏，类型参数

func maxInt(sl []int) int {
	if len(sl) == 0 {
		panic("slice is empty")
	}
	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func maxString(sl []string) string {
	if len(sl) == 0 {
		panic("slice is empty")
	}
	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func maxFloat(sl []float64) float64 {
	if len(sl) == 0 {
		panic("slice is empty")
	}
	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func maxAny(sl []any) any {
	if len(sl) == 0 {
		panic("slice is empty")
	}
	max := sl[0]
	for _, v := range sl[1:] {
		//这里必须要做类型断言，否则没法比较的！毕竟是强类型语言，不像py，直接写是不可以的!
		// any不能随便比较大小
		//if v>max{
		//
		//}
		switch v.(type) {
		case int:
			if v.(int) > max.(int) {
				max = v
			}
		case string:
			if v.(string) > max.(string) {
				max = v
			}
		case float64:
			if v.(float64) > max.(float64) {
				max = v
			}
		}
	}
	return max

}
func TestMax(t *testing.T) {
	sint := []int{3, 1, 3, 7}
	fmt.Println(maxInt(sint))
	// 按字典排序
	sstr := []string{"111", "22", "44", "66", "77", "10"}
	fmt.Println(maxString(sstr))

	sfloat := []float64{1.01, 2.02, 3.03, 5.05, 7.07, 0.01}
	fmt.Println(maxFloat(sfloat))

	// 这里这样调用是不行的，必须是any类型,强类型语言就是这样的！不能随便调用，不像py随便写！
	//fmt.Println(maxAny(sint))

}

func TestAny(t *testing.T) {
	i := maxAny([]any{1, 2, -4, -6, 7, 0})
	m := i.(int)
	fmt.Println(m)

	fmt.Println(maxAny([]any{"11", "22", "44", "66", "77", "10"}))
	fmt.Println(maxAny([]any{1.01, 2.02, 3.03, 5.05, 7.07, 0.01}))

}
