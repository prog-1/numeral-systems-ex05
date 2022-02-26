package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	ab := LongMul(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, LongMul(a, b))

}
