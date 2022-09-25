package pointer

import (
	"fmt"
	"go-basic-exercises/pointer/entity"
	"testing"
	"unsafe"
)

func TestNew(t *testing.T) {
	// 注意&，*，&的是取指针地址，*是取值
	var d = new(int)       // 返回指针 ,开辟内存，返回*int
	fmt.Println(d, &d, *d) // 0xc00009e278    0xc00000a038  0
	//fmt.Printf()
	//fmt.Printf("%d", d)
	*d = 9
	fmt.Println(*d)

	var a int64 = 8
	var ptr *int64 = &a // 声明指针变量
	fmt.Printf("%v\n", ptr)
	fmt.Println(*ptr) // 8

	// 通过指针修改值
	*ptr = 9
	fmt.Printf("%v\n", a)

	var c float32 = 3.14
	// ptr=&c 会报错，goland直接给提示了，这一点Goland做的还是很强大的
	ptr2 := &c
	fmt.Printf("%v, %v\n", ptr2, c)

	c = 6.28
	fmt.Println(*ptr2)

}

func TestForrange(t *testing.T) {
	s := []int{3, 4, 5, 6, 7}
	for _, v := range s {
		v = 10 // 这里改的是副本，要记得书上的那个图P144
		fmt.Println(v, &v)
	}
	fmt.Println(s)
}

// 利用unsafe强转
func TestUnsafe(t *testing.T) {
	var a int32 = 8
	var f int64 = 64
	ptr := &a
	fmt.Printf("%T %v %v\n", ptr, ptr, *ptr)

	ptr = (*int32)(unsafe.Pointer(&f)) //強轉了,本來是*int64类型的
	fmt.Printf("%T %v %v\n", ptr, ptr, *ptr)

	*ptr = 10
	fmt.Println(a, f)

}

// 利用unsafe修改导出变量
func TestStruct(t *testing.T) {
	stu := new(entity.Student) //  *student
	// stu.id
	fmt.Printf("%+v\n", stu)

	// 第一个namestring类型
	// 强制类型转换，将*Student类型指针转换成通用的指针类型，由于第一个字段类型是字符串是string，因此可以用（*string）进行指针转换
	// 结构体的开始内存地址就是结构体中第一个字段的内存地址，此时变量p的类型是*string切指向结构体中name字段的内存区域
	p := (*string)(unsafe.Pointer(stu))

	*p = "jack" // 突破第一个私有字段
	fmt.Printf("%+v\n", stu)

	// 第二个字段需要指针运算，第一个是字符串长度为16->uintptr(16)
	// 思考这里为什么要这样写？
	/*
			第二个字段复杂，需要进行指针运算，在Go中，借助unsafe包也可以对指针进行运算，指针运算需要unitptr类型
			unsafe.Pointer(stu)首先获取通用结构体的内存地址，然后考虑第一个指针是字符串类型，占用长度为16字节因此偏移16字节即可，即uintptr(unsafe.Pointer(stu)) + uintptr(16)
		    然后用unsafe.Pointer(stu)将其转成通用型指针，并赋值给指针变量ptr_id,此时的ptr_id和结构体stu中的id私有字段中内存中指向同一块内存地址
	*/
	// ptr_id :=(*int)(unsafe.Pointer(stu)) 这样是错误的
	ptr_id := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(stu)) + uintptr(16)))
	*ptr_id = 1 // 突破第二个私有字段
	fmt.Printf("%+v\n", stu)

}

func TestWildPointer(t *testing.T) {
	// 指针未初始化，野指针
	var ptr *int
	// 打印指针ptr指向的地址
	fmt.Printf("%p\n", ptr) // 0x0
	// *ptr = 3

	fmt.Println(ptr) // <nil>
	var a int = 7
	// 可以重新指向
	ptr = &a
	*ptr = 3
	fmt.Println(a) //3

}

func TestGeek(t *testing.T) {
	var a int = 13
	var p *int = &a
	fmt.Println(p)
}

type foo struct {
	id   string
	age  int8
	addr string
}

func TestSizeOf(t *testing.T) {
	var p1 *int
	var p2 *bool
	var p3 *byte
	var p4 *[20]int
	var p5 *foo
	var p6 unsafe.Pointer

	println(unsafe.Sizeof(p1)) // 8
	println(unsafe.Sizeof(p2)) // 8
	println(unsafe.Sizeof(p3)) // 8
	println(unsafe.Sizeof(p4)) // 8
	println(unsafe.Sizeof(p5)) // 8
	println(unsafe.Sizeof(p6)) // 8

}

// 指针读取
func TestGeek1(t *testing.T) {
	var a int = 17
	var p *int = &a
	println(*p)
	(*p) += 3
	println(a)
	fmt.Println(p) // 和下面打印的是一样的
	fmt.Printf("%p\n", p)
}

func TestGeek2(t *testing.T) {
	var a int = 5
	var b int = 6
	var p *int = &a // 指向变量a所在内存单元
	println(*p)     //输出变量a的值
	p = &b          //指向变量b所在内存单元
	println(*p)     //输出变量b的值
}

func TestGeek3(t *testing.T) {
	var a int = 5
	var p1 *int = &a // p1指向变量a所在内存单元
	var p2 *int = &a // p2指向变量b所在内存单元
	*p1 += 5         // 通过p1修改变量a的值
	println(*p2)     // 10 对变量a的修改可以通过另外一个指针变量p2的解引用反映出来
}

// 二级指针
func TestGeek4(t *testing.T) {
	var a int = 5
	var p1 *int = &a
	println(*p1)

	var b int = 55
	var p2 *int = &b
	println(*p2)

	var pp **int = &p1
	println(**pp)
	pp = &p2
	println(**pp)

}

func foo1(pp **int) {
	var b = 35
	var p1 *int = &b
	(*pp) = p1
}

func TestGeek5(t *testing.T) {
	var a int = 5
	var p *int = &a
	println(*p)

	foo1(&p)
	println(*p)
}
