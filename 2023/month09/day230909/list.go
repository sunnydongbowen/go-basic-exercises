package main

// @program:   go-basic-exercises
// @file:      list.go
// @author:    bowen
// @time:      2023-09-09 10:13
// @description:

type List interface {
	Add(idx int, val any) error
	Append(val any)
	Delete(idx int) (any, error)
	toSlice() ([]any, error)
}

type linklist struct {
	head *node
	tail *node
}

func (l linklist) Add(idx int, val any) error {
	//TODO implement me
	panic("implement me")
}

func (l linklist) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l linklist) Delete(idx int) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (l linklist) toSlice() ([]any, error) {
	//TODO implement me
	panic("implement me")
}
