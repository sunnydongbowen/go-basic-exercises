package self

import (
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"reflect"
)

// @program:     go-basic-exercises
// @file:        General.go
// @author:      bowen
// @create:      2022-11-12 19:56
// @description: 给一个参数，判断是什么类型，如果是基本类型，就给出基本类型()
// 如果是结构体，就给出结构体的字段以及value
// 如果是一级指针，给出具体类型
// 二级指针不做判断

func reflectapp(entidy any) (err error) {
	val := reflect.ValueOf(entidy)
	// 传参为nil
	if entidy == nil {
		return errors.New("传参不能为空")
	}
	// 传参胃
	switch val.Kind() {
	case reflect.String:
		fmt.Printf("字符串类型,值为%v\n",)
	}
	case

	return nil
}
