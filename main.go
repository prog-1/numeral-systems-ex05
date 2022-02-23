package main

import (
	"fmt"
	"strings"
)

func LongMul(a, b string) string {
	za, zb := a, b
	if strings.ContainsAny(a, "-") || strings.ContainsAny(a, "+") {
		a = a[1:]
	}
	if strings.ContainsAny(b, "-") || strings.ContainsAny(b, "+") {
		b = b[1:]
	}
	if a == "0" || b == "0" {
		return "0"
	}
	ab := make([]byte, len(a)+len(b))
	ia := 0
	i := len(a) - 1
	for ; i >= 0; i-- {
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
	if strings.IndexAny(res, "0") == 0 {
		res = strings.TrimPrefix(res, "0")
	}
	if strings.ContainsAny(za, "-") && !strings.ContainsAny(zb, "-") || strings.ContainsAny(zb, "-") && !strings.ContainsAny(za, "-") {
		res = "-" + res
	}
	return res
}

func main() {
	var a, b string
	fmt.Println("Examples:")
	for _, t := range []struct {
		a, b string
	}{
		{"1101100100000011", "110111"},
		{"1001", "101"},
	} {
		fmt.Printf("%v * %v = %v\n", t.a, t.b, LongMul(t.a, t.b))
	}
	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)
	fmt.Printf("%v * %v = %v\n", a, b, LongMul(a, b))
}
