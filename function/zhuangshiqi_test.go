package function

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

// tony专栏 演示装饰器
func TestZhuang(t *testing.T) {
	r := strings.NewReader("hello, gopher!\n")
	lr := io.LimitReader(r, 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

}
