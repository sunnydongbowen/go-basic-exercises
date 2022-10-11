package mysql

import (
	"fmt"
	"strconv"
	"testing"
)

// slice
func BatchInsert(users []User) error {
	_, err := dbx.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}

// 产生数据
func generate() (users []User) {
	var usersList = []User{}
	for i := 0; i < 100; i++ {
		v := strconv.Itoa(i)
		user := User{Name: "bowen" + v, Age: i}
		usersList = append(usersList, user)
	}
	//fmt.Println(usersList)
	return usersList
}

// 批量插入
func TestBatchInsert(t *testing.T) {
	if err := initDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}
	users := generate()
	BatchInsert(users)
	defer dbx.Close()
}

// 测试产生数据
func TestGen(t *testing.T) {
	generate()
}
