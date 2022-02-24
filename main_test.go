package main

import (
	"strings"
	"testing"
)

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want string
	}{
		// Checking signed operations.
		{"10", "101", "1010"},
		{"+10", "101", "1010"},
		{"10", "+101", "1010"},
		{"+10", "+101", "1010"},
		{"-10", "101", "-1010"},
		{"10", "-101", "-1010"},
		{"-10", "-101", "1010"},
		{"+10", "-101", "-1010"},
		{"-10", "+101", "-1010"},
		// Resulting zero must never contain sign literals.
		{"0", "0", "0"},
		{"-0", "0", "0"},
		{"0", "-0", "0"},
		{"1", "0", "0"},
		{"1", "+0", "0"},
		{"1", "-0", "0"},
		{"+1", "-0", "0"},
		{"-1", "+0", "0"},

		// Other examples.
		{"1101100100000011", "110111", "1011101001111110100101"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := LongMul(tc.a, tc.b); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

var (
	benchA = strings.Repeat("1", 10000)
	benchB = strings.Repeat("0", 1)

	// benchA = strings.Repeat("1", 10000)
	// benchB = strings.Repeat("1", 200)

	// benchA = strings.Repeat("101101", 10000)
	// benchB = strings.Repeat("10", 1000)
)

func BenchmarkLongMul(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = LongMul(benchA, benchB)
	}
}

func BenchmarkLongMulAnotherVersion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = LongMulAnotherVersion(benchA, benchB)
	}
}

// Conclusion #1: The higher the numbers, the more ns, B and allocs the program needs to complete the task.
// Conclusion #2: My another (optimized) version of this program doesn't really do anything faster or better.
