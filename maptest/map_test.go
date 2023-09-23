package maptest

import (
	"fmt"
	"testing"
)

// geek time 学习笔记

func TestMap(t *testing.T) {
	//var m map[string]int // m=nil  panic
	m := map[string]int{}
	m["key"] = 1
	fmt.Println(m)
}

// 初始化切片
func TestDefi(t *testing.T) {
	var m map[string]int
	m = make(map[string]int, 10)
	m["age"] = 123
	m["phone"] = 123
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func TestMap2(t *testing.T) {

	m1 := map[int][]string{
		1: []string{"val1_1", "val1_2"},
		3: []string{"val3_1", "val3_2", "val3_3"},
		7: []string{"val7_1"},
	}
	fmt.Println(m1)

	m2 := map[int][]string{
		1: {"val1_1", "val1_2"},
		3: {"val3_1", "val3_2", "val3_3"},
		7: {"val7_1"},
	}
	fmt.Println(m2)

	type Position struct {
		x float64
		y float64
	}

	//m2 := map[Position]string{
	//	Position{29.935523, 52.568915}: "school",
	//	Position{25.352594, 113.304361}: "shopping-mall",
	//	Position{73.224455, 111.804306}: "hospital",
	//}

	m3 := map[Position]string{
		{29.935523, 52.568915}:  "school",
		{25.352594, 113.304361}: "shopping-mall",
		{73.224455, 111.804306}: "hospital",
	}
	fmt.Println(m3)
}

func TestInsert(t *testing.T) {
	m := make(map[int]string)
	m[1] = "value1"
	m[2] = "value2"
	m[3] = "value3"
	fmt.Println(m)
}

func TestChange(t *testing.T) {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	fmt.Println(m)
	fmt.Println(len(m))
	m["key1"] = 11
	m["key2"] = 12
	fmt.Println(m)
	m["key3"] = 3
	fmt.Println(len(m))
}

func TestFind(t *testing.T) {
	m := make(map[string]int)
	//v := m["key1"]
	v, ok := m["key1"]
	// 可以自己试着写这样一个函数
	if !ok {
		fmt.Println("key不存在")
		return
	}
	fmt.Println(m)
	fmt.Println(v)
}

func TestDelete(t *testing.T) {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(m)
	delete(m, "key2")
	fmt.Println(m)

}

func TestIteration(t *testing.T) {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}
	fmt.Printf("{")
	for k, v := range m {
		fmt.Printf("[%d, %d] ", k, v)
	}
	fmt.Printf("}\n")
}

func IteratorMap(m map[int]int) {
	fmt.Printf("{")
	for k, v := range m {
		fmt.Printf("[%d, %d] ", k, v)
	}
	fmt.Printf("}\n")
}

func TestIterMap(t *testing.T) {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}
	for i := 0; i < 3; i++ {
		IteratorMap(m)
	}
}

func foo(m map[string]int) {
	m["key1"] = 11
	m["key2"] = 12
}

func TestFoo(t *testing.T) {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(m)
	foo(m)
	fmt.Println(m)
}
