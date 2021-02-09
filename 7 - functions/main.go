package main

import (
	"fmt"
)

func main() {

	a := Hello()
	fmt.Println(a)

	// Can return multiple values
	b, c := TwoValues()
	fmt.Println(b, c)

	// Functions are values
	d := TwoValues
	fmt.Println(d())
}
func Hello() string {
	return "Hello there"
}

func TwoValues() (string, string) {
	return "Hello", "World"
}
