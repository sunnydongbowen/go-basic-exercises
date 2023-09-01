package day230901

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      map_test.go
// @author:    bowen
// @time:      2023-09-01 10:17
// @description: 笔记复习

func TestMap(t *testing.T) {
	var info map[string](int)
	fmt.Println(info == nil, info)

	info = make(map[string]int, 10)
	fmt.Println(info)

	info["小明"] = 10
	info["小张"] = 11
	fmt.Println(info)
	fmt.Println(info["小明"])

	// 判断key是否存在
	value, ok := info["小张"]
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Println(value)
	}

	// 遍历map
	for k, v := range info {
		fmt.Println(k, v)
	}

	delete(info, "小明")
	fmt.Println(info)
}
