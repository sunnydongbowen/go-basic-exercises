package reflect

// 汪明并发编程笔记或者是Go语言学习笔记中的demo？
// 利用反射修改结构体的值

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimaryType(t *testing.T) {
	msg := "hello"
	typ := reflect.TypeOf(msg) // string
	fmt.Println(typ)
	typ = reflect.TypeOf(&msg) // *string
	fmt.Println(typ)
	// 返回该类型的元素类型，
	//如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
	fmt.Println(typ.Elem()) //string

	// reflect.TypeOf(msg).Elem() // panic
	v := reflect.ValueOf(&msg)
	//Elem返回v持有的接口保管的值的Value封装，
	//或者v持有的指针指向的值的Value封装。
	//如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，
	//会返回Value零值
	fmt.Println(v.Elem()) // hello  Elem返回v持有的接口保管的值的Value封装，

	//
	if v.Elem().CanInterface() {
		vs := v.Elem().Interface().(string)
		//vs := v.Elem().Interface().(int)
		fmt.Println(vs)
	}
	// 可以设置才能修改
	if v.Elem().CanSet() {
		v.Elem().SetString("bowen")
	}
	fmt.Println(msg)

	// 非指针，修改的是副本
	v = reflect.ValueOf(msg)
	if v.CanSet() {
		v.SetString("bowen1")
	}
	fmt.Println(msg)
}

func TestElem(t *testing.T) {
	msg := "hello"
	fmt.Println("修改前的内存地址:", &msg)
	valueOf := reflect.ValueOf(&msg)
	valueOf.Elem().SetString("bowen")
	// valueOf.SetString("bowen") 这样写会panic掉
	fmt.Println(msg)
	fmt.Println("reflect修改后的内存地址:", &msg)
	msg = "Go"
	fmt.Println("重新赋值后的内存地址:", &msg)
	fmt.Println(msg)

	//typeOf:=reflect.TypeOf(msg)
	//Print(msg)

}
