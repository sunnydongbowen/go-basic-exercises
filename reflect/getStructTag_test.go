package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetStructTag(t *testing.T) {
	stu := Student{Name: "Jack", Age: 19, id: "Y001"}
	typ := reflect.TypeOf(&stu).Elem()
	fmt.Println(typ) // reflect.Student,返回的是指针指向的数据类型,如果没有加Elem()，则返回*reflect.Student
	// 获取tag
	m := make(map[string]string)
	for i := 0; i < typ.NumField(); i++ {
		m[typ.Field(i).Name] = typ.Field(i).Tag.Get("json")
	}
	fmt.Println(m)

	if f, ok := typ.FieldByName("Name"); ok {
		fmt.Println("Name字段的table标识值:", f.Tag.Get("table"))
	}
	if f, ok := typ.FieldByName("id"); ok {
		fmt.Println("id字段的iskey值:", f.Tag.Get("iskey"))
	}
	fmt.Println(stu)

}
