package main

import (
	"fmt"
	"strings"
)

func LongMul(a, b string) string {
	isNegative := (b[0] == '-') != (a[0] == '-')
	if a[0] == '-' || a[0] == '+' {
		a = a[1:]
	}
	if b[0] == '-' || b[0] == '+' {
		b = b[1:]
	}
	if a == "0" || b == "0" {
		return "0"
	}
	str := make([]byte, len(a)+len(b))
	ia := 0
	i := len(a) - 1
	for ; i >= 0; i-- {
		var carry byte
		ib := 0
		for j := len(b) - 1; j >= 0; j-- {
			mul := (a[i]-'0')*(b[j]-'0') + str[ia+ib] + carry
			carry = mul / 2
			str[ia+ib] = mul % 2
			ib++
		}
		if carry > 0 {
			str[ia+ib] += carry
		}
		ia++
	}
	var res string
	for i := len(str) - 1; i >= 0; i-- {
		res += fmt.Sprint(str[i])
	}
	if strings.IndexAny(res, "0") == 0 {
		res = strings.TrimPrefix(res, "0")
	}
	if isNegative {
		return "-" + res
	}
	return res
}

func main() {
	var a, b string
	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)
	fmt.Println(LongMul(a, b))
}
