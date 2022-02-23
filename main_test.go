package main

import (
	"strings"
	"testing"
)

var testLongMul = LongMul

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want string
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
		{"1001", "101", "101101"},
		{"1111", "-1111", "-11100001"},
		{"-1001", "-101", "101101"},
		{"+1001", "101", "101101"},
		{"+1001", "+101", "101101"},
		{"1101100100000011", "110111", "1011101001111110100101"},
	} {
		if got := testLongMul(tc.a, tc.b); got != tc.want {
			t.Errorf("LongSum(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
		}
	}
}

var (
	benchA = strings.Repeat("1", 10000)
	benchB = strings.Repeat("9", 1)
)

func BenchmarkLongMul(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = testLongMul(benchA, benchB)
	}
}
