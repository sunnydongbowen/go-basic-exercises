package geek_homeworkv1

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// 没有字段的结构体
func TestDemo1(t *testing.T) {
	empty := Empty{}
	typ := reflect.TypeOf(empty)
	if typ.NumField() == 0 {
		fmt.Println("为空")
	}
	// 这里用val判断也可以的
}

// 获取结构体字段
func TestDemo2(t *testing.T) {
	entity := BaseEntity{}
	val := reflect.ValueOf(entity)
	typ := val.Type()
	// 获取字段名称   这里我发现不管val还是typ,其实都能获取到参数
	num := typ.NumField()
	for i := 0; i < num; i++ {
		name := typ.Field(i).Name
		fmt.Println(name) // 能取出这两个字段
		//value := val.Field(i).Interface()
		//fmt.Println(value)
	}
	// 获取结构体的value
	for i := 0; i < num; i++ {
		value := val.Field(i).Interface()
		fmt.Printf("%d,%T\n", value, value)
	}
}

// 测试strings.builder
func TestString(t *testing.T) {
	bd := strings.Builder{}
	bd.WriteString("bowen")
	bd.WriteString("努力")
	fmt.Println(bd.String())
}

// 指针类型的如何拿到fileName和value
func TestPointer(t *testing.T) {
	entity := &BaseEntity{
		CreateTime: 123,
		UpdateTime: ptrInt64(234),
	}

	val := reflect.ValueOf(entity) // 这里传值要注意，不然会被报错！
	// 要转一下
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := reflect.TypeOf(entity)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	//typ := val.Type()

	numField := typ.NumField()

	//a := []interface{}{}
	for i := 0; i < numField; i++ {
		f := typ.Field(i)
		fmt.Println(f.Name)

	}

	//
	for i := 0; i < numField; i++ {
		fmt.Println(val.Field(i).Interface())
	}
	//fmt.Println(a)
}

func TestMultipointer(t *testing.T) {
	entity := func() interface{} {
		entity := &BaseEntity{CreateTime: 123, UpdateTime: ptrInt64(234)}
		return &entity
	}()

	val := reflect.ValueOf(entity)

	//typ := reflect.TypeOf(val)
	fmt.Println(val.Kind())
	fmt.Println(val.Elem().Kind())
	fmt.Println(val.Elem().Kind())

}

func ReadNestStruct(entity any) {
	val := reflect.ValueOf(entity).Elem()
	typ := reflect.TypeOf(val)
	for i := 0; i < val.NumField(); i++ {
		if val.Type().Field(i).Type.Kind() == reflect.Struct {
			ReadNestStruct()

		}
	}

}

func TestComposition(t *testing.T) {
	entity := User{
		BaseEntity: BaseEntity{
			CreateTime: 123,
			UpdateTime: ptrInt64(456),
		},
		Id:       789,
		NickName: sql.NullString{String: "Tom", Valid: true},
	}
	val := reflect.ValueOf(entity).Elem()
	typ := reflect.TypeOf(val)

	for i := 0; i < val.NumField(); i++ {
		if val.Type().Field(i).Type.Kind() == reflect.Struct {

		}
	}

}
