package main

import "fmt"

func main() {

	// 1. Array - fixed size

	// 1.1 Initialize + declare
	a := [5]int{10, 20, 30, 40, 50} //Defining and populating an array
	a[1] = 25
	fmt.Println("a:", a)

	// 1.2 Declare, then initialize
	var b [5]int
	b[0] = 10
	b[1] = 20
	b[2] = 30
	b[3] = 40
	b[4] = 50
	fmt.Println("b:", b)

	// 2. Slice - dynamic size array
	// 2.1 Can be defined above
	var c []int
	fmt.Println("c empty:", c)
	c = []int{10, 20, 30, 40}
	fmt.Println("c not empty:", c)

	// 2.2 Or like this
	d := []int{10, 20, 30, 40, 50}
	fmt.Println("d:", d)

	// 2.3 Or like this
	s := make([]int, 4)
	s = append(s, 60, 70)
	fmt.Println("s:", s)

	// Manipulate slice
	s[2] = s[len(s)-1]
	fmt.Println("s:", s)
	fmt.Println("s(3-5)", s[2:5])

	// Range
	for k, v := range s {
		fmt.Printf("%d is %d\n", k, v)
	}
}
