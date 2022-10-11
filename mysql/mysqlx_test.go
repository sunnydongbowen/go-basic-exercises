package mysql

import (
	"database/sql/driver"
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
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

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
// 感觉这个没有下面第三个方便
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args

	_, err := dbx.Exec(query, args...)
	return err
}

func TestBatchInsertUsers2(t *testing.T) {
	if err := initDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}
	defer dbx.Close()
	var users = []interface{}{
		User{Name: "bowen1002", Age: 30},
		User{Name: "bowen10", Age: 10},
	}
	fmt.Println(BatchInsertUsers2(users))
}
func queryByName() {
	sqlStr := "select id, name, age from user where name=:name" // 结构体里要加tag
	u := User{
		Name: "bowen",
	}
	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err := dbx.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

func TestQueryByName(t *testing.T) {
	if err := initDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}
	queryByName()
	defer dbx.Close()
}

// QueryByIDs 根据给定ID查询
func QueryByIDS(ids []int) (Users []User, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT id, name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = dbx.Rebind(query)
	err = dbx.Select(&Users, query, args...)
	return
}

func TestQueryByIDS(t *testing.T) {
	if err := initDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}
	defer dbx.Close()
	userList, err := QueryByIDS([]int{1, 3, 5, 7})
	for _, user := range userList {
		fmt.Println(user)
	}
	if err != nil {
		fmt.Println("QueryByIDS failed,err: ", err)
		return
	}
	//defer dbx.Close()
}

// 这种插入还挺方便的
// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []User) error {
	_, err := dbx.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}

func TestBatchInsertUsers3(t *testing.T) {
	if err := initDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}
	defer dbx.Close()
	var users2 = []User{
		{Name: "jade", Age: 91},
		{Name: "Jack", Age: 87},
		{Name: "小明", Age: 11},
	}
	BatchInsertUsers3(users2)
}
