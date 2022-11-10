package geek_homeworkv2

import (
	"errors"
	"reflect"
	"strings"
)

var errInvalidEntity = errors.New("invalid entity")

func InsertStmt(entity interface{}) (string, []interface{}, error) {
	// 实例为空的场景
	if entity == nil {
		return "", nil, errInvalidEntity
	}
	val := reflect.ValueOf(entity)
	typ := val.Type()
	// 判断指针的情况
	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = val.Type()
	}
	// 不是结构体类型，也给抛错
	if typ.Kind() != reflect.Struct {
		return "", nil, errInvalidEntity
	}

	bd := strings.Builder{}
	_, _ = bd.WriteString("INSERT INTO `")
	bd.WriteString(typ.Name()) // 这个是类型表名，也就是结构体名。
	bd.WriteString("`(")
	fields, values := fieldNameAndValues(val)
	for i, name := range fields {
		if i > 0 {
			bd.WriteRune(',')
		}
		bd.WriteRune('`')
		bd.WriteString(name)
		bd.WriteRune('`')
	}
	bd.WriteString(") VALUES(")
	args := make([]interface{}, 0, len(values))
	for i, fd := range fields {
		if i > 0 {
			bd.WriteRune(',')
		}
		bd.WriteRune('?')
		args = append(args, values[fd])
	}
	if len(args) == 0 {
		return "", nil, errInvalidEntity
	}
	bd.WriteString(");")
	return bd.String(), args, nil
}

func fieldNameAndValues(val reflect.Value) ([]string, map[string]interface{}) {
	typ := val.Type()
	fieldNum := typ.NumField()                       // 获取数量
	fields := make([]string, 0, fieldNum)            //初始化一个结构体切片
	values := make(map[string]interface{}, fieldNum) //初始化结构体值切片

	for i := 0; i < fieldNum; i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		//
		if field.Type.Kind() == reflect.Struct && field.Anonymous {
			// 递归，核心在这
			subFields, subValues := fieldNameAndValues(fieldVal)
			for _, k := range subFields {
				if _, ok := values[k]; ok {
					// 重复字段
					continue
				}
				fields = append(fields, k)
				values[k] = subValues[k]
			}
			continue
		}
		fields = append(fields, field.Name)
		values[field.Name] = fieldVal.Interface()
	}
	// 返回字段和字段值
	return fields, values

}
