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
		FirstName: "Ross",
		LastName:  "Geller",
	}
	// If we wanted to print the full name
	fmt.Println(fullName(e.FirstName, e.LastName))

	// But it sucks to send the whole thing every time
	// So we use recievers (below)

}

func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}
