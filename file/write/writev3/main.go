package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// @program:     go-basic-exercises
// @file:        main.go
// @author:      bowen
// @create:      2022-11-28 13:12
// @description:

func main() {
	content := flag.String("fcontent", "", "-fcontent指定写入文件的内容")
	flag.Parse()

	//创建文件
	f, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	//字符串转slice
	slice := strings.Split(*content, ";")
	for _, line := range slice {
		// 写入行
		_, err := fmt.Fprintln(f, line)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("写入成功")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
