package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"testing"
)

// sqlx的使用，重要！核心的区分点在这，比原生的要常用，后面我们基本用的这个
var dbx *sqlx.DB

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func initDB() (err error) {
	dsn := "root:815qza@tcp(192.168.72.130:3306)/sql_test?charset=utf8mb4&parseTime=True"
	dbx, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed,err:%v\n", err)
		return
	}
	dbx.SetMaxOpenConns(20)
	dbx.SetMaxIdleConns(10)
	return
}

// 这个放在这干啥的？？？？
//func (u User) Value() (driver.Value, error) {
//	return []interface{}{u.Name, u.Age}, nil
//}

func queryRow() {
	sqlStr := "select id, name, age from user where id=?"
	var u User
	// 自动把数据塞到结构体里，这里挎包了，所以结构体要大写，用的是反射。
	err := dbx.Get(&u, sqlStr, 3)
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return
	}
	fmt.Println(u)
}

func TestQuerRow(t *testing.T) {
	if err := initDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}
	queryRow()
	defer dbx.Close()
}

func queryMultiRowSqlx() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []User // 这种结构类似py里列表嵌套字典

	err := dbx.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	fmt.Println(users)
}

func TestMultiRowSqlx(t *testing.T) {
	if err := initDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}
	queryMultiRowSqlx()
	defer dbx.Close()
}

func insertUser() {
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err := dbx.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "bowen",
			"age":  33,
		},
	)
	if err != nil {
		fmt.Println("insert failed,err: ", err)
		return
	}
	return
}

func TestInsertqlx(t *testing.T) {
	if err := initDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}
	insertUser()
	defer dbx.Close()
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {

}
