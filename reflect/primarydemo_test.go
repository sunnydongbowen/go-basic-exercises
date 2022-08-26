package reflect

// 阅读汪明Go并发编程实战第10章反射学习
import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// 为了代码整洁，这里写了一个函数
func Print(val any) {
	fmt.Println("type: ", reflect.TypeOf(val))
	fmt.Println("value", reflect.ValueOf(val))
	fmt.Println("kind", reflect.ValueOf(val).Kind())
}
func Test(t *testing.T) {
	num := 3.14
	Print(num)

	ti := time.Now()
	Print(ti)

	var v1 interface{}
	v1 = "hello world"
	Print(v1)

	v1 = true
	Print(v1)

	v1 = 2
	Print(v1)

	v1 = make([]string, 0)
	Print(v1) // []string
	//fmt.Println("type kind: ", reflect.TypeOf(v1).Kind()) // slice
	//fmt.Println("value Canset:", reflect.ValueOf(v1).CanSet()) // false

	v1 = nil
	Print(v1)

}
