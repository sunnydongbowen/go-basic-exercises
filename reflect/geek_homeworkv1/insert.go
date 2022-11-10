package geek_homeworkv1

import (
	"errors"
	"reflect"
	"strings"
)

var errInvalidEntity = errors.New("invalid entity")

// InsertStmt 作业里面我们这个只是生成 SQL，所以在处理 sql.NullString 之类的接口
// 只需要判断有没有实现 driver.Valuer 就可以了
func InsertStmt(entity interface{}) (string, []interface{}, error) {
	var typ reflect.Type
	val := reflect.ValueOf(entity)
	//typ := val.Type()
	if val.Kind() == reflect.Ptr && val.Elem().Kind() != reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Ptr {
		return "", nil, errors.New("参数不能为多级指针")

	}

	if entity != nil {
		typ = val.Type()
	}

	// 检测 entity 是否符合我们的要求
	// 我们只支持有限的几种输入

	// 判断空
	if entity == nil {
		return "", nil, errors.New("参数不能为nil")
	}
	// 判断空结构体
	if typ.NumField() == 0 {
		return "", nil, errors.New("结构体不能为空")
	}
	// 第三条 "simple struct
	a := []interface{}{}
	numField := typ.NumField()
	for i := 0; i < numField; i++ {
		a = append(a, val.Field(i).Interface())
	}
	//return "", a, nil
	// basic结构体，最后要拼接一个sql出来，利用反射取到这个结构体的字段.

	bd := strings.Builder{}
	if typ.Kind() == reflect.Struct {
		names := make([]string, 0, 2)
		for i := 0; i < numField; i++ {
			name := typ.Field(i).Name
			names = append(names, name)
		}

		// 使用 strings.Builder 来拼接 字符串,这个怎么用？？？ gorm也是用的这个
		bd.WriteString("INSERT INTO `BaseEntity`")
		bd.WriteString("(`")
		bd.WriteString(names[0])
		bd.WriteString("`,`")
		bd.WriteString(names[1])
		bd.WriteString("`)")
		bd.WriteString(" VALUES(?,?);")
	}

	// 第四个case 指针在上面判断过了
	// 多级指针

	//if typ.Kind() == reflect.Ptr {
	//	val := reflect.ValueOf(&entity)
	//	typ := val.Type()
	//}

	//	"INSERT INTO `BaseEntity`(`CreateTime`,`UpdateTime`) VALUES(?,?);",
	// 构造 INSERT INTO XXX，XXX 是你的表名，这里我们直接用结构体名字

	// 遍历所有的字段，构造出来的是 INSERT INTO XXX(col1, col2, col3)
	// 在这个遍历的过程中，你就可以把参数构造出来
	// 如果你打算支持组合，那么这里你要深入解析每一个组合的结构体
	// 并且层层深入进去

	// 拼接 VALUES，达成 INSERT INTO XXX(col1, col2, col3) VALUES

	// 再一次遍历所有的字段，要拼接成 INSERT INTO XXX(col1, col2, col3) VALUES(?,?,?)
	// 注意，在第一次遍历的时候我们就已经拿到了参数的值，所以这里就是简单拼接 ?,?,?

	return bd.String(), a, nil
	//panic("implement me")
}
