package day230831

// @program:   go-basic-exercises
// @file:      structFactory.go
// @author:    bowen
// @time:      2023-08-31 22:05
// @description: 笔记复习

type file struct {
	fd   int
	name string
}

type matrix struct {
	name string
}

func Newfile(fd int, name string) *file {
	return &file{
		fd:   fd,
		name: name,
	}
}

func NewMatrix(name string) *matrix {
	m := new(matrix)
	*m = matrix{name: name}
	return m
}
