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
	myVar := "Hello, World!"

	fmt.Println(myVar)

	// var var1, var2, var3 int = 1, 2, 3
	// var var1, var2, var3 = 1, 2, 3
	var1, var2, var3 := 56, 78, 90
	fmt.Println(var1, var2, var3)
}