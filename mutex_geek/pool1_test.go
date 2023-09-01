package mutex_geek

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      pool1_test.go
// @author:    bowen
// @time:      2023-04-13 20:40
// @description:

func TestPooldemo(t *testing.T) {
	var strPool = sync.Pool{
		New: func() interface{} {
			return "test str"
		},
	}
	str := strPool.Get()
	fmt.Println(str)
	strPool.Put(str)
}

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			fmt.Println("hhhh,new")
			return []byte{}
		},
	}
	for i := 0; i < 5; i++ {
		val := pool.Get()
		fmt.Println(val)
		//pool.Put(val)
		pool.Put(i)
	}
}
