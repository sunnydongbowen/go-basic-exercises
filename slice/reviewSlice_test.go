package slice

import (
	"fmt"
	"testing"
)

func TestReviewSlice(t *testing.T) {
	// 整型切片
	var a []int
	a = []int{10, 11}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 字符串切片
	ss := [...]string{"雨花台", "栖霞", "江宁"}
	for i := 0; i < len(ss); i++ {
		println(ss[i])
	}
	fmt.Println(ss)

	// 扩容及make初始化，注意前面的默认值
	ss1 := make([]string, 3, 8)
	fmt.Println(len(ss1), cap(ss1))
	fmt.Println(ss1)
	ss1 = append(ss1, "雨花台")
	ss1 = append(ss1, "江宁")
	ss1 = append(ss1, "鼓楼")
	ss1 = append(ss1, "谷里")
	fmt.Println(len(ss1), cap(ss1))
	fmt.Println(ss1)

	//整型的，前面依然有三个默认值
	ints := make([]int, 3, 10)
	fmt.Println(ints)

	//bool型切片。注意前面默认值
	bools := make([]bool, 3, 10)
	fmt.Println(bools)
	// chan切片，注意前面默认值
	chans := make([]chan int, 3, 10)
	fmt.Println(chans)

	// 这就是为什么一般使用make初始化切片的时候，要为第二个参数0的原因
	ss2 := make([]string, 0, 8)
	fmt.Println(len(ss2), cap(ss2))
	fmt.Println(ss2)
	ss2 = append(ss2, "雨花台")
	ss2 = append(ss2, "江宁")
	ss2 = append(ss2, "鼓楼")
	ss2 = append(ss2, "谷里")
	fmt.Println(len(ss2), cap(ss2))
	fmt.Println(ss2)

}
