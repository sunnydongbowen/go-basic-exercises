package day230902

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      pool_test.go
// @author:    bowen
// @time:      2023-09-02 10:19
// @description: 笔记复习，并发编程，大明课程

type User struct {
	ID   int
	Name string
}

func (u *User) Reset() {
	u.ID = 0
	u.Name = ""
}

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return &User{}
		},
	}
	// 取出来
	u1 := pool.Get().(*User)
	u1.ID = 12
	u1.Name = "TM"

	//重置后放进去
	u1.Reset()
	pool.Put(u1)
	u2 := pool.Get().(*User)
	fmt.Println(u2)
}
