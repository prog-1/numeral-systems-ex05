package main

import (
	"fmt"
)

func LongMul(a, b string) string

func main() {
	fmt.Println("Program implements signed arithmetics multiplication, which accepts two binary numbers encoded as strings.")
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Println("Answer:", LongMul(a, b))
}
