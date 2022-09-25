package main

// 类型具体化

func Sort[Elem interface{ Less(y Elem) bool }](list []Elem) {}

type book struct {
}

func (x book) Less(y book) bool {
	return true
}

func main() {
	var bookshelf []book
	Sort[book](bookshelf) // 泛型函数调用

}
