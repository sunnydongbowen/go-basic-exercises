package fanxing

import "testing"

// @program:     go-basic-exercises
// @file:        maxbench_test.go
// @author:      bowen
// @create:      2022-12-11 12:48
// @description: benchmark测试

func BenchmarkMaxInt(b *testing.B) {
	sl := []int{1, 2, 3, 4, 7, 8, 9, 0}
	for i := 0; i < b.N; i++ {
		maxInt(sl)
	}
}

func BenchmarkMaxAny(b *testing.B) {
	sl := []any{1, 2, 3, 4, 7, 8, 9, 0}
	for i := 0; i < b.N; i++ {
		maxAny(sl)
	}
}

func BenchmarkMaxGeneric(b *testing.B) {
	sl := []int{1, 2, 3, 4, 7, 8, 9, 0}
	for i := 0; i < b.N; i++ {
		maxGenerics(sl)

	}
}
