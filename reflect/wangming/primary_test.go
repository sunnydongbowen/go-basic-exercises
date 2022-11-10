package wangming

// 阅读汪明Go并发编程实战第10章反射学习
import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// 为了代码整洁，这里写了一个函数
func ReflectPrint(val any) {
	fmt.Println("type:", reflect.TypeOf(val), " value:", reflect.ValueOf(val), " kind:", reflect.ValueOf(val).Kind())
	//fmt.Println("value", reflect.ValueOf(val))
	//fmt.Println("kind", reflect.ValueOf(val).Kind())
}
func TestPrimary(t *testing.T) {
	num := 3.14
	ReflectPrint(num)

	ti := time.Now()
	ReflectPrint(ti)

	var v1 interface{}
	v1 = "hello world"
	ReflectPrint(v1)

	v1 = true
	ReflectPrint(v1)

	v1 = 2
	ReflectPrint(v1)

	v1 = make([]string, 0)
	ReflectPrint(v1) // []string
	//fmt.Println("type kind: ", reflect.TypeOf(v1).Kind()) // slice
	//fmt.Println("value Canset:", reflect.ValueOf(v1).CanSet()) // false

	v1 = nil
	ReflectPrint(v1)

}
