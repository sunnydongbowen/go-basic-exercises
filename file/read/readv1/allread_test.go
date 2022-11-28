package main

import (
	"fmt"
	"os"
	"testing"
)

// @program:     go-basic-exercises
// @file:        allread_test.go
// @author:      bowen
// @create:      2022-11-26 22:03
// @description: 汪明并发编程实战

func TestReadAll(t *testing.T) {
	// 把文件读到内存中，一般为小文件
	content, err := os.ReadFile("E:\\go-basic-exercises\\file\\read\\hello.txt")
	if err != nil {
		fmt.Println("文件读取失败:", err)
		return
	}
	fmt.Println("文件内容为:", string(content))
}
