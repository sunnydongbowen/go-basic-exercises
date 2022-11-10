package geek_iterate

import (
	"errors"
	"reflect"
)

// Iterate 迭代数组，切片，或者字符串
func Iterate(input any) ([]any, error) {
	val := reflect.ValueOf(input)
	// 这里分开定义了
	typ := val.Type()
	kind := typ.Kind()

	if kind != reflect.Slice && kind != reflect.Array && kind != reflect.String {
		return nil, errors.New("非法类型")
	}
	res := make([]any, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		ele := val.Index(i)
		res = append(res, ele.Interface())
	}
	return res, nil
}

func IterateString(input any) ([]any, error) {
	val := reflect.ValueOf(input)
	// 这里分开定义了
	typ := val.Type()
	kind := typ.Kind()

	if kind != reflect.Slice && kind != reflect.Array && kind != reflect.String {
		return nil, errors.New("非法类型")
	}
	res := make([]any, 0, val.Len())
	if kind == reflect.Slice || kind != reflect.Array {
		for i := 0; i < val.Len(); i++ {
			ele := val.Index(i)
			res = append(res, ele.Interface())
		}
	}

	//resString := make([]byte, 0, val.Len())
	//if kind == reflect.String {
	//	for i := 0; i < val.Len(); i++ {
	//		ele := val.Index(i)
	//		resString = append(resString, string(ele.Interface()))
	//	}
	//}

	return res, nil
}

func IterateMapV1(input any) ([]any, []any, error) {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Map {
		return nil, nil, errors.New("不是map类型")
	}
	l := val.Len()
	keys := make([]any, 0, l)
	values := make([]any, 0, l)

	for _, k := range val.MapKeys() {
		keys = append(keys, k.Interface())
		v := val.MapIndex(k)
		values = append(values, v.Interface())
	}
	return keys, values, nil

}

func IterateMapV2(input any) ([]any, []any, error) {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Map {
		return nil, nil, errors.New("不是map类型")
	}
	l := val.Len()
	keys := make([]any, 0, l)
	values := make([]any, 0, l)

	itr := val.MapRange()
	for itr.Next() {
		keys = append(keys, itr.Key().Interface())
		values = append(values, itr.Value().Interface())
	}
	return keys, values, nil
}
