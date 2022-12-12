package main

import (
	"encoding/json"
	"fmt"
)

// @program:     go-basic-exercises
// @file:        unmarshal.go
// @author:      bowen
// @create:      2022-12-08 22:12
// @description:

type Student struct {
	Name string
	Age  int
}

func main() {
	// 结构体转json
	s1 := &Student{"Amanda", 12}
	bytes_s1, err := json.Marshal(s1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes_s1))

	//
	var s2 string = `{"name":"Judy","age":13}`
	var stu = &Student{}
	err = json.Unmarshal([]byte(s2), stu)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal:Name:%s,Age：%d\n", stu.Name, stu.Age)
}
