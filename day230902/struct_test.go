package day230902

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      struct_test.go
// @author:    bowen
// @time:      2023-09-02 14:41
// @description: 笔记复习 struct qimi

type person struct {
	name string
	age  uint8
}

func TestPerson(t *testing.T) {
	var p1 person
	p1.name = "xxx"
	p1.age = 10
	fmt.Println(p1)

	//省略
	p2 := person{"小王", 10}
	fmt.Println(p2)

}
