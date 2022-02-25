package main

import (
	"fmt"
)

func LongMul(a, b string) (output string) {
	var sign int
	if (a[0] == '-') != (b[0] == '-') {
		sign = 1
	}
	if a[0] == '-' || a[0] == '+' {
		a = a[1:]
	}
	if b[0] == '-' || b[0] == '+' {
		b = b[1:]
	}
	if a == "0" || b == "0" {
		return "0"
	}
	ab := make([]byte, len(a)+len(b))
	var ia int
	for i := len(a) - 1; i >= 0; i-- {
		var carry byte
		ib := 0
		for j := len(b) - 1; j >= 0; j-- {
			mul := (a[i]-'0')*(b[j]-'0') + ab[ia+ib] + carry
			carry = mul / 2
			ab[ia+ib] = mul % 2
			ib++
		}
		if carry > 0 {
			ab[ia+ib] += carry
		}
		ia++
	}
	var res string
	for i := len(ab) - 1; i >= 0; i-- {
		res += fmt.Sprint(ab[i])
	}
	if res[0] == '0' {
		res = res[1:]
	}
	if sign == 1 {
		res = "-" + res
	}

	return res
}

func main() {
	for _, t := range []struct {
		a, b string
	}{
		// Checking signed operations.
		{"10", "101"},
		{"+10", "101"},
		{"10", "+101"},
		{"+10", "+101"},
		{"-10", "101"},
		{"10", "-101"},
		{"-10", "-101"},
		{"+10", "-101"},
		{"-10", "+101"},
		// Resulting zero must never contain sign literals.
		{"0", "0"},
		{"-0", "0"},
		{"0", "-0"},
		{"1", "0"},
		{"1", "+0"},
		{"1", "-0"},
		{"+1", "-0"},
		{"-1", "+0"},
	} {
		fmt.Printf("%v * %v = %v\n", t.a, t.b, LongMul(t.a, t.b))
	}
}
