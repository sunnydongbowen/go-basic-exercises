package batch_insertv2

import (
	"fmt"
	"go-basic-exercises/mysql"
	"sync"
	"testing"
)

var wg sync.WaitGroup

// 产生数据，往chan里送
func producer(ch chan<- []mysql.User) {
	// 这是是往chan中添加userlist
	for i := 0; i < 100; i++ {
		var userlist []mysql.User
		userlist = mysql.Generate()
		ch <- userlist
	}
	close(ch)
}

// 从chan里取数据
func consumer(ch <-chan []mysql.User) {
	//这里面是切片，就可以用sqlx这个方法了
	for userslist := range ch {
		mysql.BatchInsert(userslist)
	}
}

func TestInsertWithGoroutine(*testing.T) {
	// 初始化连接数据库
	if err := mysql.InitDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}

	c := make(chan []mysql.User)
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
