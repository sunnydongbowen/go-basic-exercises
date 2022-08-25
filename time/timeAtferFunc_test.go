package time

import (
	"fmt"
	"testing"
	"time"
)

var notify bool

// 注意文件名_test结尾
func TestTimeAfterFunc(t *testing.T) {
	// 注意先传时间，这函数返回的是Timer指针
	timer := time.AfterFunc(30*time.Millisecond, func() {
		fmt.Println("我再等30ms开始执行,你可以忙别的")
	})
	time.Sleep(20 * time.Millisecond)
	if notify {
		timer.Stop()
		fmt.Println("收到通知，任务已终止")
	}
	fmt.Println("主函数开始终止")
}
