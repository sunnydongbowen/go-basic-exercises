package main

// @program:   go-basic-exercises
// @file:      sum.go
// @author:    bowen
// @time:      2023-09-08 16:30
// @description:

func SumInt(vals []int) int {
	var res int
	for _, v := range vals {
		res += v
	}
	return res
}
