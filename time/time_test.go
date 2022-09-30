package time

import (
	"fmt"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	fmt.Println(time.Second)
	//now+1h
	now := time.Now()
	//当前时间+24h
	fmt.Println(now.Add(100 * time.Da))
}
