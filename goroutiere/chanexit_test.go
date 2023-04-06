package goroutiere

import (
	"fmt"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      chanexit_test.go
// @author:    bowen
// @time:      2023-04-03 11:40
// @description:

var exitChan = make(chan bool, 1)

func fchan() {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("001")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break FORLOOP
		default:
		}
	}
}
func TestChanExit(t *testing.T) {
	wg.Add(1)
	go fchan()
	time.Sleep(5 * time.Second)
	exitChan <- true
	wg.Wait()
}
