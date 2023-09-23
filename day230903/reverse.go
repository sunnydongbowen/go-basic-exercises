package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// @program:   go-basic-exercises
// @file:      reverse.go
// @author:    bowen
// @time:      2023-09-03 15:09
// @description: 笔记复习，algorithm,倒序输出

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	words := strings.Fields(input)

	for i := len(words) - 1; i > 0; i-- {
		fmt.Printf("%s ", words)
	}
}
