package test

import "testing"

func TestCounter(t *testing.T) {
	tests := []struct {
		name string
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Counter(); got != tt.want {
				t.Errorf("Counter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounter2(t *testing.T) {
	tests := []struct {
		name string
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Counter2(); got != tt.want {
				t.Errorf("Counter2() = %v, want %v", got, tt.want)
			}
		})
	}
}
