package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init() 方法，很重要！
	"testing"
)

var db *sql.DB

func TestConnectMysql(t *testing.T) {
	dsn := "root:815qza@tcp(192.168.72.130:3306)/mysql"
	// Open它不会真正去连，只是校验一下格式
	db, err := sql.Open("mysql", dsn) // 如果db是nil，最后调用nil.close() 会给出panic 空指针
	// defer db.Close() 放在这会panic，

	if err != nil {
		fmt.Println("连接mysql失败")
		return //
	}
	// 程序退出时关闭数据库连接，保证数据库不为空，因为如果有错误，走到上面if里直接就return了，根本不会走close语句。
	// 代码是压栈的过程，放在这，它这都没压进去，就不会执行了
	defer db.Close()
	// 拿到db对象之后进行增删改查

}

// db定义成sql.DB指针，err是返回值
func initMysql() (err error) {
	dsn := "root:815qza@tcp(192.168.72.130:3306)/sql_test"
	// Open它不会真正去连，只是校验一下格式
	db, err = sql.Open("mysql", dsn) // 如果db是nil，最后调用nil.close() 会给出panic 空指针
	// defer db.Close() 放在这会panic，

	if err != nil {
		fmt.Println("连接mysql失败")
		return err //
	}
	// 程序退出时关闭数据库连接，保证数据库不为空，因为如果有错误，走到上面if里直接就return了，根本不会走close语句。
	// 代码是压栈的过程，放在这，它这都没压进去，就不会执行了
	//defer db.Close()
	// 拿到db对象之后进行增删改查
	// 去真正连接数据库
	if err := db.Ping(); err != nil {
		return err
	}
	// 配置
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(10)
	return
}

type user struct {
	id   int
	age  int
	name string
}

// 查询单条数据
func QueryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 查询多条数据
func queryMultiRow() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}

}

func insert() {
	sqlStr := "insert into user(name,age) values(?,?)"
	ret, err := db.Exec(sqlStr, "jade", 38)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get insert ID failed,err:%v\v", err)
		return
	}
	fmt.Printf("insert sucess,the id is %d\n", theID)
}

func update() {
	sqlStr := "update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, 50, 3)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	// 操作影响的行数
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func deleteRow() {
	sqlStr := "delete from user where id=?"
	ret, err := db.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("delete failed,err：%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func sqlInject(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed.err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}

// 
func ()  {

}

func TestConnect(t *testing.T) {
	if err := initMysql(); err != nil {
		fmt.Printf("init mysqld failed,err：%v\n", err)
		return
	}
	fmt.Println("数据库链接成功")
	//
	//QueryRowDemo()
	//queryMultiRow()
	//insert()
	//update()
	// deleteRow()
	sqlInject("xxx' or 1=1#")
	sqlInject("xxx' union select * from user #")
	sqlInject("xxx' and (select count(*) from user) <10 #")
	defer db.Close()
}
