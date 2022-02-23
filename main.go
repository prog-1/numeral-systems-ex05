package main

import (
	"strings"

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

func main() {
}
