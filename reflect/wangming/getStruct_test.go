package wangming

import (
	"fmt"
	"reflect"
	"testing"
)

type Student struct {
	id   string `json:"id" iskey:"1"`
	Name string `json:"cname" table:"t_student"`
	Age  int    `json:"age"`
}

func (s Student) GetName() {
	fmt.Println(s.Name)
}

func (s *Student) GetAge() {
	fmt.Println(s.Age)
}

func (s Student) getID() {
	fmt.Println(s.id)
}

func PrintFileMethod(o interface{}) {
	valueOf := reflect.ValueOf(o)
	typeOf := valueOf.Type()
	fmt.Println(typeOf)

	if typeOf.Kind() == reflect.Struct {
		for i := 0; i < typeOf.NumField(); i++ { // 返回v持有的结构体类型值的字段数
			// 获取结构体字段
			field := typeOf.Field(i) //返回结构体的第i个字段（的Value封装）
			fmt.Println("结构体字段:", field)
			if valueOf.Field(i).CanInterface() {
				//私有属性报错
				value := valueOf.Field(i).Interface()
				fmt.Printf("%s:%v=%v\n", field.Name, field.Type, value)
			} else {
				fmt.Printf("%s:%v=%v\n", field.Name, field.Type, "私有字段")
			}
		}
	}
	// 获取结构体方方法,不能获取私有方法
	for i := 0; i < typeOf.NumMethod(); i++ { // 结构体的方法数目
		m := typeOf.Method(i)
		fmt.Printf("%s：%v\n", m.Name, m.Type)
	}
}
func TestGetStruct(t *testing.T) {
	stu := Student{
		Name: "小明",
		Age:  12,
		id:   "Y001",
	}
	PrintFileMethod(stu)

	a := 2
	PrintFileMethod(a)

	PrintFileMethod(&stu) // 传了指针型接recevier，没有通过if判断，所以只能拿到方法，事实上即使是指针型recevier，上面接取字段的值会报错的
}
