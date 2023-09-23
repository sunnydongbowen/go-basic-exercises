package dao

import "gorm.io/gorm"

// @program:   go-basic-exercises
// @file:      init.go
// @author:    bowen
// @time:      2023-09-22 10:33
// @description:

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
