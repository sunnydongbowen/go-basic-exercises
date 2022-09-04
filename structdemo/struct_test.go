package structdemo

//for 结构体工厂
import (
	"fmt"
	"testing"
)

func TestStructFactory(t *testing.T) {
	obj := NewFile(10, "xxx")
	fmt.Println(obj)

	m := NewMatrix("bowen")
	fmt.Print(m)
}
