package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// 和Type获取类型信息不同，Value专注于对象实例数据读写。
// 在前面章节曾提到过，接口变量会复制对象，
// 且是unaddressable的，所以要想修改目标对象，
// 就必须使用指针。
func TestVnote1(t *testing.T) {
	a := 100
	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()
	// 能不能寻址，能不能被设置
	fmt.Println(va.CanAddr(), va.CanSet()) // false false
	fmt.Println(vp.CanAddr(), vp.CanSet()) // true true
}

type User struct {
	Name string
	code int
}

// 就算传入指针，一样需要通过Elem获取目标对象。 因为被接口存储的指针本身是不能寻址和进行设置操作的。
// 注意，不能对非导出字段直接进行设置操作， 无论是当前包还是外包。
func TestVnote2(t *testing.T) {
	p := new(User)
	v := reflect.ValueOf(p).Elem()
	name := v.FieldByName("Name")
	code := v.FieldByName("code")
	fmt.Printf("name:canaddr=%v, canset=%v\n", name.CanAddr(), name.CanSet()) // name:canaddr=true, canset=true
	fmt.Printf("code:canaddr=%v, code=%v\n", code.CanAddr(), code.CanSet())   // code:canaddr=true, code=false
	if name.CanSet() {
		name.SetString("Tom")
	}
	if code.CanAddr() {
		*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 100
	}
	fmt.Printf("%+v\n", *p)

}

// Value.Pointer和Value.Int等方法类似，将Value. data存储的数据转换为指针，目标必须是指针类型。
//而UnsafeAddr返回任何CanAddr
//Value.data地址（相当于&取地址操作），比如Elem后的Value， 以及字段成员地址。
//以结构体里的指针类型字段为例， Pointer返回该字段所保存的地址， 而UnsafeAddr返回该字段自身的地址（结构对象地址+偏移量）
//可通过Interface方法进行类型推断和转换

func TestVnote3(t *testing.T) {
	type user struct {
		Name string
		Age  int
	}
	u := user{
		"q.yuwen",
		60,
	}
	v := reflect.ValueOf(&u)
	if !v.CanInterface() {
		println("CanInterface:faile.")
		return
	}
	p, ok := v.Interface().(*user)
	if !ok {
		print("Interface:failed.")
		return
	}
	p.Age++
	fmt.Printf("%+v\n", u)
}

// 也可直接使用Value.Int、Bool等方法进行类型转换， 但失败时会引发panic，且不支持ok-idiom。
// 复合类型对象设置示例
func TestVnote4(t *testing.T) {
	c := make(chan int, 4)
	v := reflect.ValueOf(c)
	if v.TrySend(reflect.ValueOf(100)) {
		fmt.Println(v.TryRecv())
	}
}

// 接口有两种nil状态，这一直是个潜在麻烦。 解决方法是用IsNil判断值是否为nil。
func TestVnote5(t *testing.T) {
	var a interface{} = nil
	var b interface{} = (*int)(nil)
	fmt.Println(a == nil)                             // true
	fmt.Println(b == nil, reflect.ValueOf(b).IsNil()) // false true
	// 也可用unsafe转换后直接判断iface.data是否为零值
	iface := (*[2]uintptr)(unsafe.Pointer(&b))
	fmt.Println(iface, iface[1] == 0)
}

// 让人很无奈的是，Value里的某些方法并未实现ok- idom或返回error，所以得自行判断返回的是否为Zero Value。
func TestVnote6(t *testing.T) {
	v := reflect.ValueOf(struct{ name string }{})
	println(v.FieldByName("name").IsValid())
	println(v.FieldByName("xxx").IsValid())

}
