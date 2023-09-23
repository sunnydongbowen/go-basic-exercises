package main

import "fmt"

// @program:   go-basic-exercises
// @file:      out.go
// @author:    bowen
// @time:      2023-09-10 14:23
// @description: 组合

type Inner struct {
}

type Outer struct {
	Inner
}

type OuterV1 struct {
	Inner
}

func (i Inner) Dosth() {
	fmt.Println("这是inner")
}

func (o OuterV1) Dosth() {
	fmt.Println("这是outerv1")
}

type OuterPtr struct {
	*Inner
}

func UseInner() {
	var o Outer
	o.Dosth()
	o.Inner.Dosth()

	// 这种很少用。
	var op *OuterPtr
	op.Dosth()

	//初始化
	o1 := Outer{
		Inner{},
	}
	op1 := OuterPtr{
		&Inner{},
	}
	o1.Dosth()
	op1.Dosth()
}
