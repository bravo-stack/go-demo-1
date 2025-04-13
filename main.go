/*
	default values for uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64, rune is 0
	default values for string is "" (empty string)
	default values for bool is false
	default value for type error is nil
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Hello, World!")
	var numerator uint8 = 130
	var denominator uint8 = 2
	var result, remainder, err = intDivision(numerator, denominator)

	switch {
		case err != nil:
				fmt.Println(err.Error())
		case remainder == 0:
				fmt.Printf("The result of the integer division is %v", result)
		default:
				fmt.Printf("The result of the integer division of %v with remainder is %v", result, remainder)
	}
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

func intDivision (numerator uint8, denominator uint8) (uint8, uint8, error) {
	var err error
	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0,0,err
	}

	var result uint8 = numerator / denominator
	var remainder uint8 = numerator % denominator
	return result, remainder, err
}