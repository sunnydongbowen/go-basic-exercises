package string

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var s string
	s = "bowen"
	fmt.Println(len(s))
	for i, v := range s {
		fmt.Printf("%d %v\n", i, string(v)) // %v的话不会输出字符串本身，输出的是字符串的对应的ASCII值,所以后面要有string(v)
	}

	s1 := "Hello 博文"
	for i, v := range s1 {
		fmt.Printf("%d %c\n", i, v) //输出索引，和索引对应的值，c表示输出单个字符，d表示十进制整数
	}

}
