package main

import (
	"bufio"
	"fmt"
	"os"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023/1/25 11:39
// @description:

func main() {
	counts := make(map[string]int)
	// 这里读的是文件？是的，传一个文件名
	files := os.Args[1:]
	// 接收输入
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// 可以接收多个文件
		for _, arg := range files {
			countsfile := make(map[string]int)
			//fmt.Println(len(files))
			// 对每个文件进行操作
			f, err := os.Open(arg)
			// 捕捉错误
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			// 调用countLines方法
			countLines(f, countsfile)
			f.Close()

			for line, n := range countsfile {
				if n > 1 {
					fmt.Printf("%s\t%d\t%s\n", arg, n, line)
				}
			}
		}
	}
	// 这一部分跟dub1是一样的
}

// 这个函数两个参数，一个文件，一个是map，主体还是dub1一样的，只是换成从文件里读了
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	//fmt.Println(f.Name())
	for input.Scan() {
		counts[input.Text()]++
	}
}
