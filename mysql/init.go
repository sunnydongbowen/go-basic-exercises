package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"sync"
)

var DBX *sqlx.DB

var Wg sync.WaitGroup

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 初始化数据库
func InitDB() (err error) {
	dsn := "root:815qza@tcp(192.168.72.130:3306)/sql_test?charset=utf8mb4&parseTime=True"
	DBX, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed,err:%v\n", err)
		return
	}
	DBX.SetMaxOpenConns(20)
	DBX.SetMaxIdleConns(10)
	return
}

// 产生数据
func Generate() (users []User) {
	var usersList = []User{}
	for i := 0; i < 1000; i++ {
		v := strconv.Itoa(i)
		user := User{Name: "bowen" + v, Age: i}
		usersList = append(usersList, user)
	}
	//fmt.Println(usersList)
	return usersList
}

// 执行插入
func BatchInsert(users []User) error {
	_, err := DBX.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}
