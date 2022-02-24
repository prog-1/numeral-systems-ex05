package main

import (
	"fmt"
	"strings"
)

func NumberExistsInBase2(num string) bool {
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, v := range num {
		if !strings.Contains(base36[:2], string(v)) {
			return false
		}
	}
	return true
}

func LongMul(a, b string) string {
	var s string
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
	m := make([]byte, len(a)+len(b))
	if NumberExistsInBase2(a) && NumberExistsInBase2(b) {
		for i, l := len(a)-1, 0; i >= 0; i-- {
			var add uint8 // uint8 = byte
			k := 0
			for j := len(b) - 1; j >= 0; j-- {
				num := (a[i]-'0')*(b[j]-'0') + m[l+k] + add // '0' = 48
				add = num / 2
				m[l+k] = num % 2
				k++
			}
			if add > 0 {
				m[l+k] = m[l+k] + add
			}
			l++
		}
		for i := len(m) - 1; i >= 0; i-- {
			s = s + fmt.Sprint(m[i])
		}
		if s[0] == '0' {
			s = s[1:]
		}
		if isNegative {
			s = "-" + s
		}
	} else {
		s = "-1"
	}
	return s
}

func LongMulAnotherVersion(a, b string) string {
	var s string
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
	m := make([]byte, len(a)+len(b))
	if NumberExistsInBase2(a) && NumberExistsInBase2(b) {
		for i, l, k := len(a)-1, 0, 0; i >= 0; i-- {
			var add uint8 // uint8 = byte
			k = 0
			for j := len(b) - 1; j >= 0; j-- {
				num := (a[i]-'0')*(b[j]-'0') + m[l+k] + add // '0' = 48
				add = num / 2
				m[l+k] = num % 2
				k++
			}
			if add > 0 {
				m[l+k] += add
			}
			l++
		}
		for i := len(m) - 1; i >= 0; i-- {
			s += fmt.Sprint(m[i])
		}
		if s[0] == '0' {
			return s[1:]
		}
		if isNegative {
			return "-" + s
		}
	} else {
		s = "-1"
	}
	return s
}

func main() {
	var a, b string
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	if LongMul(a, b) != "-1" {
		fmt.Println("The product of a and b:", LongMul(a, b))
	} else {
		fmt.Println("ERROR: At least one of the entered numbers does not exist in base 2")
	}
}
