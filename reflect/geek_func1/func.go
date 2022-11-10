package geek_func1

import (
	"errors"
	"reflect"
)

type FuncInfo struct {
	Name string         // 方法名
	In   []reflect.Type // 方法输入参数
	Out  []reflect.Type // 方法输出参数

	// 反射调用的结果
	Result []any
}

// IterateFuncs
// 输出方法信息，并执行调用

// nil 基本类型，内置类型(切片，map，channel)之类的
// 结构体，结构体指针，多级结构体

// 决策支持 :只支持结构体和结构体指针

func IterateFuncs(val any) (map[string]*FuncInfo, error) {
	//判断为空
	if val == nil {
		return nil, errors.New("输入nil")
	}
	// 只支持结构体和一级指针
	typ := reflect.TypeOf(val)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("不支持该类型")
	}
	// 正儿八经的方法，返回数量
	numMethod := typ.NumMethod()
	res := make(map[string]*FuncInfo, numMethod) //切片
	// In
	for i := 0; i < numMethod; i++ {
		method := typ.Method(i)
		mt := method.Type
		numIn := mt.NumIn()
		in := make([]reflect.Type, 0, numIn)
		for j := 0; j < numIn; j++ {
			in = append(in, mt.In(j))
		}
		// Out和上面一样
		numOut := mt.NumOut()
		out := make([]reflect.Type, 0, numOut)
		for j := 0; j < numOut; j++ {
			out = append(out, mt.Out(j))
		}

		callRes := method.Func.Call([]reflect.Value{reflect.ValueOf(val)})
		retVals := make([]any, 0, len(callRes))
		for _, cr := range callRes {
			retVals = append(retVals, cr.Interface()) // 反射形态的值转换成正常的值
		}

		res[method.Name] = &FuncInfo{
			Name:   method.Name,
			In:     in,
			Out:    out,
			Result: retVals,
		}
	}
	return res, nil
}
