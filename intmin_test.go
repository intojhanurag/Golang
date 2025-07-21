package main

import "testing"

func TestIntMin(t *testing.T) {
	got := IntMin(1, 2)
	if got != 1 {
		t.Errorf("Expected 1, got %d", got)
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
