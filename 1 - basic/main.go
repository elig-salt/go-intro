package main

// Imports are done this way
import "fmt"

func main() {

	// Variables can be declared with var
	var b string
	b = "Hello world"

	// Variables can also be declared in a shorter way :=
	// No type needed
	// the colons means we're declaring the variable
	// So we can't use := on the same variable twice
	a := "hello"

	// Once a variable is declared, we use regular =
	a = "world"

	// If does not need round brackets like in Java
	if a == "world" {
		fmt.Println("a is hello")
	} else {
		fmt.Println("a is not hello")
	}

	if b == "" {
		fmt.Println("Primitve zero values are always not nil")
	}
}
