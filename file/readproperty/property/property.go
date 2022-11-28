package property

import (
	"bufio"
	"fmt"
	"go-basic-exercises/file/selfio/myio"
	"strings"
)

// @program:     go-basic-exercises
// @file:        property.go
// @author:      bowen
// @create:      2022-11-29 12:32
// @description:

// Read读取config。Property，没有缓存
// 根据传入的字符串的key从配置文件config.property中获取key对应的值
func Read(key string) string {
	fio, err := myio.Open("config.property")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//构建scanner，并且逐行扫描，
	sc := bufio.NewScanner(fio)
	line := ""
	slice := make([]string, 2)
	for sc.Scan() {
		line = sc.Text()
		// 注释，忽略掉
		if strings.HasPrefix(line, "#") {
			continue
		}
		slice = strings.Split(line, "=")
		if len(slice) == 2 {
			// 一旦匹配key就讲获取到的对应的值返回。
			if strings.Trim(slice[0], " ") == key {
				return strings.Trim(slice[1], " ")
			} else {
				//fmt.Println("key不存在")
				return "key不存在"
			}
		}
	}
	err = fio.Close()
	if err != nil {
		fmt.Println(err)
	}
	return ""

}
