package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// @program:     go-basic-exercises
// @file:        main.go
// @author:      bowen
// @create:      2022-11-27 11:55
// @description: 分行读取

func main() {
	fptr := flag.String("fpath", "hello.txt", "-path指定文件路径读取")
	flag.Parse()

	file, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	// defer 关闭文件，释放资源
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("文件操作失败:", err)
		}
	}()
	// 分行读取
	fmt.Println("文件内容为:")
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err = sc.Err(); err != nil {
		fmt.Println("文件操作失败:", err)
	}

}
