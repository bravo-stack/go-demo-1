package main

import (
	"fmt"
)

func main() {

	var age uint8 = 23
	var amount float64 = 12345678.9

	var outp = float64(age) + amount
	var outp2 = age + uint8(amount)

	fmt.Println(age)
	fmt.Println(amount)
	fmt.Println(outp)
	fmt.Println(outp2)
}