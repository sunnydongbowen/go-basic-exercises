package gonote

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// Go 语言学习笔记的内容
type X int
type Y int

// Kind()和Type区别
func TestNote1(t *testing.T) {
	var a, b X = 100, 200
	var c Y = 300
	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)
	// kind()返回的是基础类型(底层数据类型)，TypeOf()返回的是X
	fmt.Printf("ta=%s,ta.Name()=%s, ta.Kind()=%s,\n", ta, ta.Name(), ta.Kind())

	fmt.Println(ta == tb, ta == tc)
	fmt.Println(ta.Kind() == tc.Kind())

	va, vb, vc := reflect.ValueOf(a), reflect.ValueOf(b), reflect.ValueOf(c)
	fmt.Println(va.Type().Name(), vb.Type().Name(), vc.Type().Name())
	fmt.Println(va.Kind(), vb.Kind(), vc.Kind())
}

// 使用反射构造复合数据类型,不理解实际使用场景，你一直反射操作的话，你不知道是 string 或者是 int
func TestNote2(t *testing.T) {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	fmt.Println(a, m) // [10]uint8 map[string]int

	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a)) // *reflect.rtype [10]uint8
}

// 传入对象应该区分基本类型和指针类型，它们并不属于同一个类型
func TestNote3(t *testing.T) {
	x := 100
	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)
	fmt.Println(tx, tp)               // int *int
	fmt.Println(tx.Kind(), tp.Kind()) // int ptr
	fmt.Println(tx == tp.Elem())
}

// ELem返回指针、数组、切片、字典或通道基本类型
func TestNote4(t *testing.T) {
	fmt.Println(reflect.TypeOf(map[string]int{}).Elem()) // int
	fmt.Println(reflect.TypeOf([]int32{}).Elem())        // int32
}

// 只有在获取结构体指针的基类型后，才能遍历它的字
func TestNote5(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	type manager struct {
		user
		title string
	}

	var m manager
	tp := reflect.TypeOf(&m)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)
		//
		if f.Anonymous {
			// 输出匿名字段结构
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(" ", af.Name, af.Type)
			}
		}
	}
}

// 对于匿名字段，可用多级索引（按定义顺序）直接访
func TestNote6(t *testing.T) {
	type user struct {
		name string
		age  int
	}
	type manager struct {
		user
		title string
	}
	var m manager
	tp := reflect.TypeOf(m)

	// 按名称查找
	name, _ := tp.FieldByName("name") // name string
	fmt.Println(name.Name, name.Type)

	// 按多级索引查
	age := tp.FieldByIndex([]int{0, 1}) // age int
	fmt.Println(age.Name, age.Type)
}

type A int
type B struct {
	A
}

// NumMethod returns the number of methods accessible using Method.
// For a non-interface type, it returns the number of exported methods.
// For an interface type, it returns the number of exported and unexported methods.
func (A) Av()  {}
func (*A) Ap() {}
func (B) Bv()  {}
func (*B) Bp() {}
func TestNote7(t *testing.T) {
	var b B
	tp := reflect.TypeOf(&b)
	s := []reflect.Type{tp, tp.Elem()}
	for _, tp := range s {
		fmt.Println(tp, ":")
		//fmt.Println(tp.NumMethod())
		for i := 0; i < tp.NumMethod(); i++ {
			fmt.Println(" ", tp.Method(i))
		}
	}
}

// 反射能探知当前包或外包的非导出结构成
func TestNote8(t *testing.T) {
	var s http.Server
	tp := reflect.TypeOf(s)
	for i := 0; i < tp.NumField(); i++ {
		fmt.Println(tp.Field(i).Name, tp.Field(i).Type)
	}
}

// 可用反射提取struct tag，还能自动分解。其常用于ORM映射，或数据格式验证。
func TestNote9(t *testing.T) {
	type user struct {
		name string `json:"name" type:"varchar(50)"`
		age  int    `json:"age" type:"int"`
	}
	var u user
	tp := reflect.TypeOf(u)
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		fmt.Printf("%s: %s%s\n", f.Name, f.Tag.Get("json"), f.Tag.Get("type"))
	}
}

func (X) String() string {
	return ""
}

// 辅助判断方法Implements、ConvertibleTo、 AssignableTo都是运行期进行动态调用和赋值所必需的。
func TestNote10(t *testing.T) {
	var a X
	tp := reflect.TypeOf(a)
	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(tp.Implements(st))
	it := reflect.TypeOf(0)
	fmt.Println(tp.AssignableTo(st), tp.AssignableTo(it))

}
