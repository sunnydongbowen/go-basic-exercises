package main

import (
	"flag"
	"fmt"
	"go-basic-exercises/file/readproperty/property"
)

// @program:     go-basic-exercises
// @file:        main、.go
// @author:      bowen
// @create:      2022-11-29 12:31
// @description: go并发编程笔记

func main() {
	content := flag.String("fkey", "db.uname", "-fkey指定配置文件的key")
	flag.Parse()
	fmt.Printf("key[%s]->%s", *content, property.Read(*content))
}
