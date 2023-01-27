package main

import (
	"fmt"
	"os"
	"strings"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023/1/25 16:20
// @description:
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
