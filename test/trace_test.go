package test

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"testing"
)

// for trace test

func TestTrace(t *testing.T) {
	// go tool trace trace.out
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	trace.Start(f)
	defer f.Close()
	ch := make(chan int)
	go genOne(ch)
	for c := range ch {
		fmt.Println(c)
	}
	defer trace.Stop()
	fmt.Println("trace stopped")

}
