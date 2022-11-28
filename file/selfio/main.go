package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-basic-exercises/file/selfio/myio"
	"strings"
)

// @program:     go-basic-exercises
// @file:        main.go
// @author:      bowen
// @create:      2022-11-28 22:06
// @description:

func main() {
	content := flag.String("fcontent", "", "-fcontent指定写入文件的内容")
	flag.Parse()

	//创建文件
	fio, err := myio.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//字符串转[]byte类型
	slice := strings.Split(*content, ";")
	for _, line := range slice {
		// 写入行
		_, err := fmt.Fprintln(fio, line)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("写入完成")

	fmt.Println("读取内容")
	fio, err = myio.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	sc := bufio.NewScanner(fio)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

	err = fio.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
