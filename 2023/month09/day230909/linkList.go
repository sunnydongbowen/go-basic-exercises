package main

// @program:   go-basic-exercises
// @file:      linkList.go
// @author:    bowen
// @time:      2023-09-09 10:23
// @description:

type node struct {
}

type LinkedList struct {
	head *node
	tail *node
	// 包外可访问
	Len int
}

func (l LinkedList) Add(idx int, val any) {

}

func (l *LinkedList) AddV1(idx int, val any) {

}
