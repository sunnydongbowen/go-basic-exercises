package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

// @program:   go-basic-exercises
// @file:      user.go
// @author:    bowen
// @time:      2023-09-22 9:05
// @description:

type UserDAO struct {
	db *gorm.DB
}

var (
	ErrEmailDuplicate = errors.New("邮箱冲突")
	ErrUserNotfound   = gorm.ErrRecordNotFound
)

// 与数据库对应的字段
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	//唯一索引
	Email    string `gorm:"unique"`
	Password string

	Ctime int64
	Utime int64
}

// 分两张表
type UserDetail struct {
	//生日，什么什么的
}

func NewUserDao(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

// 真正的入库程序,注册
func (dao *UserDAO) Insert(c context.Context, u User) error {
	// 存ms数
	now := time.Now().UnixMilli()

	u.Utime = now
	u.Ctime = now
	err := dao.db.WithContext(c).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == 1062 {
			//邮箱冲突
			return ErrEmailDuplicate
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(c context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(c).Where("email = ?", email).First(&u).Error
	return u, err
}
