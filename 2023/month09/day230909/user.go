package main

import "fmt"

// @program:   go-basic-exercises
// @file:      user.go
// @author:    bowen
// @time:      2023-09-09 11:30
// @description:

type User struct {
	Name string
	Age  int
}

func NewUser() {
	u := User{}
	//println(u)  //无法打印
	fmt.Printf("%v \n", u)
	fmt.Printf("%+v \n", u)

	//结构体指针
	up := &User{}
	fmt.Printf("up %+v \n", up)

	up2 := new(User)
	fmt.Printf("up2 %+v \n", up2)

	//赋值
	u4 := User{Name: "Tom", Age: 8}
	fmt.Println(u4)

	var up3 *User
	fmt.Println(up3)
	//fmt.Println(up.Name)  //报错

}

func (u User) ChangeName(name string) {
	fmt.Printf("change name中的 u的地址: %p\n", &u)
	u.Name = name
}

func (u *User) ChangeAge(age int) {
	fmt.Printf("change age中的 u的地址: %p\n", u)
	u.Age = age
}

func ChangeName(u User, name string) {

}

func ChangeAge(u *User, age int) {

}

func ChangeUser() {
	u1 := User{Name: "小明", Age: 11}
	fmt.Printf("change User中的 u的地址: %p\n", &u1)
	u1.ChangeName("大明明")
	u1.ChangeAge(13)
	fmt.Println(u1)
}
