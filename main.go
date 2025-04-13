/*
	default values for uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, rune is 0
	default values for string is "" (empty string)
	default values for bool is false
*/

package main

import (
	"fmt"
)

func main() {
	printMe("Hello, World!")
	var numerator int = 118234231
	var denominator int = 3
	var result int = int(intDivision(uint8(numerator), uint8(denominator)))
	fmt.Println("Result of intDivision is: ", result)
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

func intDivision (var1 uint8, var2 uint8) uint8 {
	var result uint8 = var1 / var2
	return result
}