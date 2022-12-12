package sqlgeek

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // 引入driver，不然不知道的，匿名引入
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// @program:     go-basic-exercises
// @file:        crud_test.go
// @author:      bowen
// @create:      2022-12-06 11:31
// @description: 大明sql编程

type sqlTestSuite struct {
	suite.Suite // 测试套件，这个需要去看一下怎么用
	// 配置字段
	driver string
	dsn    string
	// 初始化字段
	db *sql.DB
}

// 每次测试跑完都会执行这个
func (s *sqlTestSuite) TearDownTest() {
	// 删除表
	_, err := s.db.Exec("DELETE FROM test_model;")
	if err != nil {
		s.T().Fatal(err) // 这里也需要去看一下T()是什么？
	}
}

// 这其实就是一个建表的过程，
// 但是要理解Suite是什么意思？这个以前看过，
// 测试套件，所有测试跑之前先执行这个
func (s *sqlTestSuite) SetupSuite() {
	db, err := sql.Open(s.driver, s.dsn)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = s.db.ExecContext(ctx, `
CREATE TABLE IF NOT EXISTS test_model(
    id INTEGER PRIMARY KEY,
    first_name TEXT NOT NULL,
    age INTEGER,
    last_name TEXT NOT NULL
)
`)
	if err != nil {
		s.T().Fatal(err)
	}
}

type TestModel struct {
	id        int64 `eorm:"auto_increment,primary_key"`
	FirstName string
	Age       int8
	LastName  *sql.NullString
}

func (s *sqlTestSuite) TestCRUD() {
	t := s.T()
	db, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=memory")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//或者 Exec(xxx)
	res, err := db.ExecContext(ctx, "INSERT INTO `test_model`(`id`, `first_name`, `age`, `last_name`) VALUES (1, 'Tom', 18, 'Jerry')")
	if err != nil {
		t.Fatal(err)
	}
	affected, err := res.RowsAffected()

	if err != nil {
		t.Fatal(err)
	}

	if affected != 1 {
		t.Fatal(err)
	}
	// 查询
	rows, err := db.QueryContext(context.Background(), "SELECT `id`, `first_name`,`age`, `last_name` FROM `test_model` LIMIT ?", 1)
	if err != nil {
		t.Fatal()
	}
	for rows.Next() {
		tm := &TestModel{}
		err = rows.Scan(&tm.id, &tm.FirstName, &tm.Age, &tm.LastName)
		if err != nil {
			rows.Close()
			t.Fatal()
		}
		assert.Equal(t, "Tom", tm.FirstName)
	}
	rows.Close()

	//upodate
	//或者exec(xxx)
	res, err = db.ExecContext(ctx, "UPDATE `test_model` SET `first_name` = 'changed' WHERE `id` = ?", 1)
	if err != nil {
		t.Fatal(err)
	}
	affected, err = res.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	if affected != 1 {
		t.Fatal(err)
	}

	//
	row := db.QueryRowContext(ctx, "SELECT `id`, `first_name`,`age`, `last_name` FROM `test_model` LIMIT 1")
	if row.Err() != nil {
		t.Fatal(row.Err())
	}
	tm := &TestModel{}

	err = row.Scan(&tm.id, &tm.FirstName, &tm.Age, &tm.LastName)
	if err != nil {
		t.Fatal(err)
	}
	// 这里测试里面会经常用
	assert.Equal(t, "changed", tm.FirstName)
}

func TestSQLite(t *testing.T) {
	suite.Run(t, &sqlTestSuite{
		driver: "sqlite3",
		dsn:    "file:test.db?cache=shared&mode=memory",
		//dsn: "test.db",
	})
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(0)
	fmt.Print(timer.Stop())
	<-timer.C
}
