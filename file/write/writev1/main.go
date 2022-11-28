package main

import (
	"flag"
	"fmt"
	"os"
)

// @program:     go-basic-exercises
// @file:        main.go
// @author:      bowen
// @create:      2022-11-28 11:50
// @description: 汪明go并发编程实战 第八章 文件读写

func main() {
	// 从控制台获取需要写入文件的内容
	content := flag.String("fcontent", "", "-fcontent指定写入文件的内容")
	flag.Parse()
	// 创建文件
	// 返回一个文件描述符，并在当前目录创建一个hello.txt文件，如果本身文件已经存在了，会清空掉。
	f, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 返回的是内存地址，所以才需要使用*
	// 往文件中写入字符串*content,这个方法返回写入的字节数，如果有返回就返回错误。
	n, err := f.WriteString(*content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	if n > 0 {
		fmt.Println("写入完成")
	}
	if err = f.Close(); err != nil {
		fmt.Println(err)
		return
	}

}
