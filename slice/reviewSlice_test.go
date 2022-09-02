package slice

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
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
	ss1 := make([]string, 3, 4)
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
	ss2 := make([]string, 0, 2)
	fmt.Println(len(ss2), cap(ss2))
	fmt.Println(ss2)
	ss2 = append(ss2, "雨花台")
	ss2 = append(ss2, "江宁")
	ss2 = append(ss2, "鼓楼")
	ss2 = append(ss2, "谷里")
	ss3 := []string{"北京", "上海", "深圳", "杭州"}
	ss2 = append(ss2, ss3...)
	fmt.Println(len(ss2), cap(ss2))
	fmt.Println(ss2)

}

func TestAppend(t *testing.T) {
	a := make([]int, 0, 2)
	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	for i := 0; i < 50; i++ {
		a = append(a, i)
		if i%5 == 0 {
			//fmt.Printf("第%d次扩容后的地址%p\n", i, &a[0]) // 这里不能写&a，&a是不变的。&a打印的控制结构，肯定无变化的
			// 控制结构有一个指向buffer的指针，有个len，还有个cap,变化的是buffer,用unsafe拿
			//fmt.Println(unsafe.Pointer(&a))
			fmt.Println(sh1.Data)
		}
	}
}
