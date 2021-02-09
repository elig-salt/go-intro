package main

import (
	"fmt"
	"strings"
)

func main() {
	explain("Hello World")
	explain(52)
	explain(true)

	explain2("Hello world 2")
}

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("I stored string ", strings.ToUpper(i.(string))) // Need to cast
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}

func explain2(i interface{}) {
	// Shorter
	switch i := i.(type) {
	case string:
		fmt.Println("I stored string ", strings.ToUpper(i)) // No need to cast
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}
