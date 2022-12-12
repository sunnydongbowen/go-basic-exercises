package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// @program:     go-basic-exercises
// @file:        marshal2_test.go
// @author:      bowen
// @create:      2022-12-08 17:48
// @description:
type ClassA struct {
	Name  string
	Grade int
}

type StuAny struct {
	Name  any `json:"name"`
	Age   any
	sex   any
	HIgh  any
	Class any `json:"class"`
}

func TestMarsha(t *testing.T) {
	cla := ClassA{
		Name:  "A班",
		Grade: 2,
	}
	stu := StuAny{
		Name:  "小米",
		Age:   10,
		sex:   "男",
		HIgh:  true,
		Class: cla,
	}
	StuJson, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("结构体转json失败")
	}
	fmt.Println(string(StuJson))
}

type StuRead struct {
	Name string
	Age  int
}

func TestMarS(t *testing.T) {
	//方式1：只声明，不分配内存
	var stus1 []StuRead

	//方式2：分配初始值为0的内存
	stus2 := make([]StuRead, 0)

	stu1 := StuRead{
		Name: "xiaoming",
		Age:  11,
	}
	stu2 := StuRead{
		Name: "xiaoming",
		Age:  111,
	}
	stus1 = append(stus1, stu1, stu2)
	stus2 = append(stus2, stu1, stu2)
	fmt.Println(stus1)
	fmt.Println(stus2)

	json1, _ := json.Marshal(stus1)
	json2, _ := json.Marshal(stus2)

	//var stusr []StuRead

	fmt.Println(string(json1))
	fmt.Println(string(json2))

}
