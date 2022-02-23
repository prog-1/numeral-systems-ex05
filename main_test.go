package main

import (
	"strings"
	"testing"
)

var testLongMul = LongMul

func TestLongMul(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"from README", "10", "101", "1010"},
		{"from README", "+10", "101", "1010"},
		{"from README", "10", "+101", "1010"},
		{"from README", "+10", "+101", "1010"},
		{"from README", "-10", "101", "-1010"},
		{"from README", "10", "-101", "-1010"},
		{"from README", "-10", "-101", "1010"},
		{"from README", "+10", "-101", "-1010"},
		{"from README", "-10", "+101", "-1010"},

		{"from README", "0", "0", "0"},
		{"from README", "-0", "0", "0"},
		{"from README", "0", "-0", "0"},
		
		{"from README", "1", "0", "0"},
		{"from README", "1", "+0", "0"},
		{"from README", "1", "-0", "0"},
		{"from README", "+1", "-0", "0"},
		{"from README", "-1", "+0", "0"},
		{"from README", "1101100100000011", "110111", "1011101001111110100101"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongMul(tc.a, tc.b); got != tc.want {
				t.Errorf("LongMul(%v, %v) = %v, want = %v", tc.a, tc.b, got, tc.want)
			}
		})
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
