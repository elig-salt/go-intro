package main

// Imports are done this way
import "fmt"

func main() {
	// For is everything

	// Simple For loop
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)

	// Simple While loop
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n) // 8 (1*2*2*2)

	// Infinite loop
	// sum := 0
	// for {
	// 		sum++ // repeated forever
	// }
	// fmt.Println(sum) // never reached
}
