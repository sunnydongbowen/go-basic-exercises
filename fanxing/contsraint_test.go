package fanxing

import "testing"

// @program:     go-basic-exercises
// @file:        contsraint_test.go
// @author:      bowen
// @create:      2022-12-09 20:29
// @description:  tony白专栏，这个是写第二次了！一点影响都没了？？？

type C1 interface {
	~int | ~int32 | ~bool
	M1()
}

type T struct {
}

// 这里面这个receiver只传了一个类型，recevier的名字给省了
func (T) M1() {

}

type T1 int

func (t T1) M1() {

}

type tb bool

func (t tb) M1() {

}

// 只是在声明的时候不一样
func foo[P C1](t P) {

}

func TestConstraint(t *testing.T) {
	var t1 T1
	// 调用的时候其实是一样的
	foo(t1)
	// 报错
	//var t T
	//foo(t)
	// 这个也可以！
	var b tb
	foo(b)
}
