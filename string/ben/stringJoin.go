package ben

import (
	"bytes"
	"fmt"
	"strings"
)

// @program:     go-basic-exercises
// @file:        stringJoin.go
// @author:      bowen
// @create:      2022-12-01 18:57
// @description:

var sl []string = []string{
	"Rob Pike ", "Robert Griesemer ", "Ken Thompson ",
}

// + 号拼接
func concatStringByOperator(sl []string) string {
	var s string
	for _, v := range sl {
		s += v
	}
	return s
}

// 利用fmt.Sprintf拼接
func concatStringBySprintf(sl []string) string {
	var s string
	for _, v := range sl {
		//格式化拼接，两个拼一起返回一个字符串
		s = fmt.Sprintf("%s%s", s, v)
	}
	return s
}

// 利用Strings.Join函数拼接
func concatStringByJoin(sl []string) string {
	return strings.Join(sl, "")
}

// strings.Builder 这个写法其实gorm构建sql就是用的这个
func concatStringByStringsBuilder(sl []string) string {
	var b strings.Builder
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}

// strings.Builder 更高级了一点，用的Grow了，这个后面再去了解这个是什么
func concatStringByStringsBuilderWithInitSize(sl []string) string {
	var b strings.Builder
	b.Grow(64)
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBuffer(sl []string) string {
	var b bytes.Buffer
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBufferWithInitSize(sl []string) string {
	buf := make([]byte, 0, 64)
	b := bytes.NewBuffer(buf)
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}
