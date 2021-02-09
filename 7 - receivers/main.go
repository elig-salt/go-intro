package main

import "fmt"

type Employee struct {
	FirstName, LastName string
}

func fullName(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return fullName
}

func main() {
	e := Employee{
		FirstName: "X Æ A-12",
		LastName:  "Musk",
	}
	// If we wanted to print the full name
	fmt.Println("Son's name was", fullName(e.FirstName, e.LastName))

	// But it sucks to send the whole thing every time
	// So we use recievers (below)
	fmt.Printf("IT WAS %s !!!\n", e.fullName())

	e.changeLastName("X AE A-Xii")
	fmt.Printf("But then they changed it to: %s\n", e.fullName())
}

// For quicker access to fields
func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

// For modifying the internals of the struct
func (e *Employee) changeLastName(firstName string) {
	e.FirstName = firstName
}
