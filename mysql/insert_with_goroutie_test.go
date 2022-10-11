package mysql

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

// 产生数据，往chan里送
func producer(ch chan<- []User) {
	// 这是是往chan中添加userlist
	for i := 0; i < 10; i++ {
		var userlist []User
		userlist = generate()
		ch <- userlist
	}
	close(ch)
}

// 从chan里取数据
func consumer(ch <-chan []User) {
	//这里面是切片，就可以用sqlx这个方法了
	for userslist := range ch {
		BatchInsert(userslist)
	}
}

func TestInsertWithGoroutine(*testing.T) {
	// 初始化连接数据库
	if err := initDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}

	c := make(chan []User)
	wg.Add(2)

	go func() {
		producer(c)
		wg.Done()
	}()

	go func() {
		consumer(c)
		wg.Done()
	}()

	wg.Wait()
}
