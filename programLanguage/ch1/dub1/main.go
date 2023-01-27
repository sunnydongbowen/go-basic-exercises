package main

import (
	"bufio"
	"fmt"
	"os"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023/1/24 17:12
// @description:

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//这个map因为是无序的，key是唯一的。遇到11还是会加1。
		counts[input.Text()]++
		//if input.Text() == "break" {
		//	break
		//}
	}
	fmt.Println(counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
