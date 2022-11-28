package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// @program:     go-basic-exercises
// @file:        mkdir_test.go
// @author:      bowen
// @create:      2022-11-28 22:19
// @description:  go汪明，go并发编程笔记

func CreateDir(Path string) string {
	foldername := time.Now().Format("20060102")
	folderpath := filepath.Join(Path, foldername)
	// 获取特定路径的文件属性，它返回一个FIleInfo对象和error对象
	// os.IsnotExit 可以判断一个目录或文件是否存在，如果不存在，可以调用mkdir方法创建目录
	if _, err := os.Stat(folderpath); os.IsNotExist(err) {
		// 先创建文件夹
		os.MkdirAll(folderpath, os.ModePerm)
		// 修改权限
		os.Chmod(folderpath, 0777)
	}
	return folderpath

}
func TestMkd(t *testing.T) {
	// 获取当前的绝对路径
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	path = CreateDir(path)
	fmt.Println(path)
	// mkdirAll 可以用来创建多级目录， mkdir只创建单级目录
	os.MkdirAll(filepath.Join(path, "新文件夹/dic"), os.ModePerm)
	// 重命名或者移动目录
	os.Rename(filepath.Join(path, "新文件夹"), filepath.Join(path, "newfolder"))
	//删除目录，但是有子目录时不删除
	os.Remove(filepath.Join(path, "newfolder/dic"))

}
