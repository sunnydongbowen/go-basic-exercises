package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// @program:     go-basic-exercises
// @file:        marshal_test.go
// @author:      bowen
// @create:      2022-12-08 17:29
// @description:Json Marshal：将数据编码成json字符串

type Class struct {
	Name  string
	Grade int
}

type Stu struct {
	Name  string `json:"name"`
	Age   int
	sex   string
	HIgh  bool
	Class *Class `json:"class"`
}

func TestMarshal(t *testing.T) {
	// 实例化一个数据结构，用于生成json字符串
	stu := Stu{
		Name: "小米",
		Age:  11,
		sex:  "男",
		HIgh: true,
	}
	// 指针变量
	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	// 这里是可以后面再给他赋值的！
	stu.Class = cla

	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))
}
