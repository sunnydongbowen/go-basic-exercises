package geek_func1

import (
	"fmt"
	"reflect"
	"testing"
)

// 训练营demo

type User struct {
	name string
	age  int
}

func (user User) GetName() string {
	return user.name
}

func (user User) GetAge() int {
	return user.age
}

func TestKindPanic(t *testing.T) {
	typ := reflect.TypeOf(123)
	fmt.Println(typ)
	// typ.NumField()  会panic，无法调用
	typ = reflect.TypeOf(User{
		"bowen",
		10,
	})
	fmt.Println(typ, typ.NumField())

	// 换成指针试试
	typ = reflect.TypeOf(&User{
		"bowen",
		10,
	})
	// Kind()
	if typ.Kind() == reflect.Struct {
		fmt.Println("struct")
	}

	if typ.Kind() == reflect.Ptr {
		fmt.Println("pointer")
	}

	//
	numMethod := typ.NumMethod()
	fmt.Println("方法数量为:", numMethod)

	for i := 0; i < numMethod; i++ {
		method := typ.Method(i)
		fmt.Println(method)
		mt := method.Type
		fmt.Println(mt)
		numIn := mt.NumIn()
		fmt.Println(numIn)
		in := make([]reflect.Type, 0, numIn)
		// 参数打印
		for j := 0; j < numIn; j++ {
			in = append(in, mt.In(j))
			fmt.Println("入参打印:", in)
		}
	}
}
