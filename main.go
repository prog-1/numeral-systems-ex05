package main

import (
	"fmt"
)

func LongMul(a, b string) string {
	var isNegative bool
	if a[0] == '-' || b[0] == '-' {
		isNegative = true
	}
	if a[0] == '-' && b[0] == '-' {
		isNegative = false
	}
	if a[0] == '+' || a[0] == '-' {
		a = a[1:]
	}
	if b[0] == '+' || b[0] == '-' {
		b = b[1:]
	}
	if a == "0" || b == "0" {
		return "0"
	}
	ab := make([]byte, len(a)+len(b))
	for i, ia, ib := len(a)-1, 0, 0; i >= 0; i-- {
		var carry byte
		ib = 0
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
	var result string
	for i := len(ab) - 1; i >= 0; i-- {
		result += fmt.Sprint(ab[i])
	}
	if result[0] == '0' {
		result = result[1:]
	}
	if isNegative {
		return "-" + result
	}
	return result
}

func main() {
	var a, b string
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	ab := LongMul(a, b)
	fmt.Println("The multiplication of numbers: ", ab)
}
