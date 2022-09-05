package structdemo

import (
	"fmt"
	"testing"
)

type per struct {
	name string
	age  int
}

func TestStructPointer(t *testing.T) {
	// 初始化结构体
	p1 := per{"bowen", 28}
	fmt.Println(p1.name, p1.age)

	//初始化结构体指针
	//var p2 *per=&per{"bpowen", 19}
	p2 := &per{"bpowen", 19} // 这一句相当于上一句的简写
	fmt.Println(p2.name, p2.age)

	// new 方式創建
	var p3 *per = new(per)
	p3.name = "bpwenm"
	p3.age = 32
	fmt.Println(p3.name, p3.age)

}
