package function

import (
	"fmt"
	"testing"
)

func TestPrintf(t *testing.T) {
	a, s := 1, "bowen"
	fmt.Printf("%v %v\n", a, s)
	fmt.Printf("%+v %+v\n", a, s)
	fmt.Printf("%#v %#v\n", a, s)

}

func TestStruct(t *testing.T) {
	type stu struct {
		id   int32
		name string
	}
	s := &stu{id: 1, name: "小明"}
	fmt.Printf("%v \n", s)
	fmt.Printf("%+v \n", s)
	fmt.Printf("%#v \n", s)

}
