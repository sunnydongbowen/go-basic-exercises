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
// @create:      2022-11-26 23:03
// @description:

func main() {
	fptr := flag.String("fpath", "hello.txt", "-fpath指f定文件路径读取")
	bytelen := flag.Int("flen", 6, "-flen指定读取字节数")
	flag.Parse()
	file, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	// defer 关闭文件，释放资源
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("文件关闭失败 ", err)
		}
	}()

	//文件分片读取到大文件中，可以处理大文件
	r := bufio.NewReader(file)
	buffer := make([]byte, *bytelen)
	fmt.Println("文件内容为:")
	for {
		n, err := r.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
