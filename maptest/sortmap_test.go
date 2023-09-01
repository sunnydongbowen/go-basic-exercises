package maptest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      sortmap_test.go
// @author:    bowen
// @time:      2023-04-22 9:54
// @description:

func TestSortTest(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var scoremap = make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%003d", i)
		value := rand.Intn(100)
		scoremap[key] = value
	}
	fmt.Println(scoremap)

	for k, v := range scoremap {
		fmt.Println(k, v)
	}

	//var keys = make([]string, 0, 200)
	//for key := range scoremap {
	//	keys = append(keys, key)
	//}
	//sort.Strings(keys)
	//for _, key := range keys {
	//	fmt.Println(key, scoremap[key])
	//}
	//fmt.Println(scoremap)
}
