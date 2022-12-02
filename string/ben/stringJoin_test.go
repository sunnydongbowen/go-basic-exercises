package ben

import "testing"

// @program:     go-basic-exercises
// @file:        stringJoin_test.go
// @author:      bowen
// @create:      2022-12-01 19:45
// @description:

func BenchmarkConcatStringByOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByOperator(sl)
	}
}

func BenchmarkConcatStringBySprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringBySprintf(sl)
	}
}

func BenchmarkConcatStringByJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByJoin(sl)
	}
}

func BenchmarkStringByStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByStringsBuilder(sl)
	}
}

func BenchmarkConcatStringByStringsBuilderWithInitSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByStringsBuilderWithInitSize(sl)
	}
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByBytesBuffer(sl)
	}
}

func BenchmarkConcatStringByBytesBufferWithInitSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringByBytesBufferWithInitSize(sl)
	}
}
