package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var str string = "Hello, ¥§World!"

	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
}