package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestResetField(t *testing.T) {
	stu := Student{Name: "Jack", Age: 18, id: "Y001"}
	// 传的必须是指针
	v2 := reflect.ValueOf(&stu.Name)
	v2.Elem().SetString("Jerry")
	fmt.Println(stu)

	v2 = reflect.ValueOf(&stu.Age)
	v2.Elem().SetInt(20)

	//包外不能访问私有变量,因为这里在包内，所以是可以改的
	v2 = reflect.ValueOf(&stu.id)
	v2.Elem().SetString("Yoo2")

	fmt.Println(stu)

}
