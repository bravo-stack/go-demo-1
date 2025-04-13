package main

import (
	"fmt"

	"github.com/bravo-stack/go-demo-1/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
}