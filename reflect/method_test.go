package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type PC struct {
	Name string
}

func (p *PC) GetName() string {
	return p.Name
}

func (p PC) Sum(a, b int) int {
	return a + b
}

func TestMethod(t *testing.T) {
	pc := PC{
		"Yoga",
	}
	//指针类型能够调用值receiver和指针receiver的method
	v := reflect.ValueOf(&pc)
	vm := v.MethodByName("GetName")
	res := vm.Call(nil)
	fmt.Println("GetName=:", res[0].String())
}
