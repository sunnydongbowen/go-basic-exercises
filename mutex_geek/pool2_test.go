package mutex_geek

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      pool2_test.go
// @author:    bowen
// @time:      2023-04-14 10:12
// @description:

type User struct {
	ID   int
	Name string
}

func (u *User) Reset() {
	u.ID = 0
	u.Name = ""
}

func TestPool2(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return &User{}
		},
	}
	u1 := pool.Get().(*User)
	u1.ID = 12
	u1.Name = "TOM"
	u1.Reset()
	pool.Put(u1)
	u2 := pool.Get().(*User)
	fmt.Println(u2)
}
