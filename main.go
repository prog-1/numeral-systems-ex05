package main

import (
	"strings"
	"testing"

	sum "github.com/rchrdsasa207/numeral-systems-ex03"
)

func LongMul(a, b string) string {

	// Determine positive or negative result number
	isNegative := (b[0] == '-') != (a[0] == '-')
	// Remove math sign
	if a[0] == '-' || a[0] == '+' {
		a = a[1:]
	}
	if b[0] == '-' || b[0] == '+' {
		b = b[1:]
	}
	// Return zero if one of them is zero
	if a == "0" || b == "0" {
		return "0"
	}
	var res string
	var i int
	for j, v := range b {
		/*
			In base 2 theres only 0 and 1
			Anything multiplied by 0 is 0, that's why I skip multiplication by 0
			Any number multiplied by 1 remains the same, that's why I just add number to sum quee
		*/
		if v == '1' {
			res = sum.LongAddWithBase(res, a+strings.Repeat("0", len(b)-1-j), 2)
			i++
		}
	}
	// add minus if needed
	if isNegative {
		return "-" + res
	}
	return res
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

func main() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "TestLongMul", F: TestLongMul},
		},
		/* benchmarks */ nil,
		/* examples */ nil)
}
