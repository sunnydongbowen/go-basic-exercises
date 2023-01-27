package main

import (
	"fmt"
	"os"
	"strings"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023/1/24 12:34
// @description:

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
