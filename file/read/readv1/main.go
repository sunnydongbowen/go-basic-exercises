package main

import (
	"flag"
	"fmt"
	"os"
)

// @program:     go-basic-exercises
// @file:        main.go
// @author:      bowen
// @create:      2022-11-26 22:26
// @description: 从文件路径读取

func main() {
	// 根据参数进行提取
	fptr := flag.String("fpath", "hello.txt", "-fpath指定文件路径读取")
	flag.Parse()

	//把文件读取到内存中，一般为小文件
	content, err := os.ReadFile(*fptr)
	if err != nil {
		fmt.Println("文件读取失败: ", err)
		return
	}
	// 要转一下，不转的话读取出的是字节
	fmt.Println("文件内容为:", string(content))
}
