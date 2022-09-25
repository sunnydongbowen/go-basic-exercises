package main

type C1 interface {
	~int | ~int32 // 这个写法是什么意思？
	M1()
}

type T struct {
}

func (T) M1() {

}

type T1 int

func (t T1) M1() {

}

func foo[P C1](t P) {

}

func main() {
	var t1 T1
	foo(t1)

	//var t T
	//foo(t)
}
