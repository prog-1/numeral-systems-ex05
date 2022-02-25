package main

import "testing"

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		a    string
		b    string
		want string
	}{
		{"0", "0", "0"},
		{"-0", "0", "0"},
		{"0", "-0", "0"},
		{"1", "0", "0"},
		{"1", "+0", "0"},
		{"1", "-0", "0"},
		{"+1", "-0", "0"},
		{"-1", "+0", "0"},
		{"10", "101", "1010"},
		{"+10", "101", "1010"},
		{"10", "+101", "1010"},
		{"+10", "+101", "1010"},
		{"-10", "101", "-1010"},
		{"10", "-101", "-1010"},
		{"-10", "-101", "1010"},
		{"+10", "-101", "-1010"},
		{"-10", "+101", "-1010"},
		{"1101100100000011", "110111", "1011101001111110100101"},
	} {
		if got := LongMul(tc.a, tc.b); got != tc.want {
			t.Errorf("LongMul(%v, %v): got = %v, want = %v", tc.a, tc.b, got, tc.want)
		}
	}
}

var benchA, benchB string = "1010100101011111", "-1010101001111010101010"

func BenchmarkLongMulWithLongAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LongMul(benchA, benchB)
	}
}
