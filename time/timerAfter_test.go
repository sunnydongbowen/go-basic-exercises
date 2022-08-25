package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeAfter(t *testing.T) {
	cht := time.After(30 * time.Second)
	fmt.Printf("value = %v，type=…%T\n", cht, cht)

	time.Sleep(20 * time.Millisecond)
	fmt.Println("主goroutine主阻塞20ms")
	fmt.Println("cht", <-cht) // 注意这里取的就是设置的到期时间，如果没到期，到了这句话就阻塞在这里了
	time.Sleep(100 * time.Millisecond)

	//fmt.Println("cht", <-cht) //当前时间
	fmt.Println(cht) // 这样写打的是它的内存地址

}
