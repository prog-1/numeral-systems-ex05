package main

import "testing"

func BenchmarkLongMul(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		LongMul("1101100100000011", "110111")
	}
}
