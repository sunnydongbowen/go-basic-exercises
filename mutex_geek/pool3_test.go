package mutex_geek

import (
	"errors"
	"sync"
	"sync/atomic"
	"unsafe"
)

// @program:   go-basic-exercises
// @file:      pool3_test.go
// @author:    bowen
// @time:      2023-04-14 10:46
// @description:

type Mypool struct {
	pool   sync.Pool
	cnt    int32
	maxCnt int32
}

func (p *Mypool) Get() any {
	return p.Get()
}

func (p *Mypool) Put(val any) any {
	// 大对象不放回去
	if unsafe.Sizeof(val) > 1024 {
		return errors.New("超过1024，大对象不可放")
	}
	// 超过数量限制
	cnt := atomic.AddInt32(&p.cnt, 1)
	if cnt >= p.maxCnt {
		atomic.AddInt32(&p.cnt, -1)
		return errors.New("操作对象过多，退出")
	}
	p.pool.Put(val)
	return nil
}
