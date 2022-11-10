package batch_insertv1

import (
	"fmt"
	"go-basic-exercises/mysql"
	"testing"
)

// 批量插入
func TestBatchInsert(t *testing.T) {
	// 连接数据库
	if err := mysql.InitDB(); err != nil {
		fmt.Println("initDB failed,err: ", err)
		return
	}
	// 产生userslist，这种一次性生成的插入取决于Generate的数量
	users := mysql.Generate()
	// 执行插入数据
	mysql.BatchInsert(users)

	defer mysql.DBX.Close()
}

// 测试产生数据
func TestGen(t *testing.T) {
	mysql.Generate()
}
