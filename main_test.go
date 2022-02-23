package main

import "testing"

func BenchmarkLongMul(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		LongMul("1101100100000011", "110111")
	}
}

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		a, b, want string
	}{
		{"10", "101", "1010"},
		{"+10", "101", "1010"},
		{"10", "+101", "1010"},
		{"+10", "+101", "1010"},
		{"-10", "101", "-1010"},
		{"10", "-101", "-1010"},
		{"-10", "-101", "1010"},
		{"+10", "-101", "-1010"},
		{"-10", "+101", "-1010"},
		{"0", "0", "0"},
		{"-0", "0", "0"},
		{"0", "-0", "0"},
		{"1", "0", "0"},
		{"1", "+0", "0"},
		{"1", "-0", "0"},
		{"+1", "-0", "0"},
		{"-1", "+0", "0"},
		{"1101100100000011", "110111", "1011101001111110100101"},
	} {
		if got := LongMul(tc.a, tc.b); got != tc.want {
			t.Errorf("LongMul(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
		}
	}
}
