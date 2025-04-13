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
	const age = 17
	const pi float32 = 3.1412

	fmt.Println(age)
	fmt.Println(pi)
}