package sqlgeek

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// @program:     go-basic-exercises
// @file:        json.go
// @author:      bowen
// @create:      2022-12-08 15:45
// @description:

// JsonColumn 代表存储字段的 json 类型
// 主要用于没有提供默认 json 类型的数据库
// T 可以是结构体，也可以是切片或者 map
// 一切可以被 json 库所处理的类型都能被用作 T

type JsonColum[T any] struct {
	VAL   T
	Valid bool
}

// Value 返回一个 json 串。类型是 []byte
func (j JsonColum[T]) Value() (driver.Value, error) {
	if !j.Valid {
		return nil, nil
	}
	return json.Marshal(j.VAL)
}

// Scan 将 src 转化为对象
// src 的类型必须是 []byte, *[]byte, string, sql.RawBytes, *sql.RawBytes 之一

func (j *JsonColum[T]) Scan(src any) error {
	var bs []byte
	switch val := src.(type) {
	case []byte:
		bs = val
	case *[]byte: // 如果是字节型指针，就取指针给bs
		bs = *val
	case string: // 如果是string类型，就转换成字节再给bs
		bs = []byte(val)
	case *sql.RawBytes:
		bs = *val
	case sql.RawBytes:
		bs = val
	default:
		return fmt.Errorf("ekit:JsonColumn.Scan 不支持src类型%v\n", src)
	}
	if err := json.Unmarshal(bs, &j.VAL); err != nil {
		return err
	}
	j.Valid = true
	return nil
}
