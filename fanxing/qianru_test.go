package fanxing

import (
	"fmt"
	"sync"
	"testing"
)

// @program:     go-basic-exercises
// @file:        qianru_test.go
// @author:      bowen
// @create:      2022-12-12 10:39
// @description: tony白專刊，泛型类型与类型嵌入

// T类型的切片
type Slice[T any] []T

func (s Slice[T]) String() string {
	if len(s) == 0 {
		return ""
	}
	//把ｓ的第一个赋值给res
	var result = fmt.Sprintf("%v", s[0])
	for _, v := range s[1:] {
		// 拼接
		result = fmt.Sprintf("%v,%v", result, v)
	}
	return result
}

type Lockable[T any] struct {
	t T
	Slice[int]
	sync.Mutex
}

type Foo struct {
	Slice[int]
}

type FooS struct {
	Slice[string]
}

func TestQianru(t *testing.T) {
	n := Lockable[string]{
		t:     "hello",
		Slice: []int{1, 2, 3},
	}
	println(n.String())

	f := Foo{
		Slice: []int{12, 11, 13},
	}
	println(f.String())

	fs := FooS{
		Slice: []string{"bowen", "hello", "hhh"},
	}
	println(fs.String())
}
