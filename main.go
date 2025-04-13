/*
	default values for uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, rune is 0
	default values for string is "" (empty string)
	default values for bool is false
	default value for type error is nil
*/

package main

import (
	"fmt"
)

func main() {
	printMe("Hello, World!")
	var numerator uint8 = 131
	var denominator uint8 = 3
	var result, remainder uint8 = intDivision(numerator, denominator)
	fmt.Println("Result of intDivision is: ", result)
	fmt.Println("Remainder of intDivision is: ", remainder)
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

func intDivision (var1 uint8, var2 uint8) (uint8, uint8) {
	var result uint8 = var1 / var2
	var remainder uint8 = var1 % var2
	return result, remainder
}