package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func changeNameVal(p Person) {
	p.firstName = "Unchanged"
}
func changeName(p *Person) {
	p.firstName = "Bob"
}

func changeSecondItemVal(arr [5]int) {
	arr[1] = 1
}
func changeSecondItem(arr *[5]int) {
	arr[1] = 1
}

func changeSecondItemSlice(arr []int) {
	arr[1] = 1
}

func main() {
	person := Person{
		firstName: "Alice",
		lastName:  "Dow",
	}

	changeNameVal(person)
	fmt.Println("Did not change", person)

	changeName(&person)
	fmt.Println("Did change", person)

	// Arrays - By VAL!!
	var a = [5]int{}
	changeSecondItemVal(a)
	fmt.Println("Uncahnged", a) // Unchanged
	changeSecondItem(&a)
	fmt.Println("Changed", a) // Changed

	var b = []int{0, 0, 0, 0, 0}
	changeSecondItemSlice(b)
	fmt.Println("Changed Slice by val??", b) // Changed (even though by val?)

	// Array - by val :/
	// Slices, Maps, Channels - always sent by ref :)
}
