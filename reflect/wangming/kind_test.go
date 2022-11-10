package wangming

// 阅读汪明Go并发编程实战第10章反射学习

import (
	"fmt"
	"reflect"
	"testing"
)

// 是别人并不知道我给他的是什么类型，他拿过来要用kind判断一下，是相应的类型才走相应的处理

func PrintValue(obj interface{}) {
	v := reflect.ValueOf(obj) //返回值是value类型的
	k := v.Kind()             // Valueof调用的
	switch k {
	case reflect.Int:
		fmt.Printf("int value is：%d\n", v.Int()) // Int方法只能Int系列的类型才能调用
	case reflect.Float32:
		fmt.Printf("Float32 value is：%f\n", v.Float())
	case reflect.Float64:
		fmt.Printf("Float64 value is：%f\n", v.Float())
	case reflect.String:
		fmt.Printf("String value is：%s\n", v.String())
	case reflect.Bool:
		fmt.Printf("bool value is:%t\n", v.Bool())
	default:
		fmt.Printf("unkown type:%v", v)
	}
}

func TestKind(t *testing.T) {
	var o interface{}
	// interface value -> reflect oject
	o = 3.14115
	ReflectPrint(o)
	// reflect obj->interface value
	iv := reflect.ValueOf(o).Interface()
	fmt.Printf("%+v %T\n", iv, iv) //

	//interface value->reflect object
	refObject := reflect.ValueOf(iv)
	fmt.Println(refObject.Kind(), refObject.Float()) // float64
	//fmt.Println(refObject.Float()) // value

	PrintValue(o)

	o = "hello GO"
	PrintValue(o)

	o = true
	PrintValue(o)

	o = 7
	PrintValue(o)
}
