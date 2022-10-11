package const_t

import (
	"fmt"
	"testing"
)

const (
	a = iota + 1 // 1, iota = 0
	b            // 2, iota = 1
	c            // 3, iota = 2
)

func TestConst(t *testing.T) {
	fmt.Println(a, b, c)
}
