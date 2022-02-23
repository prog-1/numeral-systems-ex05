package main

import (
	"fmt"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsValidNumber(num string, base int) bool {
	for _, v := range num {
		if !strings.Contains(base36[:base], string(v)) {
			return false
		}
	}
	return true
}
func LongAddBase(a, b string, base int) (sum string) {
	var carry byte
	if len(a) > len(b) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	} else {
		a = strings.Repeat("0", len(b)-len(a)) + a
	}
	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0') + carry
		carry = num / byte(base)
		sum = string(num%byte(base)+'0') + sum
	}
	if carry != 0 {
		sum = "1" + sum
	}
	return sum
}

func LongMul(a, b string) (output string) {
	var minus int
	if a[0] == '-' && b[0] != '-' || a[0] != '-' && b[0] == '-' {
		minus = 1
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
	toSum := make([]string, strings.Count(b, "1"))
	var i int
	for j, v := range b {
		if v == '1' {
			toSum[i] = a + strings.Repeat("0", len(b)-1-j)
			i++
		}
	}
	for _, v := range toSum {
		output = LongAddBase(output, v, 2)
	}
	if minus == 1 {
		output = "-" + output
	}
	return output
}

func main() {
	var a, b string
	base := 2
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	fmt.Println(LongMul(a, b))
	if IsValidNumber(a, base) && IsValidNumber(b, base) {
		fmt.Printf("Multiplication of %v and %v = %v", a, b, LongMul(a, b))
	} else {
		fmt.Printf("The number %v or number %v is represented incorrect in base%v", a, b, base)
	}
}
