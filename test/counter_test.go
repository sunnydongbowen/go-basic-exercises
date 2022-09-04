package test

// benchmark test
import "testing"

func BenchmarkCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Counter()
	}
}

func BenchmarkCounter2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Counter2()
	}
}
