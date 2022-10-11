package function

import (
	"fmt"
	"testing"
)

func sum(para ...int) int {
	sum := 0
	for _, v := range para {
		sum = sum + v
	}
	return sum
}

func TestPara(t *testing.T) {
	sum1 := sum(1, 2)
	sum2 := sum(1, 2, 3)
	fmt.Println(sum1, sum2)

}
