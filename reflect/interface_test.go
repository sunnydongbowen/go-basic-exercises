package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type IDuck interface {
	SingGua1()
	SingGua2()
}
type Goose struct {
}

func (g Goose) SingGua2() {
	fmt.Println("Goose sing gua2 gua2 gua2...")
}

func (g Goose) SingGua1() {
	fmt.Println("Goose sing gua gua gua....")
}

// Determine if the interface is implemented
func TestDetermine(t *testing.T) {
	intf := new(IDuck)                      // *IDUCK
	intfType := reflect.TypeOf(intf).Elem() // reflect.IDuck

	goose := Goose{}
	srcType := reflect.TypeOf(&goose) //*reflect.Goose
	fmt.Println(intfType, srcType)
	// 判断是否实现了
	if srcType.Implements(intfType) {
		fmt.Printf("%v 实现了 %v 接口\n", srcType, intfType) // *reflect.Goose 实现了 reflect.IDuck 接口
	} else {
		fmt.Printf("%v 没有实现 %v 接口\n", srcType, intfType)
	}
	fmt.Println()
	// 传非指针的情况
	srcType = reflect.TypeOf(goose)
	if srcType.Implements(intfType) {
		fmt.Printf("%v 实现了 %v 接口\n", srcType, intfType) // *reflect.Goose 实现了 reflect.IDuck 接口
	} else {
		fmt.Printf("%v 没有实现 %v 接口\n", srcType, intfType)
	}
}

func TestNoPointer(t *testing.T) {
	intf := new(IDuck)                      // *IDUCK
	intfType := reflect.TypeOf(intf).Elem() // reflect.IDuck

	goose := Goose{}
	// 指针型的
	srcType := reflect.TypeOf(&goose) //*reflect.Goose
	fmt.Println(intfType, srcType)
	// 判断是否实现了
	if srcType.Implements(intfType) {
		fmt.Printf("%v 实现了 %v 接口\n", srcType, intfType) // *reflect.Goose 实现了 reflect.IDuck 接口
	} else {
		fmt.Printf("%v 没有实现 %v 接口\n", srcType, intfType)
	}
	fmt.Println()
	// 传非指针的情况
	srcType = reflect.TypeOf(goose)
	if srcType.Implements(intfType) {
		fmt.Printf("%v 实现了 %v 接口\n", srcType, intfType) // *reflect.Goose 实现了 reflect.IDuck 接口
	} else {
		fmt.Printf("%v 没有实现 %v 接口\n", srcType, intfType)
	}

}
