package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// @program:     go-basic-exercises
// @file:        dirIter_test.go
// @author:      bowen
// @create:      2022-11-28 23:03
// @description: go实战编程笔记

// 遍历目录和文件
func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("%s\n", path)
	return nil
}

func TestDirIter(t *testing.T) {
	// 遍历当前目录
	if root, err := os.Getwd(); err == nil {
		// 遍历指定目录，该方法接收两个参数，一个是根目录root，一个是递归的回调函数walkFn，
		err := filepath.Walk(root, visit)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
