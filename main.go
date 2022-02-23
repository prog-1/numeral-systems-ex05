package main

import (
	"fmt"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Byte2Number(a byte) int {
	for i, c := range base36 {
		if c == rune(a) {
			return i
		}
	}
	return -1
}

func LongAdd(a, b string, base int) (s string) {
	if len(a) < len(b) {
		a, b = b, a
	}
	iB := len(b)
	var r []rune
	tmp := 0
	for i := len(a) - 1; i >= 0; i-- {
		if iB > 0 {
			iB--
			if Byte2Number(a[i])+Byte2Number(b[iB])+tmp > base-1 {
				if Byte2Number(a[i])+Byte2Number(b[iB])+tmp > 9 {
					r = append(r, rune(base36[i]))
				} else {
					r = append(r, rune(Byte2Number(a[i])+Byte2Number(b[iB])-base+tmp))
					tmp = 1
				}

			} else {
				if Byte2Number(a[i])+Byte2Number(b[iB])+tmp > 9 {
					r = append(r, rune(base36[i]))
				} else {
					r = append(r, rune(Byte2Number(a[i])+Byte2Number(b[iB])+tmp))
					tmp = 0
				}

			}
		} else {
			if Byte2Number(a[i])+tmp > base-1 {
				if Byte2Number(a[i])+tmp > 9 {
					r = append(r, rune(base36[i]))
				} else {
					r = append(r, 0)
					r = append(r, 1)
					tmp = 0
				}

			} else {
				if Byte2Number(a[i])+tmp > 9 {
					r = append(r, rune(base36[i]))
				} else {
					r = append(r, rune(Byte2Number(a[i])+tmp))
					tmp = 0
				}

			}

		}

	}
	if tmp != 0 {
		r = append(r, 1)
	}

	var r2 []rune

	for i := len(r) - 1; i >= 0; i-- {
		r2 = append(r2, r[i]+48)
	}
	s = string(r2)
	return s
}
func LongMul(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	var carry byte
	var cnt int
	var res string
	for i := range b {
		var tmp []byte
		for c := 0; c < cnt; c++ {
			tmp = append(tmp, '0')
		}
		for i1 := range a {
			mul := carry + (a[len(a)-1-i1]-'0')*(b[len(b)-1-i]-'0')
			tmp = append(tmp, mul+'0')

		}
		for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
			tmp[i], tmp[j] = tmp[j], tmp[i]
		}

		res = LongAdd(string(res), string(tmp), 2)

		cnt++
	}

	if (a[0] == '-') != (b[0] == '-') {
		res = "-" + res[1:]

	}
	return string(res)

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