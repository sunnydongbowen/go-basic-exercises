package main

// 注意这种交互方式是没办法以测试方式进行的，所以我加上了mian包，并且加了mian函数，
// 把go文件名的_test去掉了。
import (
	"bufio"
	"fmt"
	"os"
)

func useScan() {
	var s string
	fmt.Print("请输入内容:")
	fmt.Scanln(&s)
	// 没有换行
	fmt.Printf("你输入的内容是:%s\n", s)
}

func useBuffer() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容:")
	s, _ := reader.ReadString('\n')
	fmt.Println("你输入的内容是:", s)
}
func main() {
	useScan()
	//useBuffer()
}
