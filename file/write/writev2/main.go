package main

import (
	"flag"
	"fmt"
	"os"
)

// @program:     go-basic-exercises
// @file:        main.go
// @author:      bowen
// @create:      2022-11-28 12:04
// @description: 分字节写入,这个程序有点难，要理解一下了！涉及到字节的，都比较难

func main() {
	content := flag.String("fcontent", "", "-fcontent指定写入文件的内容")
	bytelen := flag.Int("flen", 3, "-flen指定读取的字节数")
	flag.Parse()

	// 创建文件
	f, err := os.Create("hello.txt")
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println("文件操作失败", err)
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 字符串转[]byte类型
	b := ([]byte(*content))
	fmt.Println(b)
	// for循环分批次写入字节
	p1 := *bytelen
	fmt.Println(p1) // 3
	i := 0
	// 调用write方法分字节写入时，由于内容长度不一定就是每批字节的正数倍，因此最后一个字节切片需要进行特殊处理，
	//否则写入的内容可能有额外字节，下面这段逻辑是比较难的。要理清楚就好了。不会的用笔画一画。分析一下过程。
	for {
		if i*p1 > len(b) {
			break
		}
		if (i+1)*p1 > len(b) {
			n, err := f.Write(b[i*p1 : len(b)])
			if err != nil {
				fmt.Println(err)
				break
			}
			if n > 0 {
				fmt.Println("写入字节h：", n)
			}
		} else {
			n, err := f.Write(b[i*p1 : (i+1)*p1])
			if err != nil {
				fmt.Println(err)
			}
			if n > 0 {
				fmt.Println("写入字节：", n)
			}

		}
		i++
	}
	fmt.Println("写入完毕")
}
